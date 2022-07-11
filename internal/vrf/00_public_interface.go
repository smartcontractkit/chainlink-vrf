package vrf

import (
	"bytes"
	"context"
	"crypto/rand"
	"math/big"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

func OffchainConfig(
	keyID contract.KeyID,
) []byte {
	return keyID[:]
}

func OnchainConfig(confDelays map[uint32]struct{}) []byte {
	var onchainConfig [256]byte
	if len(confDelays) != 8 {
		panic("There must be 8 confirmation delays")
	}
	index := 0
	for delay, _ := range confDelays {
		delayBigInt := big.NewInt(0).SetUint64(uint64(delay))
		delayBinary := delayBigInt.Bytes()
		paddingBytes := bytes.Repeat([]byte{0}, 32-len(delayBinary))
		delayBinaryFull := bytes.Join([][]byte{paddingBytes, delayBinary}, []byte{})
		copy(onchainConfig[index*32:(index+1)*32], delayBinaryFull)
		index++
	}
	return onchainConfig[:]
}

func NewVRFReportingPluginFactory(
	keyProvider KeyProvider,
	coordinator vrf_types.CoordinatorInterface,
	blockhashes vrf_types.Blockhashes,
	serializer vrf_types.ReportSerializer,
	logger commontypes.Logger,
	juelsPerFeeCoin vrf_types.JuelsPerFeeCoin,
	confirmationDelays map[uint32]struct{},

) types.ReportingPluginFactory {
	period, err := coordinator.BeaconPeriod(context.Background())
	if err != nil {
		panic(err)
	}
	return &vrfReportingPluginFactory{
		&localArgs{
			coordinator, confirmationDelays, blockhashes, keyProvider, serializer,
			juelsPerFeeCoin, period, logger, rand.Reader,
		},
	}
}
