package altbn_128

import (
	"crypto/cipher"

	"github.com/smartcontractkit/ocr2vrf/altbn_128/scalar"
	"go.dedis.ch/kyber/v3"
)

type G2 struct{ r cipher.Stream }

var _ kyber.Group = (*G2)(nil)

func (c *G2) String() string {
	return "AltBN-128 G₂"
}

func (c *G2) ScalarLen() int {
	panic("not implemented")
}

func (c *G2) Scalar() kyber.Scalar {
	return scalar.NewScalarInt64(0)
}

func (c *G2) PointLen() int {
	return len(c.Point().(*g2Point).G2.Marshal())
}

func (c *G2) Point() kyber.Point {
	return newG2Point()
}
