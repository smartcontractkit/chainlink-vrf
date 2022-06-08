package ocr2vrf

import (
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	dkg_contract "github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

type DKGVRFArgs struct {
	Logger commontypes.Logger

	BinaryNetworkEndpointFactory types.BinaryNetworkEndpointFactory

	V2Bootstrappers []commontypes.BootstrapperLocator

	OffchainKeyring types.OffchainKeyring

	OnchainKeyring types.OnchainKeyring

	OffchainConfigDigester types.OffchainConfigDigester

	VRFContractConfigTracker types.ContractConfigTracker

	VRFContractTransmitter types.ContractTransmitter

	VRFDatabase types.Database

	VRFLocalConfig types.LocalConfig

	VRFMonitoringEndpoint commontypes.MonitoringEndpoint

	DKGContractConfigTracker types.ContractConfigTracker

	DKGContract dkg_contract.OnchainContract

	DKGContractTransmitter types.ContractTransmitter

	DKGDatabase types.Database

	DKGLocalConfig types.LocalConfig

	DKGMonitoringEndpoint commontypes.MonitoringEndpoint

	Blockhashes     vrf_types.Blockhashes
	Serializer      vrf_types.ReportSerializer
	JulesPerFeeCoin vrf_types.JuelsPerFeeCoin
	Coordinator     vrf_types.CoordinatorInterface

	ConfirmationDelays []uint32

	Esk   dkg.EncryptionSecretKey
	Ssk   dkg.SigningSecretKey
	KeyID dkg_contract.KeyID
}
