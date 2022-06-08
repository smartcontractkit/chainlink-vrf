package dkg

import (
	"crypto/rand"
	"sync"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/sign/anon"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
)

func NewReportingPluginFactory(
	esk EncryptionSecretKey,
	ssk SigningSecretKey,
	keyID contract.KeyID,
	contract contract.OnchainContract,
	logger commontypes.Logger,
	keyConsumer KeyConsumer,
) types.ReportingPluginFactory {
	dkgInProgress, testmode, xxxDKGTestingOnly := false, false, (*dkg)(nil)
	return &dkgReportingPluginFactory{
		&localArgs{esk, ssk, keyID, contract, logger, keyConsumer,
			rand.Reader,
		},
		dkgInProgress,
		sync.Mutex{},
		testmode,
		xxxDKGTestingOnly,
	}
}

func NewOnchainContract(
	dkg contract.DKG, keyGroup kyber.Group,
) contract.OnchainContract {
	return contract.OnchainContract{dkg, keyGroup}
}

func OffchainConfig(
	epks EncryptionPublicKeys,
	spks SigningPublicKeys,
	encryptionGroup anon.Suite,
	translator point_translation.PubKeyTranslation,
) ([]byte, error) {
	rc := &offchainConfig{epks, spks, encryptionGroup, translator}
	return rc.MarshalBinary()
}

func OnchainConfig(keyID contract.KeyID) ([]byte, error) {
	return (&onchainConfig{keyID}).Marshal(), nil
}

func NewPluginConfig(
	epks EncryptionPublicKeys,
	spks SigningPublicKeys,
	encryptionGroup anon.Suite,
	translator point_translation.PubKeyTranslation,
	keyID contract.KeyID,
) *PluginConfig {
	return &PluginConfig{
		offchainConfig{epks, spks, encryptionGroup, translator},
		onchainConfig{keyID},
	}
}

func SanityCheckConfigs(p *PluginConfig, rpf types.ReportingPluginFactory) error {
	d, ok := rpf.(*dkgReportingPluginFactory)
	if !ok {
		return errors.Errorf("plugin factory is not for DKG")
	}
	args, err := p.NewDKGArgs([32]byte{}, d.l, 0, 5, 3)
	if err != nil {
		return errors.Wrap(err, "could not construct DKG args")
	}
	return args.SanityCheckArgs()
}

func UnmarshalPluginConfig(offchainBinaryConfig, onchainBinaryConfig []byte) (*PluginConfig, error) {
	return unmarshalPluginConfig(offchainBinaryConfig, onchainBinaryConfig)
}
