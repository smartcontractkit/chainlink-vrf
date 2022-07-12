package vrf

import (
	"fmt"
	"math/big"
	"strings"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/pkg/errors"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/mod"
	kshare "go.dedis.ch/kyber/v3/share"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

func m(x int64) *mod.Int { return mod.NewInt64(x, bn256.P) }

var three = m(3)

func affineCoordinates(p kyber.Point) (*big.Int, *big.Int, error) {

	data, err := p.MarshalBinary()

	if len(data) != 32 {
		return nil, nil, errors.Errorf("attempt to unmarshal g1Point data of wrong length")
	}

	xData := make([]byte, len(data))
	copy(xData, data)
	xData[0] &= 0x7F
	x := mod.NewIntBytes(xData, bn256.P, mod.BigEndian)
	tmp := m(0).Mul(x, x)
	_ = tmp.Mul(tmp, x)
	ySq := tmp.Add(tmp, three)
	y := m(0)
	if !y.Sqrt(ySq) {
		return nil, nil, errors.Errorf("no point on curve with given x ordinate 0x%s", x)
	}
	yParity := (data[0] & 0x80) == 0x80
	yData, err := y.MarshalBinary()
	if err != nil {
		return nil, nil, errors.Wrapf(err, "while marshalling y data")
	}
	currentParity := (yData[31] & 1) == 1
	if yParity != currentParity {

		_ = y.Neg(y)
	}
	yData, err = y.MarshalBinary()
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not re-marshal y after possible negation")
	}
	if (yData[31]&1 == 1) != yParity {
		panic("failed to set correct parity for y")
	}
	return &x.V, &y.V, nil

}

type pointerContributions []*kshare.PubShare

func (c pointerContributions) String() string {
	var rv []string
	for _, cc := range c {
		rv = append(rv, fmt.Sprintf("{I: %d, V: %s}", cc.I, cc.V))
	}
	return fmt.Sprintf("[%s]", strings.Join(rv, ", "))
}

func (s sigRequest) vrfOutput(
	block vrf_types.Block, kd dkg.KeyData,
) (kyber.Point, error) {
	h := block.VRFHash(s.configDigest, kd.PublicKey)

	hpoint := altbn_128.NewHashProof(h).HashPoint
	output := kd.SecretShare.Mul(hpoint)
	pk := s.i.Index(kd.Shares).(kshare.PubShare).V
	if !validateSignature(s.pairing, hpoint, pk, output) {
		return nil, errors.Errorf("could not verify own contribution to signature")
	}
	fmt.Println("KEY SHARES : ", s.i, kd.Shares)
	return output, nil
}
