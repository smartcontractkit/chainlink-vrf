package vrf

import (
	"fmt"
	"io"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	dkg_contract "github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

type vrfReportingPluginFactory struct {
	l *localArgs
}

var _ types.ReportingPluginFactory = (*vrfReportingPluginFactory)(nil)

type localArgs struct {
	coordinator        vrf_types.CoordinatorInterface
	confirmationDelays map[uint32]struct{}
	blockhashes        vrf_types.Blockhashes
	keyProvider        KeyProvider
	serializer         vrf_types.ReportSerializer
	juelsPerFeeCoin    vrf_types.JuelsPerFeeCoin
	period             uint16

	logger     commontypes.Logger
	randomness io.Reader
}

func (v *vrfReportingPluginFactory) NewReportingPlugin(
	c types.ReportingPluginConfig,
) (types.ReportingPlugin, types.ReportingPluginInfo, error) {

	var keyID dkg_contract.KeyID
	copy(keyID[:], c.OffchainConfig[:32])

	if c.N > int(player_idx.MaxPlayer) {
		return nil, types.ReportingPluginInfo{},
			errors.Errorf("too many players: %d > %d", c.N, player_idx.MaxPlayer)
	}
	players, err := player_idx.PlayerIdxs(player_idx.Int(c.N))
	if err != nil {
		return nil, types.ReportingPluginInfo{},
			errors.Wrap(err, "could not determine local player DKG index")
	}
	tbls, err := newSigRequest(
		keyID,
		v.l.keyProvider,
		player_idx.Int(c.N),
		player_idx.Int(c.F),
		common.Hash(c.ConfigDigest),
		*players[c.OracleID],
		&altbn_128.PairingSuite{},
		v.l.blockhashes,
		v.l.serializer,
		time.Hour,
		v.l.logger,
		v.l.juelsPerFeeCoin,
		v.l.coordinator,
		v.l.confirmationDelays,
		v.l.period,
	)
	if err != nil {
		return nil, types.ReportingPluginInfo{},
			errors.Wrap(err, "could not create new VRF Beacon reporting plugin")
	}
	return tbls, types.ReportingPluginInfo{
		Name: fmt.Sprintf("vrf instance %v", tbls.i),
		Limits: types.ReportingPluginLimits{
			MaxQueryLength:       200000,
			MaxObservationLength: 200000,
			MaxReportLength:      200000,
		},
	}, nil
}
