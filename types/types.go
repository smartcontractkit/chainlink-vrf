package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/ocr2vrf/internal/common/ocr"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
)

type CoordinatorInterface interface {
	ReportBlocks(
		ctx context.Context,
		slotInterval uint16,
		confirmationDelays map[uint32]struct{},
		retransmissionDelay time.Duration,
		maxBlocks, maxCallbacks int,
	) (
		blocks []Block,
		callbacks []AbstractCostedCallbackRequest,
		err error,
	)

	ReportWillBeTransmitted(context.Context, AbstractReport) error

	DKGVRFCommittees(context.Context) (dkg, vrf OCRCommittee, err error)

	ProvingKeyHash(context.Context) (common.Hash, error)

	BeaconPeriod(ctx context.Context) (uint16, error)

	ReportIsOnchain(
		ctx context.Context,
		epoch uint32, round uint8,
	) (presentOnchain bool, err error)

	ConfirmationDelays(ctx context.Context) ([]uint32, error)

	KeyID(ctx context.Context) (contract.KeyID, error)
}

type ReportSerializer interface {
	SerializeReport(AbstractReport) ([]byte, error)

	DeserializeReport([]byte) (AbstractReport, error)

	MaxReportLength() uint

	ReportLength(AbstractReport) uint
}

type JuelsPerFeeCoin interface {
	JuelsPerFeeCoin() (*big.Int, error)
}

type Blockhashes interface {
	OnchainVerifiableBlocks(
		context.Context,
	) (startHeight uint64, hashes []common.Hash, err error)

	CurrentHeight(context.Context) (uint64, error)
}

type AbstractCostedCallbackRequest struct {
	BeaconHeight      uint64
	ConfirmationDelay uint32
	SubscriptionID    uint64
	Price             *big.Int
	RequestID         uint64
	NumWords          uint16
	Requester         common.Address
	Arguments         []byte
	GasAllowance      *big.Int
	RequestHeight     uint64
	RequestBlockHash  common.Hash
}

type AbstractVRFOutput struct {
	BlockHeight       uint64
	ConfirmationDelay uint32
	VRFProof          [32]byte
	Callbacks         []AbstractCostedCallbackRequest
}

type AbstractReport struct {
	Outputs           []AbstractVRFOutput
	JulesPerFeeCoin   *big.Int
	RecentBlockHeight uint64
	RecentBlockHash   common.Hash
}

type Block struct {
	Height            uint64
	ConfirmationDelay uint32
	Hash              common.Hash
}

type (
	PubKeyTranslation  = point_translation.PubKeyTranslation
	PairingTranslation = point_translation.PairingTranslation
	PlayerIdxInt       = player_idx.Int
	OCRCommittee       = ocr.OCRCommittee
)
