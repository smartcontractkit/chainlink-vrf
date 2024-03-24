package secp256k1

import (
	"math/big"

	secp256k1BTCD "github.com/btcsuite/btcd/btcec/v2"

	"go.dedis.ch/kyber/v3"
)

type Secp256k1 struct{}

var s256 *secp256k1BTCD.KoblitzCurve = secp256k1BTCD.S256()

func (*Secp256k1) String() string { return "Secp256k1" }

var egScalar kyber.Scalar = newScalar(big.NewInt(0))
var egPoint kyber.Point = &secp256k1Point{newFieldZero(), newFieldZero()}

func (*Secp256k1) ScalarLen() int { return egScalar.MarshalSize() }

func (*Secp256k1) Scalar() kyber.Scalar { return newScalar(big.NewInt(0)) }

func (*Secp256k1) PointLen() int { return egPoint.MarshalSize() }

func (*Secp256k1) Point() kyber.Point {
	return &secp256k1Point{newFieldZero(), newFieldZero()}
}
