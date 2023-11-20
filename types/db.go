package types

import (
	"context"

	ocr_types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/chainlink-vrf/types/hash"
)

type DKGSharePersistence interface {
	WriteShareRecords(
		ctx context.Context,
		cfgDgst ocr_types.ConfigDigest,
		keyID [32]byte,
		shareRecords []PersistentShareSetRecord,
	) error

	ReadShareRecords(
		cfgDgst ocr_types.ConfigDigest,
		keyID [32]byte,
	) (retrievedShares []PersistentShareSetRecord, err error)
}

type PersistentShareSetRecord struct {
	Dealer               player_idx.PlayerIdx
	MarshaledShareRecord []byte
	Hash                 hash.Hash
}
