package vrf

import (
	"bytes"
	"context"
	"crypto/rand"
	"math/big"
	"sort"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

func OffchainConfig() []byte {

	return nil
}

func OnchainConfig(confDelays map[uint32]struct{}) []byte {
	var onchainConfig [256]byte
	if len(confDelays) != 8 {
		panic("There must be 8 confirmation delays")
	}
	index := 0
	delays := make([]int, 0, 8)
	for delay := range confDelays {
		delays = append(delays, int(delay))
	}
	sort.Ints(delays)
	for _, delay := range delays {
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
	keyID contract.KeyID,
	keyProvider KeyProvider,
	coordinator vrf_types.CoordinatorInterface,
	blockhashes vrf_types.Blockhashes,
	serializer vrf_types.ReportSerializer,
	logger commontypes.Logger,
	juelsPerFeeCoin vrf_types.JuelsPerFeeCoin,
) (types.ReportingPluginFactory, error) {
	contractKeyID, err := coordinator.KeyID(context.Background())
	if err != nil {
		return &vrfReportingPluginFactory{}, errors.Wrap(err, "could not get key ID")
	}
	if keyID != contractKeyID {
		return &vrfReportingPluginFactory{}, errors.New("provided keyID is different from coordinator keyID")
	}
	period, err := coordinator.BeaconPeriod(context.Background())
	if err != nil {
		return &vrfReportingPluginFactory{}, errors.Wrap(err, "could not get beacon period")
	}
	return &vrfReportingPluginFactory{
		&localArgs{
			keyID:           keyID,
			coordinator:     coordinator,
			blockhashes:     blockhashes,
			keyProvider:     keyProvider,
			serializer:      serializer,
			juelsPerFeeCoin: juelsPerFeeCoin,
			period:          period,
			logger:          logger,
			randomness:      rand.Reader,
		},
	}, nil
}
