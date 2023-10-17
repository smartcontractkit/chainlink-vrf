package recovery

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/ocr2vrf/dkg"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	"github.com/smartcontractkit/ocr2vrf/internal/util"
	recovery_types "github.com/smartcontractkit/ocr2vrf/types"

	"go.dedis.ch/kyber/v3/sign/anon"
)

type dummyDKG struct {
	client            util.CommittingClient
	recoveryClient    common.Address
	kd                contract.OnchainKeyData
	recoveryCommittte recovery_types.OCRCommittee
}

func (d dummyDKG) InitiateDKG(
	ctx context.Context,
	committee recovery_types.OCRCommittee,
	f player_idx.Int,
	keyID contract.KeyID,
	epks contract.EncryptionPublicKeys,
	spks contract.SigningPublicKeys,
	encGroup anon.Suite,
	translator point_translation.PubKeyTranslation,
) error {

	panic("implement me")
}

var _ contract.DKG = (*dummyDKG)(nil)

func (d dummyDKG) GetKey(
	ctx context.Context,
	keyID dkg.KeyID,
	configDigest [32]byte,
) (contract.OnchainKeyData, error) {
	return d.kd, nil
}

func (d *dummyDKG) AddClient(
	ctx context.Context,
	keyID [32]byte,
	clientAddress common.Address,
) error {
	d.recoveryClient = clientAddress
	return nil
}

func (d *dummyDKG) Address() common.Address {
	return common.HexToAddress("0x123456789abcdef12345")
}

func (d dummyDKG) CurrentCommittee(
	ctx context.Context,
) (recovery_types.OCRCommittee, error) {
	return d.recoveryCommittte, nil
}
