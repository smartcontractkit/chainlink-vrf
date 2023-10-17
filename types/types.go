package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/ocr2vrf/internal/common/ocr"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
)

type CoordinatorInterface interface {
	DKGRecoveryCommittees(context.Context) (dkg, committee OCRCommittee, err error)

	ProvingKeyHash(context.Context) (common.Hash, error)

	KeyID(ctx context.Context) (contract.KeyID, error)

	ContractID(context.Context) common.Address
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

type AbstractReport struct {
	AccountToRecover common.Address
	Recoverer        common.Address
	Success          bool
}

type (
	PubKeyTranslation  = point_translation.PubKeyTranslation
	PairingTranslation = point_translation.PairingTranslation
	PlayerIdx          = player_idx.PlayerIdx
	PlayerIdxInt       = player_idx.Int
	OCRCommittee       = ocr.OCRCommittee
)

func UnmarshalPlayerIdx(b []byte) (*PlayerIdx, []byte, error) {
	return player_idx.Unmarshal(b)
}

func RawMarshalPlayerIdxInt(i PlayerIdxInt) []byte {
	return player_idx.RawMarshal(i)
}
