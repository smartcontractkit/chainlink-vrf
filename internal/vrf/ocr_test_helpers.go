package vrf

import (
	"context"
	"crypto/cipher"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/signatures/cryptotest"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/sign/anon"
	"go.dedis.ch/kyber/v3/util/random"
	"go.dedis.ch/kyber/v3/xof/blake2xb"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	dkg_contract "github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	"github.com/smartcontractkit/ocr2vrf/internal/vrf/protobuf"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"

	"github.com/stretchr/testify/require"
)

func UnmarshalObservation(t *testing.T, observation []byte) *protobuf.Observation {
	obsProto := &protobuf.Observation{}
	err := proto.Unmarshal(observation, obsProto)
	require.Nil(t, err)
	return obsProto
}

var _ vrf_types.CoordinatorInterface = (*CoordinatorReturnsBlocksTooEarly)(nil)

type CoordinatorReturnsBlocksTooEarly struct {
	*TestCoordinator
	backend *backends.SimulatedBackend
	delay   uint32
}

func (c CoordinatorReturnsBlocksTooEarly) ReportBlocks(
	ctx context.Context,
	slotInterval uint16,
	confirmationDelays map[uint32]struct{},
	retransmissionDelay time.Duration, maxBlocks int, maxCallbacks int,
) (
	blocks []vrf_types.Block,
	cbs []vrf_types.AbstractCostedCallbackRequest,
	err error,
) {
	blocks = make([]vrf_types.Block, 0)
	currentHeight := c.backend.Blockchain().CurrentBlock().Number().Uint64()

	maxFulfillBlockheight := uint64(0)

	if currentHeight > uint64(c.delay) {
		maxFulfillBlockheight = (currentHeight - uint64(c.delay)) / uint64(slotInterval) * uint64(slotInterval)
	}

	for currentHeight := uint64(slotInterval); currentHeight <= maxFulfillBlockheight+uint64(slotInterval); currentHeight += uint64(slotInterval) {
		blocks = append(blocks, vrf_types.Block{
			currentHeight,
			c.delay,
			common.Hash{111},
		},
		)
	}
	return blocks, nil, nil
}

var _ vrf_types.CoordinatorInterface = (*CoordinatorReturnsUnknownConfiramtionDelay)(nil)

type CoordinatorReturnsUnknownConfiramtionDelay struct {
	*TestCoordinator
	maxHeight uint64
}

func (c CoordinatorReturnsUnknownConfiramtionDelay) ReportBlocks(
	ctx context.Context,
	slotInterval uint16,
	confirmationDelays map[uint32]struct{},
	retransmissionDelay time.Duration, maxBlocks int, maxCallbacks int,
) (
	blocks []vrf_types.Block,
	cbs []vrf_types.AbstractCostedCallbackRequest,
	err error,
) {
	blocks = make([]vrf_types.Block, 0)
	si := uint64(slotInterval)
	maxConfirmationDelay := uint32(0)
	for cf, _ := range confirmationDelays {
		if cf > maxConfirmationDelay {
			maxConfirmationDelay = cf
		}
	}
	for height := si; height <= c.maxHeight; height += si {
		block := vrf_types.Block{height, maxConfirmationDelay + 1, common.Hash{111}}
		blocks = append(blocks, block)
	}
	return blocks, nil, nil
}

var _ vrf_types.CoordinatorInterface = (*CoordinatorReturnsNonBeaconHeight)(nil)

type CoordinatorReturnsNonBeaconHeight struct {
	*TestCoordinator
	maxHeight uint64
}

func (c CoordinatorReturnsNonBeaconHeight) ReportBlocks(
	ctx context.Context,
	slotInterval uint16,
	confirmationDelays map[uint32]struct{},
	retransmissionDelay time.Duration, maxBlocks int, maxCallbacks int,
) (
	blocks []vrf_types.Block,
	cbs []vrf_types.AbstractCostedCallbackRequest,
	err error,
) {
	blocks = make([]vrf_types.Block, 0)
	sI := uint64(slotInterval)
	for currentHeight := sI + 1; currentHeight <= c.maxHeight; currentHeight += sI {
		for delay := range confirmationDelays {
			block := vrf_types.Block{currentHeight, delay, common.Hash{111}}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil, nil
}

var _ vrf_types.Blockhashes = corruptedBlockhashes{}

type corruptedBlockhashes struct {
	EthereumBlockhashes
	t *testing.T
}

func (e corruptedBlockhashes) OnchainVerifiableBlocks(
	ctx context.Context,
) (startHeight uint64, hashes []common.Hash, err error) {
	hashes = make([]common.Hash, 0, numBlocks)
	zeroHash := common.Hash{}
	currentHeight, err := e.CurrentHeight(ctx)
	startHeight = currentHeight - numBlocks
	for height := startHeight; height <= currentHeight; height++ {
		block := e.Blockchain.GetBlockByNumber(height)
		if block == nil {
			panic("could not get an earlier block")
		}
		if block.Hash() == zeroHash {
			panic("block with zero hash")
		}
		var hash [32]byte
		_, err = rand.Read(hash[:])
		require.NoError(e.t, err)
		hashes = append(hashes, hash)
	}
	return startHeight, hashes, nil
}

var _ vrf_types.Blockhashes = BlockhashesWithErr{}

type BlockhashesWithErr struct {
	EthereumBlockhashes
}

func (e BlockhashesWithErr) OnchainVerifiableBlocks(
	ctx context.Context,
) (startHeight uint64, hashes []common.Hash, err error) {
	hashes = make([]common.Hash, 0, numBlocks)
	zeroHash := common.Hash{}
	currentHeight, err := e.CurrentHeight(ctx)
	startHeight = currentHeight - numBlocks
	block := e.Blockchain.GetBlockByNumber(startHeight)
	for height := startHeight; height <= currentHeight; height++ {
		if block == nil {
			panic("could not get an earlier block")
		}
		if block.Hash() == zeroHash {
			panic("block with zero hash")
		}
		hashes = append(hashes, block.Hash())
	}
	return startHeight, hashes, errors.Errorf("Error in OnchainVerifiableBlocks")
}

var _ vrf_types.Blockhashes = BlockhashesWithoutCurrentBlock{}

type BlockhashesWithoutCurrentBlock struct {
	EthereumBlockhashes
}

func (e BlockhashesWithoutCurrentBlock) OnchainVerifiableBlocks(
	ctx context.Context,
) (startHeight uint64, hashes []common.Hash, err error) {
	hashes = make([]common.Hash, 0, numBlocks)
	zeroHash := common.Hash{}
	currentHeight, err := e.CurrentHeight(ctx)
	startHeight = currentHeight - numBlocks + 1
	block := e.Blockchain.GetBlockByNumber(startHeight)
	for height := startHeight; height < currentHeight; height++ {
		if block == nil {
			panic("could not get earlier block")
		}
		if block.Hash() == zeroHash {
			panic("block with zero hash")
		}
		hashes = append(hashes, block.Hash())
	}
	return startHeight, hashes, nil
}

var _ vrf_types.Blockhashes = BlockhashesCorruptedCurrentHeight{}

type BlockhashesCorruptedCurrentHeight struct {
	*EthereumBlockhashes
}

func (e BlockhashesCorruptedCurrentHeight) CurrentHeight(ctx context.Context) (uint64, error) {
	return uint64(100), errors.Errorf("error in CurrentHeight")
}

var _ vrf_types.CoordinatorInterface = (*coordinatorWithCorruptedKeyHash)(nil)

type coordinatorWithCorruptedKeyHash struct {
	*TestCoordinator
	corruptedPubKey kyber.Point
}

func (c *coordinatorWithCorruptedKeyHash) ProvingKeyHash(
	ctx context.Context,
) (common.Hash, error) {
	keyBytes, err := c.corruptedPubKey.MarshalBinary()
	if err != nil {
		return common.Hash{}, errors.Wrap(
			err, "error while reading s_provingKeyHash",
		)
	}
	keyHash := common.BytesToHash(crypto.Keccak256(keyBytes))
	return keyHash, nil
}

var _ vrf_types.CoordinatorInterface = (*coordinatorDKGVRFCommitteesReturnsError)(nil)

type coordinatorDKGVRFCommitteesReturnsError struct {
	errorNum int
	*TestCoordinator
}

func (c coordinatorDKGVRFCommitteesReturnsError) DKGVRFCommittees(context.Context) (dkg, vrf vrf_types.OCRCommittee, err error) {
	noCommittee := vrf_types.OCRCommittee{}
	if c.errorNum == 1 {
		return noCommittee, noCommittee, errors.Errorf("error in DKGVRFCommittees")
	}

	if c.errorNum == 2 {
		vrfCommittee := vrf_types.OCRCommittee{
			make([]common.Address, 10, 10),
			make([]common.Address, 10, 10),
		}
		dkgComittee := vrf_types.OCRCommittee{
			make([]common.Address, 8, 8),
			make([]common.Address, 10, 10),
		}
		return vrfCommittee, dkgComittee, nil
	}

	vrfCommittee := vrf_types.OCRCommittee{
		make([]common.Address, 10),
		make([]common.Address, 10),
	}
	dkgCommittee := vrf_types.OCRCommittee{
		make([]common.Address, 10),
		make([]common.Address, 10),
	}
	for i := range vrfCommittee.Signers {
		rand.Read(vrfCommittee.Signers[i][:])
		rand.Read(vrfCommittee.Transmitters[i][:])
		if i < 9 {
			copy(dkgCommittee.Signers[i][:], vrfCommittee.Signers[i][:])
			copy(dkgCommittee.Transmitters[i][:], vrfCommittee.Transmitters[i][:])
		} else {
			if c.errorNum == 3 {
				rand.Read(dkgCommittee.Signers[i][:])
				copy(dkgCommittee.Transmitters[i][:], vrfCommittee.Transmitters[i][:])
			} else if c.errorNum == 4 {
				copy(dkgCommittee.Signers[i][:], vrfCommittee.Signers[i][:])
				rand.Read(dkgCommittee.Transmitters[i][:])
			} else {
				panic("Bad error number")
			}
		}

	}
	return vrfCommittee, dkgCommittee, nil
}

var _ vrf_types.CoordinatorInterface = (*coordinatorWithNoKeyHash)(nil)

type coordinatorWithNoKeyHash struct {
	*TestCoordinator
}

func (c coordinatorWithNoKeyHash) ProvingKeyHash(ctx context.Context) (common.Hash, error) {
	return common.Hash{}, errors.Errorf("error in ProvingKeyHash")
}

var _ KeyProvider = (*keyProviderNilPrivateShare)(nil)

type keyProviderNilPrivateShare struct {
	kt *keyTransceiver
}

func (kd keyProviderNilPrivateShare) KeyLookup(p dkg_contract.KeyID) dkg.KeyData {
	return dkg.KeyData{
		kd.kt.kd.PublicKey,
		kd.kt.kd.Shares,
		nil,
		1,
		true,
	}
}

type keyProviderBadPublicKey struct {
	t  *testing.T
	kt *keyTransceiver
}

func (kd keyProviderBadPublicKey) KeyLookup(p dkg_contract.KeyID) dkg.KeyData {
	shares := make([]share.PubShare, 5)
	for i := range shares {
		shares[i] = share.PubShare{i, (&altbn_128.PairingSuite{}).G2().Point().Base()}
	}
	r := cryptotest.NewStream(kd.t, 22999)
	return dkg.KeyData{
		(&altbn_128.PairingSuite{}).G2().Point().Pick(r),
		kd.kt.kd.Shares,
		nil,
		1,
		true,
	}
}

type outputsType []vrf_types.AbstractVRFOutput

var _ sort.Interface = outputsType(nil)

func (os outputsType) Len() int { return len(os) }
func (os outputsType) Less(i, j int) bool {
	if os[i].BlockHeight < os[j].BlockHeight {
		return true
	}
	if os[i].BlockHeight > os[j].BlockHeight {
		return false
	}
	return os[i].ConfirmationDelay < os[j].ConfirmationDelay
}
func (os outputsType) Swap(i, j int) {
	os[i], os[j] = os[j], os[i]
}

type brokenMarshalPoint struct {
	kyber.Point
}

func (b brokenMarshalPoint) MarshalBinary() ([]byte, error) {
	return nil, errors.Errorf("error - brokenMarshalPoint")
}

type brokenMarshalGroup struct {
	kyber.Group
}

func (b brokenMarshalGroup) Point() kyber.Point {
	return &brokenMarshalPoint{b.Group.Point()}
}

func (b brokenMarshalPoint) Pick(r cipher.Stream) kyber.Point { b.Point.Pick(r); return b }
func (b brokenMarshalPoint) Base() kyber.Point                { b.Point.Base(); return b }
func (b brokenMarshalPoint) String() string {
	return fmt.Sprintf("&brokenMarshalPoint{%s}", b.Point)
}
func (b brokenMarshalPoint) Equal(o kyber.Point) bool {
	om, ok := o.(*brokenMarshalPoint)
	return ok && b.Point.Equal(om.Point)
}
func (b brokenMarshalPoint) Add(a, b2 kyber.Point) kyber.Point {
	b.Point.Add(a.(*brokenMarshalPoint).Point, b2.(*brokenMarshalPoint).Point)
	return b
}
func (b brokenMarshalPoint) Sub(a, b2 kyber.Point) kyber.Point {
	b.Point.Sub(a.(*brokenMarshalPoint).Point, b2.(*brokenMarshalPoint).Point)
	return b
}
func (b brokenMarshalPoint) Mul(s kyber.Scalar, p kyber.Point) kyber.Point {
	if p == nil {
		p = &brokenMarshalPoint{b.Point.Clone().Base()}
	}
	b.Point.Mul(s, p.(*brokenMarshalPoint).Point)
	return b
}
func (b brokenMarshalPoint) Null() kyber.Point { b.Point.Null(); return b }
func (b brokenMarshalPoint) Clone() kyber.Point {
	return &brokenMarshalPoint{b.Point.Clone()}
}
func (b brokenMarshalPoint) Neg(p kyber.Point) kyber.Point {
	b.Point.Neg(p.(*brokenMarshalPoint).Point)
	return b
}

func (b *brokenMarshalGroup) XOF(seed []byte) kyber.XOF { return blake2xb.New(seed) }

func (b *brokenMarshalGroup) Write(w io.Writer, objs ...interface{}) error {
	panic("not implemented")
}
func (b *brokenMarshalGroup) Read(r io.Reader, objs ...interface{}) error {
	panic("not implemented")
}
func (b *brokenMarshalGroup) RandomStream() cipher.Stream { return random.New() }

var _ anon.Suite = (*brokenMarshalGroup)(nil)

type pairingWithBadPair struct {
	pairing.Suite
	r cipher.Stream
}

func (p pairingWithBadPair) Pair(p1 kyber.Point, p2 kyber.Point) kyber.Point {
	rp1, rp2 := p1.Pick(p.r), p2.Pick(p.r)
	return p.Suite.Pair(rp1, rp2)
}

var _ vrf_types.CoordinatorInterface = (*CoordinatorReturnsBadCallbacks)(nil)

type CoordinatorReturnsBadCallbacks struct {
	*TestCoordinator
	errorCode int
}

const (
	nonBeaconHeight = iota + 0
	unknownConfirmationDelay
	priceTooLarge
	requestIDTooLarge
	gasAllowanceTooLarge
)

func (c CoordinatorReturnsBadCallbacks) ReportBlocks(
	ctx context.Context,
	slotInterval uint16,
	confirmationDelays map[uint32]struct{},
	retransmissionDelay time.Duration, maxBlocks int, maxCallbacks int,
) (
	blocks []vrf_types.Block,
	cbs []vrf_types.AbstractCostedCallbackRequest,
	err error,
) {
	pendingBlocks, pendingCallbacks, err := c.TestCoordinator.ReportBlocks(ctx,
		slotInterval,
		confirmationDelays,
		retransmissionDelay,
		maxBlocks,
		maxCallbacks,
	)
	maxConfirmationDelay := uint32(0)
	for cf, _ := range confirmationDelays {
		if cf > maxConfirmationDelay {
			maxConfirmationDelay = cf
		}
	}
	numOfCallbacks := len(pendingCallbacks)
	if numOfCallbacks > 0 {
		index := numOfCallbacks / 2
		switch c.errorCode {
		case nonBeaconHeight:
			pendingCallbacks[index].BeaconHeight++
		case unknownConfirmationDelay:
			pendingCallbacks[index].ConfirmationDelay++
		case priceTooLarge:
			bytes := make([]byte, 13, 13)
			bytes[0] = 128
			pendingCallbacks[index].Price = big.NewInt(0).SetBytes(bytes)
		case requestIDTooLarge:
			pendingCallbacks[index].RequestID = MaxRequestID.Uint64() + 1
		case gasAllowanceTooLarge:
			bytes := make([]byte, 13, 13)
			bytes[0] = 128
			pendingCallbacks[index].GasAllowance = big.NewInt(0).SetBytes(bytes)
		}
	}
	return pendingBlocks, pendingCallbacks, nil
}

func isBlockInProof(h uint64, delay uint32, blockhash common.Hash, proofs []*protobuf.VRFResponse) bool {
	for _, proof := range proofs {
		var temp common.Hash
		copy(temp[:], proof.Blockhash[:])
		if proof.Height == h && proof.Delay == delay && temp == blockhash {
			return true
		}
	}
	return false
}
