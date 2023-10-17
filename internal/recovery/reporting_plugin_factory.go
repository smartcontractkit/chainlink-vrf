package recovery

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	dkg_contract "github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	recovery_types "github.com/smartcontractkit/ocr2vrf/types"
)

type recoveryReportingPluginFactory struct {
	l *localArgs
}

var _ types.ReportingPluginFactory = (*recoveryReportingPluginFactory)(nil)

type localArgs struct {
	keyID       dkg_contract.KeyID
	coordinator recovery_types.CoordinatorInterface
	keyProvider KeyProvider
	serializer  recovery_types.ReportSerializer

	logger commontypes.Logger
}

func (v *recoveryReportingPluginFactory) NewReportingPlugin(
	c types.ReportingPluginConfig,
) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
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
		v.l.keyID,
		v.l.keyProvider,
		player_idx.Int(c.N),
		player_idx.Int(c.F),
		common.Hash(c.ConfigDigest),
		*players[c.OracleID],
		&altbn_128.PairingSuite{},
		v.l.serializer,
		time.Hour,
		v.l.logger,
		v.l.coordinator,
	)
	if err != nil {
		return nil, types.ReportingPluginInfo{},
			errors.Wrap(err, "could not create new Recovery Beacon reporting plugin")
	}
	return tbls, types.ReportingPluginInfo{
		Name: fmt.Sprintf("recovery instance %v", tbls.i),
		Limits: types.ReportingPluginLimits{
			MaxQueryLength:       200000,
			MaxObservationLength: 200000,
			MaxReportLength:      200000,
		},
	}, nil
}
