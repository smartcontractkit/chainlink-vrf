package dkg

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/hash"
	"github.com/smartcontractkit/ocr2vrf/internal/pvss"
)

type dkgReportingPluginFactory struct {
	l *localArgs

	dkgInProgress bool
	dipMtx        sync.Mutex

	testmode          bool
	xxxDKGTestingOnly *dkg
}

var _ types.ReportingPluginFactory = (*dkgReportingPluginFactory)(nil)

func (d *dkgReportingPluginFactory) NewReportingPlugin(
	c types.ReportingPluginConfig,
) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	d.dipMtx.Lock()
	defer d.dipMtx.Unlock()
	emptyInfo := types.ReportingPluginInfo{}
	if d.dkgInProgress {
		return nil, emptyInfo, errors.Errorf(
			"attempt to initiate DKG round while an earlier DKG round is in progress",
		)
	}
	d.dkgInProgress = true
	a, err := unmarshalPluginConfig(c.OffchainConfig, c.OnchainConfig)
	if err != nil {
		return nil, emptyInfo,
			errors.Wrap(err, "could not read offchain plugin config")
	}
	if c.N > int(player_idx.MaxPlayer) {
		return nil, emptyInfo,
			errors.Errorf("too many players: %d > %d", c.N, player_idx.MaxPlayer)
	}
	args, err := a.NewDKGArgs(
		c.ConfigDigest, d.l, c.OracleID, player_idx.Int(c.N), player_idx.Int(c.F),
	)
	if err != nil {
		return nil, emptyInfo, errors.Wrap(err, "could not construct DKG args")
	}
	dkg, err := d.NewDKG(args)
	if err != nil {
		return nil, emptyInfo, errors.Wrap(err, "while creating reporting plugin")
	}
	if d.testmode {
		d.xxxDKGTestingOnly = dkg
	}
	dkg.keyConsumer.KeyInvalidated(dkg.keyID)
	return dkg, types.ReportingPluginInfo{
		Name: fmt.Sprintf("dkg instance %v", dkg.selfIdx),
		Limits: types.ReportingPluginLimits{
			MaxQueryLength:       1000,
			MaxObservationLength: 100_000,
			MaxReportLength:      10_000,
		},
		UniqueReports: true,
	}, nil
}

func (d *dkgReportingPluginFactory) NewDKG(a *NewDKGArgs) (*dkg, error) {
	if err := a.SanityCheckArgs(); err != nil {
		return nil, errors.Wrap(err, "could not construct new DKG")
	}
	shareSet, err := pvss.NewShareSet(
		a.cfgDgst, a.t, a.selfIdx, a.encryptionGroup, a.translator, a.epks,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not create own share set")
	}
	shareRecords := newShareRecords()
	myShareRecord, err := newShareRecord(a.signingGroup(), shareSet, a.ssk, a.cfgDgst)
	if err != nil {
		return nil, errors.Wrap(err, "could not create own share record")
	}
	shareRecords.set(myShareRecord, hash.Zero)
	completed := false
	return &dkg{
		a.t, a.selfIdx, a.cfgDgst, a.keyID, a.keyConsumer, shareRecords, myShareRecord,
		a.esk, a.epks, a.ssk, a.spks,
		a.encryptionGroup, a.translationGroup, a.translator,
		sync.RWMutex{}, nil, a.contract, completed, d.markCompleted,
		a.logger,
		a.randomness,
	}, nil
}

func (d *dkgReportingPluginFactory) markCompleted() {
	d.dipMtx.Lock()
	defer d.dipMtx.Unlock()
	d.dkgInProgress = false
}

func (d *dkgReportingPluginFactory) SetKeyConsumer(k KeyConsumer) {
	d.l.keyConsumer = k
}
