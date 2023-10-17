package recovery

import (
	"bytes"
	"context"
	"math"
	"math/big"
	"sort"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (s *sigRequest) ocrsSynced(ctx context.Context) error {
	deployedDKG, deployedRecovery, err := s.coordinator.DKGRecoveryCommittees(ctx)
	if err != nil {
		return errors.Wrap(err, failedRetrieveOCRCommitteesMsg)
	}
	if len(deployedDKG.Signers) != len(deployedRecovery.Signers) ||
		len(deployedDKG.Transmitters) != len(deployedRecovery.Transmitters) {
		return errors.Errorf(
			committeesWithDifferentSizesMsg+" %s != %s", deployedDKG, deployedRecovery,
		)
	}
	for i, s := range deployedDKG.Signers {
		if s != deployedRecovery.Signers[i] {
			return errors.Errorf(
				signersMismatchMsg+" %s != %s", s, deployedRecovery.Signers[i],
			)
		}
	}
	for i, s := range deployedDKG.Transmitters {
		if s != deployedRecovery.Transmitters[i] {
			return errors.Errorf(
				transmittersMismatchMsg+" %s != %s", s, deployedRecovery.Transmitters[i],
			)
		}
	}
	keyData := s.keyProvider.KeyLookup(s.keyID)
	if !keyData.Present {
		return errors.Errorf(noDistributedKeyMsg)
	}
	keyBytes, err := keyData.PublicKey.MarshalBinary()
	if err != nil {
		return errors.Wrap(err, failedSerializeLocalKey)
	}
	onchainKeyHash, err := s.coordinator.ProvingKeyHash(ctx)
	if err != nil {
		return errors.Wrap(err, failedRetrieveOnchainKeyMsg)
	}
	localKeyHash := common.BytesToHash(crypto.Keccak256(keyBytes))
	if localKeyHash != onchainKeyHash {
		return errors.Errorf(incorrectPublicKeyMsg+" : 0x%x != 0x%x ", localKeyHash, onchainKeyHash)
	}
	if keyData.SecretShare == nil {
		return errors.Errorf(noLocalShareMsg)
	}
	return nil
}

func sortBigInt(l []*big.Int) []*big.Int {
	sort.Sort(byValue(l))
	return l
}

type byValue []*big.Int

func (a byValue) Len() int           { return len(a) }
func (a byValue) Less(i, j int) bool { return a[i].Cmp(a[j]) < 0 }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func medianBigInt(l []*big.Int) *big.Int {
	sortBigInt(l)
	midPoint := len(l) / 2
	if len(l)%2 == 1 {
		return l[midPoint]
	}
	if len(l) == 0 {

		panic("list must be populated")
	}

	midPointTotal := big.NewInt(0).Add(l[midPoint-1], l[midPoint])
	return midPointTotal.Div(midPointTotal, big.NewInt(2))
}

var (
	maxUint16 = big.NewInt(0).SetUint64(math.MaxUint16)
	maxUint24 = big.NewInt(0).SetBytes(bytes.Repeat([]byte{0xff}, 3))
	maxUint48 = big.NewInt(0).SetBytes(bytes.Repeat([]byte{0xff}, 6))
	maxUint32 = big.NewInt(0).SetUint64(math.MaxUint32)
	maxUint64 = big.NewInt(0).SetUint64(math.MaxUint64)
	maxUint96 = big.NewInt(0).SetBytes(bytes.Repeat([]byte{0xff}, 12))
)

var (
	MaxNumWords          = maxUint16
	MaxConfirmationDelay = maxUint24
	MaxRequestID         = maxUint48
	MaxPrice             = maxUint96
	MaxGasAllowance      = maxUint96
	MaxSubscriptionID    = maxUint64
)

func init() {

	if MaxNumWords.Cmp(maxUint32) > 0 {
		panic("MaxNumWords needs new backing type")
	}
	if MaxConfirmationDelay.Cmp(maxUint32) > 0 {
		panic("MaxConfirmationDelay needs new backing type")
	}
	if MaxRequestID.Cmp(maxUint64) > 0 {
		panic("MaxRequestID needs new backing type")
	}
	if MaxSubscriptionID.Cmp(maxUint64) > 0 {
		panic("MaxSubcriptionID needs new backing type")
	}
}

const (
	excessGasAllowanceMsg              = "gas allowance too large"
	unknownConfirmationDelayMsg        = "uknown confirmation delay"
	nonBeaconHeightInCallbackMsg       = "callback with non-beacon height"
	priceTooLargeMsg                   = "price too large"
	requestIdTooLargeMsg               = "requestID too large"
	noLocalShareMsg                    = "No local secret keyshare available"
	incorrectPublicKeyMsg              = "keyHash mismatch"
	noDistributedKeyMsg                = "no distributed key available"
	failedSerializeLocalKey            = "could not serialize local view of key"
	failedRetrieveOCRCommitteesMsg     = "failed to retrieve OCR committees"
	committeesWithDifferentSizesMsg    = "committee sizes differ"
	signersMismatchMsg                 = "committee signers differ"
	transmittersMismatchMsg            = "committee transmitters differ"
	failedRetrieveOnchainKeyMsg        = "could not retrieve onchain view of key hash"
	failedReadContributionMsg          = "could not read Recovery contribution"
	nonBeaconHeightInBlockMsg          = "block output provided for non-beacon height"
	unknownConfirmationDelayInBlockMsg = "block output provided for unknown confirmation delay"
)
