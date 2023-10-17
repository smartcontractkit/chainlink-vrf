package ocr2recovery

import (
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	dkg_contract "github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	recovery_types "github.com/smartcontractkit/ocr2vrf/types"
)

type DKGRecoveryArgs struct {
	DKGLogger      commontypes.Logger
	RecoveryLogger commontypes.Logger

	BinaryNetworkEndpointFactory types.BinaryNetworkEndpointFactory

	V2Bootstrappers []commontypes.BootstrapperLocator

	OffchainKeyring types.OffchainKeyring

	OnchainKeyring types.OnchainKeyring

	DKGOffchainConfigDigester types.OffchainConfigDigester

	RecoveryOffchainConfigDigester types.OffchainConfigDigester

	RecoveryContractConfigTracker types.ContractConfigTracker

	RecoveryContractTransmitter types.ContractTransmitter

	RecoveryDatabase types.Database

	RecoveryLocalConfig types.LocalConfig

	RecoveryMonitoringEndpoint commontypes.MonitoringEndpoint

	DKGContractConfigTracker types.ContractConfigTracker

	DKGContract dkg_contract.OnchainContract

	DKGContractTransmitter types.ContractTransmitter

	DKGDatabase types.Database

	DKGLocalConfig types.LocalConfig

	DKGMonitoringEndpoint commontypes.MonitoringEndpoint

	DKGSharePersistence recovery_types.DKGSharePersistence

	Serializer  recovery_types.ReportSerializer
	Coordinator recovery_types.CoordinatorInterface

	ConfirmationDelays []uint32

	Esk   dkg_contract.EncryptionSecretKey
	Ssk   dkg_contract.SigningSecretKey
	KeyID dkg_contract.KeyID

	DKGReportingPluginFactoryDecorator      func(factory types.ReportingPluginFactory) types.ReportingPluginFactory
	RecoveryReportingPluginFactoryDecorator func(factory types.ReportingPluginFactory) types.ReportingPluginFactory
}
