package dkg

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/sign/anon"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
)

type PluginConfig struct {
	offchainConfig offchainConfig
	onchainConfig  onchainConfig
}

type offchainConfig struct {
	epks []kyber.Point
	spks []kyber.Point

	encryptionGroup anon.Suite
	translator      point_translation.PubKeyTranslation
}

type onchainConfig struct{ contract.KeyID }

func (o *offchainConfig) MarshalBinary() ([]byte, error) {

	rv := [][]byte{{0}}

	if len(o.epks) > int(player_idx.MaxPlayer) {
		return nil, errors.Errorf("too many players")
	}
	if len(o.spks) != len(o.epks) {
		return nil, errors.Errorf(
			"num public keys don't match; len(epks)=%d, len(spks)=%d",
			len(o.epks), len(o.spks),
		)
	}
	rv = append(rv, player_idx.RawMarshal(player_idx.Int(len(o.epks))))

	ename := o.encryptionGroup.String()
	if len(ename) > math.MaxUint8 {
		return nil, errors.Errorf("name for encryption group too long")
	}
	rv = append(rv, []byte{uint8(len(ename))}, []byte(ename))

	for _, pk := range o.epks {
		pkm, err := pk.MarshalBinary()
		if err != nil {
			return nil, errors.Wrap(err, "could not marshal encryption key in PluginConfig")
		}
		rv = append(rv, pkm)
	}

	for _, pk := range o.spks {
		pkm, err := pk.MarshalBinary()
		if err != nil {
			return nil, errors.Wrap(err, "could not marshal signing key in PluginConfig")
		}
		rv = append(rv, pkm)
	}

	tname := o.translator.Name()
	if len(tname) > math.MaxUint8 {
		return nil, errors.Errorf("name for public key translator too long")
	}
	rv = append(rv, []byte{uint8(len(tname))}, []byte(tname))

	return bytes.Join(rv, []byte{}), nil
}

func (o *onchainConfig) Marshal() []byte {
	return append([]byte{}, o.KeyID[:]...)
}

func unmarshalBinaryOffchainConfig(offchainBinaryConfig []byte) (*offchainConfig, error) {

	if len(offchainBinaryConfig) == 0 {
		return nil, errors.Errorf(
			"no binary for reporting plugin offchain config; did you pass it when you set the config?",
		)
	}

	versionNum, offchainBinaryConfig := offchainBinaryConfig[0], offchainBinaryConfig[1:]
	if versionNum != 0 {
		return nil, errors.Errorf(
			"unknown binary version number, %d for reporting plugin offchain config",
			versionNum)
	}

	numPlayers, offchainBinaryConfig, err := player_idx.RawUnmarshal(offchainBinaryConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read number of players")
	}

	if len(offchainBinaryConfig) == 0 {
		return nil, errors.Errorf("missing length for encryption-group name")
	}
	enameLen, offchainBinaryConfig := int(offchainBinaryConfig[0]), offchainBinaryConfig[1:]
	if len(offchainBinaryConfig) < enameLen {
		return nil, errors.Errorf("input too short for expected encryption-group name")
	}
	ename, offchainBinaryConfig := string(offchainBinaryConfig[:enameLen]), offchainBinaryConfig[enameLen:]
	encGroup, ok := encryptionGroupRegistry[ename]
	if !ok {
		return nil, errors.Errorf("unrecognized encryption-group name, '%s'", ename)
	}

	epkLen := encGroup.PointLen()
	epks := make([]kyber.Point, numPlayers)
	for i := range epks {
		var pkm []byte
		if len(offchainBinaryConfig) < encGroup.PointLen() {
			return nil, errors.Errorf("input too short for expected encryption public key")
		}
		pkm, offchainBinaryConfig = offchainBinaryConfig[:epkLen], offchainBinaryConfig[epkLen:]
		epks[i] = encGroup.Point()
		if err := epks[i].UnmarshalBinary(pkm); err != nil {
			return nil, errors.Wrapf(err, "could not read encryption key")
		}
	}

	spkLen := SigningGroup.PointLen()
	spks := make([]kyber.Point, numPlayers)
	for i := range spks {
		var pkm []byte
		if len(offchainBinaryConfig) < SigningGroup.PointLen() {
			return nil, errors.Errorf("input too short for expected signing public key")
		}
		pkm, offchainBinaryConfig = offchainBinaryConfig[:spkLen], offchainBinaryConfig[spkLen:]
		spks[i] = SigningGroup.Point()
		if err := spks[i].UnmarshalBinary(pkm); err != nil {

			return nil, errors.Wrapf(err, "could not read signing key")
		}
	}

	if len(offchainBinaryConfig) == 0 {
		return nil, errors.Errorf("missing translator name length")
	}
	tnameLen, offchainBinaryConfig := int(offchainBinaryConfig[0]), offchainBinaryConfig[1:]
	if len(offchainBinaryConfig) < tnameLen {
		return nil, errors.Errorf("input too short for expected translator name")
	}
	tname, offchainBinaryConfig := string(offchainBinaryConfig[:tnameLen]), offchainBinaryConfig[tnameLen:]
	translator, ok := translatorRegistry[tname]
	if !ok {
		return nil, errors.Errorf("unrecognized translator name")
	}

	if len(offchainBinaryConfig) > 0 {
		return nil, errors.Errorf("overage in PluginConfig representation")
	}

	return &offchainConfig{epks, spks, encGroup, translator}, nil
}

func unmarshalBinaryOnchainConfig(onchainBinaryConfig []byte) (rv onchainConfig, err error) {
	if len(onchainBinaryConfig) != len(contract.KeyID{}) {
		return rv, errors.Errorf("onchainConfig binary is wrong length")
	}
	copy(rv.KeyID[:], onchainBinaryConfig)
	return rv, nil
}

func unmarshalPluginConfig(offchainBinaryConfig, onchainBinaryConfig []byte) (*PluginConfig, error) {
	offchainConfig, err := unmarshalBinaryOffchainConfig(offchainBinaryConfig)
	if err != nil {
		return nil, errors.Wrap(err, "while unmarshaling offchaincomponent of config")
	}
	onchainConfig, err := unmarshalBinaryOnchainConfig(onchainBinaryConfig)
	if err != nil {
		return nil, errors.Wrap(err, "while unmarshaling onchaincomponent of config")
	}
	return &PluginConfig{*offchainConfig, onchainConfig}, nil
}

type NewDKGArgs struct {
	t                          player_idx.Int
	selfIdx                    *player_idx.PlayerIdx
	cfgDgst                    types.ConfigDigest
	esk                        kyber.Scalar
	epks                       []kyber.Point
	ssk                        kyber.Scalar
	spks                       []kyber.Point
	keyID                      contract.KeyID
	keyConsumer                KeyConsumer
	encryptionGroup            anon.Suite
	translationGroup           kyber.Group
	translator                 point_translation.PubKeyTranslation
	contract                   contract.OnchainContract
	logger                     commontypes.Logger
	randomness                 io.Reader
	xxxTestingOnlySigningGroup anon.Suite
}

func (p *PluginConfig) NewDKGArgs(
	d types.ConfigDigest,
	l *localArgs,
	oID commontypes.OracleID,
	n, t player_idx.Int,
) (*NewDKGArgs, error) {
	oc := p.offchainConfig
	translationGroup, err := oc.translator.TargetGroup(oc.encryptionGroup)
	if err != nil {
		return nil, errors.Wrap(err, "could not determine translation target group")
	}
	players, err := player_idx.PlayerIdxs(n)
	if err != nil {
		return nil, errors.Wrap(err, "could not determine local player index")
	}
	selfIdx := players[oID]
	return &NewDKGArgs{
		t, selfIdx, d, l.esk, oc.epks, l.ssk, oc.spks, p.onchainConfig.KeyID,
		l.keyConsumer, oc.encryptionGroup, translationGroup, oc.translator,
		l.contract, l.logger, l.randomness, nil,
	}, nil
}

type localArgs struct {
	esk         kyber.Scalar
	ssk         kyber.Scalar
	keyID       contract.KeyID
	contract    contract.OnchainContract
	logger      commontypes.Logger
	keyConsumer KeyConsumer
	randomness  io.Reader
}

func (oc *offchainConfig) String() string {
	epks := make([]string, len(oc.epks))
	spks := make([]string, len(oc.spks))
	for i := range epks {
		epk, err := oc.epks[i].MarshalBinary()
		if err != nil {
			epks[i] = "unmarshallable: " + err.Error()
		} else {
			epks[i] = hexutil.Encode(epk)
		}
		spk, err := oc.spks[i].MarshalBinary()
		if err != nil {
			epks[i] = "unmarshallable: " + err.Error()
		} else {
			spks[i] = hexutil.Encode(spk)
		}
	}
	return fmt.Sprintf(`PluginConfig{
  epks: %s,
  spks: %s,
  encryptionGroup: %s,
  translator: %s,
}`,
		strings.Join(epks, ", "),
		strings.Join(spks, ", "),
		oc.encryptionGroup,
		oc.translator,
	)
}
