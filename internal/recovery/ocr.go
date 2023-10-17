package recovery

import (
	"bytes"
	"context"
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/twystd/tweetnacl-go/tweetnacl"
	kshare "go.dedis.ch/kyber/v3/share"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/gethwrappers/recoverybeacon"
	recovery_types "github.com/smartcontractkit/ocr2vrf/types"
	commonUtil "github.com/smartcontractkit/ocr2vrf/util"
)

var cipherSetMapping map[common.Address]struct{} = make(map[common.Address]struct{})
var reportedMapping map[common.Address]struct{} = make(map[common.Address]struct{})

var _ types.ReportingPlugin = (*sigRequest)(nil)

func (s *sigRequest) Query(
	_ context.Context, _ types.ReportTimestamp,
) (types.Query, error) {
	return nil, nil
}

type ObservationPayload struct {
	marshalledShare  []byte
	addressToRecover common.Address
	recoverer        common.Address
}

func (s *sigRequest) Observation(
	ctx context.Context, rts types.ReportTimestamp, _ types.Query,
) (types.Observation, error) {
	if err := s.ocrsSynced(ctx); err != nil {
		return nil, errors.Wrap(err, failedConstructObservation)
	}

	owner, ec := commonUtil.GetOwnerAndClient()

	keyData := s.keyProvider.KeyLookup(s.keyID)
	secretShare := keyData.SecretShare
	distributedPublicKeyBytes, err := keyData.PublicKey.MarshalBinary()
	commonUtil.PanicErr(err)

	beacon, err := recoverybeacon.NewRecoveryBeacon(s.coordinator.ContractID(ctx), ec)
	commonUtil.PanicErr(err)

	recoveryRequests, err := beacon.GetRecoveryRequests(nil)
	commonUtil.PanicErr(err)
	if len(recoveryRequests) > 0 {
		for _, r := range recoveryRequests {

			if bytes.Equal(r.RecoveryPubKeyBytes, commonUtil.NodePublicKeyEncryptionKeyPair.PublicKey) {

				if !bytes.Equal(r.DistributedPublicKey, distributedPublicKeyBytes) {
					PanicErr(errors.New("distributed key does not match request"))
				}

				keyData := s.keyProvider.KeyLookup(s.keyID)
				secretShare := keyData.SecretShare
				answerPoint := commonUtil.GetPointFromCipher(r.Cipher, r.Nonce, commonUtil.NodePublicKeyEncryptionKeyPair, r.EphermeralKeyPubKeyBytes)
				recoveryShare := secretShare.Mul(answerPoint)
				marshalled, err := recoveryShare.MarshalBinary()
				PanicErr(err)
				if len(marshalled) != 32 {
					PanicErr(errors.New("marshalled point size is not correct"))
				}

				if (len(r.AddressToRecover.Bytes()) != 20) || (len(r.Recoverer.Bytes()) != 20) {
					PanicErr(errors.New(fmt.Sprintf(
						"addresses are not correct lengths %v %v",
						len(r.AddressToRecover.Bytes()),
						len(r.Recoverer.Bytes()),
					)))
				}

				marshalledObservation := marshalled
				marshalledObservation = append(marshalledObservation, r.AddressToRecover.Bytes()...)
				marshalledObservation = append(marshalledObservation, r.Recoverer.Bytes()...)
				PanicErr(err)
				s.logger.Info("[RECOVERY OCR] WE HAVE SUCCESSFULLY MARSHALLED AN OBSERVATION", commontypes.LogFields{})
				return marshalledObservation, nil
			}
		}
	}

	request, err := beacon.GetMostRecentEnrollmentRequest(nil)
	commonUtil.PanicErr(err)

	if (len(request.AddressToEnroll) == 0) || (request.AddressToEnroll == common.Address{}) {
		s.logger.Info("[RECOVERY OCR] NO ADDRESS FOUND", commontypes.LogFields{})
		cipherSetMapping = make(map[common.Address]struct{})
		reportedMapping = make(map[common.Address]struct{})
		return nil, nil
	}

	_, ok := cipherSetMapping[request.AddressToEnroll]
	if ok {
		s.logger.Info("[RECOVERY OCR] ALREADY TRANSMITTED", commontypes.LogFields{})
		return nil, nil
	}

	recipientPublicKey := request.PublicKeyBytes

	acctPoint := altbn_128.NewHashProof(common.HexToHash(request.AddressToEnroll.String())).HashPoint
	acctPointBytes, err := acctPoint.MarshalBinary()
	share := secretShare.Mul(acctPoint)

	kd := s.keyProvider.KeyLookup(s.keyID)
	pk := s.i.Index(kd.Shares).(kshare.PubShare).V
	valid := ValidateSignature(s.pairing, acctPoint, pk, share)
	if !valid {
		s.logger.Info("[RECOVERY OCR] SHARE IS NOT VALID", commontypes.LogFields{"pk": pk.String()})
		return nil, errors.New("share is not valid")
	}

	ephemeralKeyPair, err := tweetnacl.CryptoBoxKeyPair()
	PanicErr(err)

	cipher, ephemeralPubKeyBytes, nonce := commonUtil.GetCipherFromPoint(share, recipientPublicKey, ephemeralKeyPair)
	if len(cipher) != 48 {
		PanicErr(errors.New("enrollment share encrypted incorrectly"))
	}
	tx, err := beacon.PostCipher(owner, request.AddressToEnroll, recoverybeacon.RecoveryBeaconTypesEnrollmentResponse{
		PlayerIdx:                s.i.GetIdx(),
		Threshold:                kd.T,
		Cipher:                   cipher,
		EphermeralKeyPubKeyBytes: ephemeralPubKeyBytes,
		Nonce:                    nonce,
		DistributedPublicKey:     distributedPublicKeyBytes,
		AccountPointBytes:        acctPointBytes,
		RecoveryPubKeyBytes:      commonUtil.NodePublicKeyEncryptionKeyPair.PublicKey,
	})
	fmt.Printf("tx %s\n", tx.Hash())
	if err != nil {
		s.logger.Info("[RECOVERY OCR] ERROR SENDING TX", commontypes.LogFields{"err": err})
	} else {
		s.logger.Info("[RECOVERY OCR] POSTING ON-CHAIN", commontypes.LogFields{"tx": tx.Hash()})
		cipherSetMapping[request.AddressToEnroll] = struct{}{}
	}

	return nil, nil

}

func (s *sigRequest) Report(
	ctx context.Context,
	ts types.ReportTimestamp,
	_ types.Query,
	obs []types.AttributedObservation,
) (bool, types.Report, error) {

	s.logger.Info("[RECOVERY OCR] GENERATING REPORT", commontypes.LogFields{"ts": ts})

	_, ec := commonUtil.GetOwnerAndClient()
	beacon, err := recoverybeacon.NewRecoveryBeacon(s.coordinator.ContractID(ctx), ec)
	commonUtil.PanicErr(err)

	for i, o := range obs {
		if len(o.Observation) == 0 {
			s.logger.Info("[RECOVERY OCR] GETTING EMPTY OBSERVATIONS", commontypes.LogFields{"index": i})
			return false, nil, nil
		}
	}

	var shares []*kshare.PubShare
	var recoverer common.Address = common.Address{}
	var addressToRecover common.Address = common.Address{}
	for _, o := range obs {

		marshalledShare := o.Observation[:32]
		marshalledAddressToRecover := o.Observation[32:52]
		marshalledRecoveryer := o.Observation[52:]

		if (len(marshalledShare) != 32) || (len(marshalledAddressToRecover) != 20) || (len(marshalledRecoveryer) != 20) {
			PanicErr(errors.New(fmt.Sprintf(
				"addresses are not correct lengths %v %v %v",
				len(marshalledShare),
				len(marshalledAddressToRecover),
				len(marshalledRecoveryer),
			)))
		}

		obsAddressToRecover := common.BytesToAddress(marshalledAddressToRecover)
		obsRecoveryer := common.BytesToAddress(marshalledRecoveryer)
		if (recoverer != common.Address{}) && (recoverer != obsRecoveryer) {
			commonUtil.PanicErr(errors.New(fmt.Sprintf("recoverer is not correct %v", recoverer)))
		}
		if (addressToRecover != common.Address{}) && (addressToRecover != obsAddressToRecover) {
			commonUtil.PanicErr(errors.New(fmt.Sprintf("addressToRecover is not correct %v", addressToRecover)))
		}
		addressToRecover = obsAddressToRecover
		recoverer = obsRecoveryer

		sharePoint := commonUtil.G1Point.Clone()
		err = sharePoint.UnmarshalBinary(marshalledShare)
		commonUtil.PanicErr(err)
		shares = append(shares, &kshare.PubShare{
			I: int(o.Observer),
			V: sharePoint,
		})
		s.logger.Info("[RECOVERY OCR] MARSHALLED SHARE", commontypes.LogFields{"share": o.Observation})
	}

	if len(shares) < int(s.t+1) {
		s.logger.Info("[RECOVERY OCR] NOT ENOUGH SHARES", commontypes.LogFields{"shareCount": len(shares)})
		return false, nil, nil
	}

	recovery, err := beacon.GetRecovery(nil, addressToRecover)
	commonUtil.PanicErr(err)

	commit, err := kshare.RecoverCommit(s.pairing.G1(), shares, int(s.t+1), len(shares))
	commonUtil.PanicErr(err)
	commitBinary, err := commit.MarshalBinary()
	commonUtil.PanicErr(err)

	recoveryIsValid := bytes.Equal(commitBinary, recovery)
	if recoveryIsValid {
		s.logger.Info("[RECOVERY OCR] RECOVERY IS VALID", commontypes.LogFields{"commitBinary": commitBinary})
	} else {
		s.logger.Info("[RECOVERY OCR] RECOVERY IS NOT VALID", commontypes.LogFields{"commitBinary": commitBinary, "recovery": recovery})
	}

	r, err := s.serializer.SerializeReport(recovery_types.AbstractReport{
		AccountToRecover: addressToRecover,
		Recoverer:        recoverer,
		Success:          recoveryIsValid,
	})
	commonUtil.PanicErr(err)

	return true, r, nil
}

func (s *sigRequest) ShouldAcceptFinalizedReport(
	ctx context.Context, ts types.ReportTimestamp, r types.Report,
) (bool, error) {
	return true, nil
}

func (s *sigRequest) ShouldTransmitAcceptedReport(
	ctx context.Context, ts types.ReportTimestamp, _ types.Report,
) (bool, error) {
	return true, nil
}

func (s *sigRequest) Start() error { return nil }

func (s *sigRequest) Close() error { return nil }

type heightDelay struct {
	height uint64
	delay  uint32
}

type hds []heightDelay

var _ sort.Interface = hds(nil)

func (h hds) Len() int { return len(h) }
func (h hds) Less(i, j int) bool {
	if h[i].height < h[j].height {
		return true
	}
	if h[i].height > h[j].height {
		return false
	}
	return h[i].delay < h[j].delay
}
func (h hds) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

const (
	failedConstructObservation             = "could not construct observation"
	failedListPendingBlocks                = "Observation: could not list pending requests"
	noOutputsRequiredNotTransmittingReport = "no Recovery outputs required; not transmitting report"
	notEnoughContributions                 = "not enough contributions for block"
	wrongShare                             = "wrong share provided"
	outOfRangeObserver                     = "not enough players for observer index"
	noObservationInRound                   = "no observation required on this round"
	failedReadJulesPerFeeCoin              = "error while reading JulesPerFeeCoin"
	failedReadVerifiableBlocks             = "could not get verifiable blocks"
	failedMarshalObservation               = "Error while marshaling Observation"
	unknownConfirmationDelay               = "unknown confirmation delay"
	earlyBlockReportBlocks                 = "ReportBlocks returned a block too early"
	invalidBlockReportBlocks               = "ReportBlocks returned a non-beacon height"
	failedMarshalRecoveryProof             = "could not marshal Recovery proof"
	noValidDataToIncludeInReport           = "no valid data to include in report"
	currentBlockIsNotInVerifiableBlocks    = "verifiable blocks don't include current block:"
	largeFeeCoin                           = "fee-coin exchange rate too large:"
	failedReadCurrentHeight                = "could not determine current chain height for confirmation threshold"
	initialObservation                     = "initial observation"
	failedVerifyRecoveryOutput             = "could not verify distributed Recovery output"
	failedParseObservation                 = "failed to parse observation"
	duplicateHashErr                       = "duplicate hash observed"
	noConsensusOnRecentBlockhash           = "no consensus achieved on most recent block hash"
	notEnoughAppearancesCallback           = "insufficient number of appearances for a callback"
	skipErrMsg                             = "skipping callback due to error"
	noConsensusOnOrphanBlockCallbacksMsg   = "there is no consensus on any of the callbacks of an orphan block"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
