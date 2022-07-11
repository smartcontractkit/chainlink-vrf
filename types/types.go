package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
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

type OCRCommittee struct{ Signers, Transmitters []common.Address }

type Block struct {
	Height            uint64
	ConfirmationDelay uint32
	Hash              common.Hash
}

type PairingTranslation = point_translation.PairingTranslation
