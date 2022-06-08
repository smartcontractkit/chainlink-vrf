package dkg

import (
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/sign/anon"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
)

func NewReportingPluginFactory(
	esk EncryptionSecretKey,
	ssk SigningSecretKey,
	keyID KeyID,
	contract OnchainContract,
	logger commontypes.Logger,
	keyConsumer KeyConsumer,
) types.ReportingPluginFactory {
	return dkg.NewReportingPluginFactory(
		esk,
		ssk,
		keyID,
		contract,
		logger,
		keyConsumer,
	)
}

func NewOnchainContract(
	dkg DKG, keyGroup kyber.Group,
) contract.OnchainContract {
	return contract.OnchainContract{dkg, keyGroup}
}

func OffchainConfig(
	epks EncryptionPublicKeys,
	spks SigningPublicKeys,
	encryptionGroup anon.Suite,
	translator PubKeyTranslation,
) ([]byte, error) {
	return dkg.OffchainConfig(epks, spks, encryptionGroup, translator)
}

func OnchainConfig(keyID KeyID) ([]byte, error) {
	return dkg.OnchainConfig(keyID)
}

func NewPluginConfig(
	epks EncryptionPublicKeys,
	spks SigningPublicKeys,
	encryptionGroup anon.Suite,
	translator PubKeyTranslation,
	keyID KeyID,
) *PluginConfig {
	return dkg.NewPluginConfig(epks, spks, encryptionGroup, translator, keyID)
}

func SanityCheckConfigs(
	p *PluginConfig,
	rpf types.ReportingPluginFactory,
) error {
	return dkg.SanityCheckConfigs(p, rpf)
}

func UnmarshalPluginConfig(
	offchainBinaryConfig, onchainBinaryConfig []byte) (*PluginConfig, error) {
	return dkg.UnmarshalPluginConfig(offchainBinaryConfig, onchainBinaryConfig)
}

type (
	EncryptionPublicKeys = dkg.EncryptionPublicKeys
	EncryptionSecretKey  = dkg.EncryptionSecretKey
	SigningPublicKeys    = dkg.SigningPublicKeys
	SigningSecretKey     = dkg.SigningSecretKey
	PluginConfig         = dkg.PluginConfig
	KeyConsumer          = dkg.KeyConsumer

	KeyID           = contract.KeyID
	DKG             = contract.DKG
	OnchainContract = contract.OnchainContract

	PubKeyTranslation = point_translation.PubKeyTranslation
)
