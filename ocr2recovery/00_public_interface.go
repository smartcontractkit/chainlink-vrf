package ocr2recovery

import (
	"sync"

	"go.uber.org/multierr"

	"github.com/pkg/errors"
	offchainreporting "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	"github.com/smartcontractkit/ocr2vrf/internal/recovery"
	"github.com/smartcontractkit/ocr2vrf/internal/util"
)

type OCR2Recovery struct {
	dkg, recovery  *offchainreporting.Oracle
	keyTransceiver *keyTransceiver
}

type EthereumReportSerializer = recovery.EthereumReportSerializer

func NewOCR2Recovery(a DKGRecoveryArgs) (*OCR2Recovery, error) {
	transceiver := keyTransceiver{a.KeyID, nil, sync.RWMutex{}}
	dkgReportingPluginFactory := dkg.NewReportingPluginFactory(
		a.Esk,
		a.Ssk,
		a.KeyID,
		a.DKGContract,
		a.DKGLogger,
		&transceiver,
		a.DKGSharePersistence,
	)

	recoveryReportingPluginFactory, err := recovery.NewRecoveryReportingPluginFactory(
		a.KeyID,
		&transceiver,
		a.Coordinator,
		a.Serializer,
		a.RecoveryLogger,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate Recovery reporting plugin factory")
	}

	if a.DKGReportingPluginFactoryDecorator != nil {
		dkgReportingPluginFactory = a.DKGReportingPluginFactoryDecorator(dkgReportingPluginFactory)
	}

	if a.RecoveryReportingPluginFactoryDecorator != nil {
		recoveryReportingPluginFactory = a.RecoveryReportingPluginFactoryDecorator(recoveryReportingPluginFactory)
	}

	deployedDKG, err := offchainreporting.NewOracle(offchainreporting.OracleArgs{
		BinaryNetworkEndpointFactory: a.BinaryNetworkEndpointFactory,
		V2Bootstrappers:              a.V2Bootstrappers,
		ContractConfigTracker:        a.DKGContractConfigTracker,
		ContractTransmitter:          a.DKGContractTransmitter,
		Database:                     a.DKGDatabase,
		LocalConfig:                  a.DKGLocalConfig,
		Logger:                       a.DKGLogger,
		MonitoringEndpoint:           a.DKGMonitoringEndpoint,
		OffchainConfigDigester:       a.DKGOffchainConfigDigester,
		OffchainKeyring:              a.OffchainKeyring,
		OnchainKeyring:               a.OnchainKeyring,
		ReportingPluginFactory:       dkgReportingPluginFactory,
	})
	if err != nil {
		return nil, util.WrapError(err, "while setting up new DKG oracle")
	}
	confirmationDelays := make(map[uint32]struct{}, len(a.ConfirmationDelays))
	for _, d := range a.ConfirmationDelays {
		confirmationDelays[d] = struct{}{}
	}

	deployedRecovery, err := offchainreporting.NewOracle(offchainreporting.OracleArgs{
		BinaryNetworkEndpointFactory: a.BinaryNetworkEndpointFactory,
		V2Bootstrappers:              a.V2Bootstrappers,
		ContractConfigTracker:        a.RecoveryContractConfigTracker,
		ContractTransmitter:          a.RecoveryContractTransmitter,
		Database:                     a.RecoveryDatabase,
		LocalConfig:                  a.RecoveryLocalConfig,
		Logger:                       a.RecoveryLogger,
		MonitoringEndpoint:           a.RecoveryMonitoringEndpoint,
		OffchainConfigDigester:       a.RecoveryOffchainConfigDigester,
		OffchainKeyring:              a.OffchainKeyring,
		OnchainKeyring:               a.OnchainKeyring,
		ReportingPluginFactory:       recoveryReportingPluginFactory,
	})
	if err != nil {
		return nil, util.WrapError(err, "while setting up Recoveryy oracle")
	}
	return &OCR2Recovery{deployedDKG, deployedRecovery, &transceiver}, nil
}

func OffchainConfig() []byte {
	return recovery.OffchainConfig()
}

func OnchainConfig(confDelays map[uint32]struct{}) []byte {
	return recovery.OnchainConfig(confDelays)
}

func (o *OCR2Recovery) Start() error {
	if err := o.dkg.Start(); err != nil {
		return util.WrapError(err, "starting DKG oracle")
	}
	if err := util.WrapError(o.recovery.Start(), "starting Recovery oracle"); err != nil {
		return multierr.Append(err, util.WrapError(
			o.dkg.Close(),
			"closing DKG process after starting Recovery process failed",
		))
	}
	return nil
}

func (o *OCR2Recovery) Close() error {
	return multierr.Append(
		util.WrapError(o.dkg.Close(), "while closing DKG process"),
		util.WrapError(o.recovery.Close(), "while closing Recovery process"),
	)
}
