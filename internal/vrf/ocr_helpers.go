package vrf

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	"go.dedis.ch/kyber/v3/share"
	kshare "go.dedis.ch/kyber/v3/share"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/ocr2vrf/internal/vrf/protobuf"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

func (s *sigRequest) ocrsSynced(ctx context.Context) error {
	dkg, vrf, err := s.coordinator.DKGVRFCommittees(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve OCR committees")
	}
	if len(dkg.Signers) != len(vrf.Signers) || len(dkg.Transmitters) != len(vrf.Transmitters) {
		return errors.Errorf("committee sizes differ %s != %s", dkg, vrf)
	}
	for i, s := range dkg.Signers {
		if s != vrf.Signers[i] {
			return errors.Errorf("committee signers differ: %s != %s", s, vrf.Signers[i])
		}
	}
	for i, s := range dkg.Transmitters {
		if s != vrf.Transmitters[i] {
			return errors.Errorf("committee transmitters differ: %s != %s", s, vrf.Transmitters[i])
		}
	}
	keyData := s.keyProvider.KeyLookup(s.keyID)
	if !keyData.Present {
		return errors.Errorf("no distributed key available")
	}
	keyBytes, err := keyData.PublicKey.MarshalBinary()
	if err != nil {
		return errors.Wrap(err, "could not serialize local view of key")
	}
	onchainKeyHash, err := s.coordinator.ProvingKeyHash(ctx)
	if err != nil {
		return errors.Wrap(err, "could not retrieve onchain view of key hash")
	}
	localKeyHash := common.BytesToHash(crypto.Keccak256(keyBytes))
	if localKeyHash != onchainKeyHash {
		return errors.Errorf("keyHash mismatch: 0x%x != 0x%x ", localKeyHash, onchainKeyHash)
	}
	if keyData.SecretShare == nil {
		return errors.Errorf("No local secret keyshare available")
	}
	return nil
}

func getBlock(height uint64, confdelay uint32, blocks []vrf_types.Block) *vrf_types.Block {

	for _, b := range blocks {
		if b.Height == height && b.ConfirmationDelay == confdelay {
			return &b
		}
	}
	return nil
}

func requestIDUint48(
	r uint64, l commontypes.Logger, id commontypes.OracleID,
) (uint64, error) {
	if r > maxUint48.Uint64() {
		return 0, errors.Errorf("requestID too large: %d", r)
		l.Warn(
			"callback requestID too large",
			commontypes.LogFields{"oracleID": id, "bad requestID": r},
		)
	}
	return r, nil
}

func callbackHash(c *protobuf.CostedCallback) (common.Hash, error) {
	s, err := proto.Marshal(c)
	if err != nil {
		return common.Hash{}, errors.Wrap(
			err, "could not serialize callback for indexing",
		)
	}
	return crypto.Keccak256Hash(s), nil
}

func addCallback(
	callbacks map[common.Hash]vrf_types.AbstractCostedCallbackRequest,
	c *protobuf.CostedCallback,
	oracle commontypes.OracleID, confDelays map[uint32]struct{}, beaconPeriod uint16,
	l commontypes.Logger,
) (common.Hash, error) {
	if err := sanityCheckCallback(c, l, oracle, confDelays, beaconPeriod); err != nil {
		return common.Hash{}, err
	}
	h, err := callbackHash(c)
	if err != nil {

		l.Error(
			"could not add callback",
			commontypes.LogFields{"err": err, "source": oracle, "callback": c},
		)
		return common.Hash{}, errors.Wrap(err, "could not add callback")
	}
	callbacks[h] = vrf_types.AbstractCostedCallbackRequest{
		c.Callback.Height,
		c.Callback.ConfDelay,
		c.Callback.SubscriptionID,
		big.NewInt(0).SetBytes(c.Price),
		c.Callback.RequestId,
		uint16(c.Callback.NumWords),
		common.BytesToAddress(c.Callback.Requester),
		c.Callback.Arguments,
		big.NewInt(0).SetBytes(c.GasAllowance),
		c.Callback.RequestHeight,
		common.BytesToHash(c.Callback.RequestBlockHash),
	}
	return h, nil
}

func (s *sigRequest) storeCallbacksByBlocks(costedCallbacks []*protobuf.CostedCallback,
	callbacksByBlock map[heightDelay]map[common.Hash]struct{}, callbackCounts map[common.Hash]uint64,
	callbacks map[common.Hash]vrf_types.AbstractCostedCallbackRequest, observer commontypes.OracleID) {

	seenCallbackHashes := make(
		map[common.Hash]struct{}, len(costedCallbacks),
	)
	for _, cb := range costedCallbacks {
		h, err := addCallback(callbacks, cb, observer, s.confirmationDelays, s.period, s.logger)
		if err != nil {

			continue
		} else {
			if _, present := seenCallbackHashes[h]; present {
				s.logger.Warn("duplicate callback received", commontypes.LogFields{
					"oracleID": observer, "duplicate callback": cb,
				})
				continue
			}
			seenCallbackHashes[h] = struct{}{}
			callbackCounts[h]++
			cbBlock := heightDelay{cb.Callback.Height, cb.Callback.ConfDelay}
			if _, present := callbacksByBlock[cbBlock]; !present {
				callbacksByBlock[cbBlock] = make(map[common.Hash]struct{})
			}
			callbacksByBlock[cbBlock][h] = struct{}{}
		}
	}
}

func (s *sigRequest) parseVRFProofs(
	proofs []*protobuf.VRFResponse,
	vrfContributions map[vrf_types.Block]map[commontypes.OracleID]share.PubShare,
	observer commontypes.OracleID,
	player *player_idx.PlayerIdx,
	kd dkg.KeyData,
) {
	pubShare := player.Index(kd.Shares).(kshare.PubShare)

	seenBlocks := make(map[heightDelay]struct{}, len(proofs))
	for _, output := range proofs {
		if _, present := s.confirmationDelays[output.Delay]; !present {
			s.logger.Warn(
				"block output provided for unknown confirmation delay",
				commontypes.LogFields{
					"oracleID": observer, "delay": output.Delay,
					"known delays": s.confirmationDelays,
				})
			continue
		}
		blockhash := common.BytesToHash(output.Blockhash)
		b := vrf_types.Block{output.Height, output.Delay, blockhash}
		hd := heightDelay{b.Height, b.ConfirmationDelay}
		if _, p := seenBlocks[hd]; p {
			s.logger.Warn(
				"multiple outputs requested for same block/delay pair",
				commontypes.LogFields{"oracleID": observer, "block": b})
			continue
		}
		seenBlocks[hd] = struct{}{}
		contribution := s.pairing.G1().Point()
		if err := contribution.UnmarshalBinary(output.Sig.Sig); err != nil {
			s.logger.Warn("could not read VRF contribution", commontypes.LogFields{
				"oracleID": observer, "error": err,
				"contribution": fmt.Sprintf("0x%x", output.Sig.Sig),
			})
			continue
		}
		h := b.VRFHash(s.configDigest, kd.PublicKey)
		hashPoint := altbn_128.NewHashProof(h).HashPoint
		if _, present := vrfContributions[b]; !present {
			vrfContributions[b] = make(map[commontypes.OracleID]kshare.PubShare)
		}
		p, g2 := s.pairing.Pair, s.pairing.G2().Point().Base()
		if !p(hashPoint, pubShare.V).Equal(p(contribution, g2)) {
			s.logger.Warn("wrong share provided", commontypes.LogFields{
				"oracleID": observer, "sigShare": contribution,
				"keyShare": pubShare.V, "hashPoint": hashPoint,
				"pubKey": kd.PublicKey, "configDigest": s.configDigest, "block": b,
			})
			continue
		}
		vrfContributions[b][observer] = player.PubShare(contribution)
	}
}

func (s *sigRequest) aggregateOutputs(blocks vrf_types.Blocks,
	vrfContributions map[vrf_types.Block]map[commontypes.OracleID]share.PubShare,
	callbacksByBlock map[heightDelay]map[common.Hash]struct{}, callbackCounts map[common.Hash]uint64,
	callbacks map[common.Hash]vrf_types.AbstractCostedCallbackRequest) (outputs []vrf_types.AbstractVRFOutput) {

	outputs = make([]vrf_types.AbstractVRFOutput, 0, len(vrfContributions))
	for _, b := range blocks {
		hd := heightDelay{b.Height, b.ConfirmationDelay}
		delete(callbacksByBlock, hd)
		if len(vrfContributions[b]) <= 2*int(s.n)/3 {
			s.logger.Debug(
				"not enough contributions for block",
				commontypes.LogFields{
					"block": b, "num contributions": len(vrfContributions[b]),
				})
			continue
		}
		shares := make([]*kshare.PubShare, 0, len(vrfContributions[b]))
		for _, c := range vrfContributions[b] {
			shares = append(shares, &kshare.PubShare{c.I, c.V})
		}
		kd := s.keyProvider.KeyLookup(s.keyID)
		output, err := kshare.RecoverCommit(
			s.pairing.G1(), shares, int(kd.T), len(shares),
		)
		if err != nil {
			s.logger.Error(
				"failed to recover distributed VRF output",
				commontypes.LogFields{"error": err, "shares": shares, "t": s.t},
			)
			continue
		}
		h := b.VRFHash(s.configDigest, kd.PublicKey)
		hpoint := altbn_128.NewHashProof(h).HashPoint
		if !validateSignature(s.pairing, hpoint, kd.PublicKey, output) {
			s.logger.Error(
				"could not verify distributed VRF output",
				commontypes.LogFields{"distributed signature": output},
			)
			continue
		}
		proof, err := output.MarshalBinary()
		if err != nil {
			s.logger.Error(
				"could not serialize VRF output for onchain transmission",
				commontypes.LogFields{"error": err, "output": output},
			)
		}
		ccallbacks := make(
			[]vrf_types.AbstractCostedCallbackRequest, 0, len(callbacksByBlock[hd]),
		)

		chashes := make([]string, 0, len(callbacksByBlock[hd]))
		for ch := range callbacksByBlock[hd] {
			chashes = append(chashes, ch.Hex())
		}
		sort.Strings(chashes)
		for _, chs := range chashes {
			ch := common.HexToHash(chs)
			if callbackCounts[ch] > 2*uint64(s.n)/3 {
				ccallbacks = append(ccallbacks, callbacks[ch])
			}
		}
		outputs = append(outputs, vrf_types.AbstractVRFOutput{
			BlockHeight:       b.Height,
			ConfirmationDelay: b.ConfirmationDelay,
			VRFProof:          common.BytesToHash(proof),
			Callbacks:         ccallbacks,
		})
	}
	return
}

func sanityCheckCallback(
	c *protobuf.CostedCallback, l commontypes.Logger, oracle commontypes.OracleID,
	confDelays map[uint32]struct{}, beaconPeriod uint16,
) error {
	if rem := c.Callback.Height % uint64(beaconPeriod); rem != 0 {
		l.Warn("callback with non-beacon height", commontypes.LogFields{
			"height": c.Callback.Height, "period": beaconPeriod, "remainder": rem})
		return errors.Errorf(
			"non-beacon height: %d ∤ %d", beaconPeriod, c.Callback.Height,
		)
	}
	if _, present := confDelays[c.Callback.ConfDelay]; !present {
		l.Warn(
			"unknown confirmation delay",
			commontypes.LogFields{
				"delay": c.Callback.ConfDelay, "good delays": confDelays,
				"source": oracle, "callback": c,
			},
		)
		return errors.Errorf("confirmation delay too large")
	}
	price := big.NewInt(0).SetBytes(c.Price)
	if price.Cmp(MaxPrice) > 0 {
		l.Warn("price too large", commontypes.LogFields{
			"price": price, "max": MaxPrice, "callback": c, "source": oracle,
		})
		return errors.Errorf("price delay too large")
	}
	if uint64(c.Callback.RequestId) > MaxRequestID.Uint64() {
		l.Warn("requestID too large", commontypes.LogFields{
			"requestID": c.Callback.RequestId, "max": maxUint48, "callback": c,
			"source": oracle,
		})
		return errors.Errorf("requestID too large")
	}
	if uint64(c.Callback.NumWords) > MaxNumWords.Uint64() {
		l.Warn("numWords too large", commontypes.LogFields{
			"numWords": c.Callback.NumWords, "max": MaxNumWords, "callback": c,
			"source": oracle,
		})
		return errors.Errorf("numWords too large")
	}
	if len(c.Callback.Requester) > 20 {
		l.Warn("requester bytes too long to be address", commontypes.LogFields{
			"requester": c.Callback.Requester, "maxlen": 20, "source": oracle,
		})
		return errors.Errorf("requester bytes too long")
	}
	allowance := big.NewInt(0).SetBytes(c.GasAllowance)
	if allowance.Cmp(MaxGasAllowance) > 0 {
		l.Warn("gas allowance too large", commontypes.LogFields{
			"allowance": allowance, "max allowance": MaxGasAllowance,
			"source": oracle,
		})
	}
	return nil
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
