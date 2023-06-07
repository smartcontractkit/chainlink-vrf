package secp256k1

import (
	"crypto/cipher"
	"fmt"
	"math/big"

	"go.dedis.ch/kyber/v3/util/random"
)

var q = s256.P

type fieldElt big.Int

func newFieldZero() *fieldElt { return (*fieldElt)(big.NewInt(0)) }

func (f *fieldElt) int() *big.Int { return (*big.Int)(f) }

func (f *fieldElt) modQ() *fieldElt {
	if f.int().Cmp(q) != -1 || f.int().Cmp(bigZero) == -1 {

		f.int().Mod(f.int(), q)
	}
	return f
}

func fieldEltFromBigInt(v *big.Int) *fieldElt { return (*fieldElt)(v).modQ() }

func fieldEltFromInt(v int64) *fieldElt {
	return fieldEltFromBigInt(big.NewInt(int64(v))).modQ()
}

var fieldZero = fieldEltFromInt(0)
var bigZero = big.NewInt(0)

func (f *fieldElt) String() string {
	return fmt.Sprintf("fieldElt{%x}", f.int())
}

func (f *fieldElt) Equal(g *fieldElt) bool {
	if f == (*fieldElt)(nil) && g == (*fieldElt)(nil) {
		return true
	}
	if f == (*fieldElt)(nil) {
		return false
	}
	if g == (*fieldElt)(nil) {
		return false
	}
	return bigZero.Cmp(newFieldZero().Sub(f, g).modQ().int()) == 0
}

func (f *fieldElt) Add(a, b *fieldElt) *fieldElt {
	f.int().Add(a.int(), b.int())
	return f.modQ()
}

func (f *fieldElt) Sub(a, b *fieldElt) *fieldElt {
	f.int().Sub(a.int(), b.int())
	return f.modQ()
}

func (f *fieldElt) Set(v *fieldElt) *fieldElt {
	f.int().Set(v.int())
	return f.modQ()
}

func (f *fieldElt) SetInt(v *big.Int) *fieldElt {
	f.int().Set(v)
	return f.modQ()
}

func (f *fieldElt) Pick(rand cipher.Stream) *fieldElt {
	return f.SetInt(random.Int(q, rand))
}

func (f *fieldElt) Neg(g *fieldElt) *fieldElt {
	f.int().Neg(g.int())
	return f.modQ()
}

func (f *fieldElt) Clone() *fieldElt { return newFieldZero().Set(f.modQ()) }

func (f *fieldElt) SetBytes(buf [32]byte) *fieldElt {
	f.int().SetBytes(buf[:])
	return f.modQ()
}

func (f *fieldElt) Bytes() [32]byte {
	bytes := f.modQ().int().Bytes()
	if len(bytes) > 32 {
		panic("field element longer than 256 bits")
	}
	var rv [32]byte
	copy(rv[32-len(bytes):], bytes)
	return rv
}

var two = big.NewInt(2)

func fieldSquare(y *fieldElt) *fieldElt {
	return fieldEltFromBigInt(newFieldZero().int().Exp(y.int(), two, q))
}

var sqrtPower = s256.QPlus1Div4()

func maybeSqrtInField(v *fieldElt) *fieldElt {
	s := newFieldZero()
	s.int().Exp(v.int(), sqrtPower, q)
	if !fieldSquare(s).Equal(v) {
		return nil
	}
	return s
}

var three = big.NewInt(3)
var seven = fieldEltFromInt(7)

func rightHandSide(x *fieldElt) *fieldElt {
	xCubed := newFieldZero()
	xCubed.int().Exp(x.int(), three, q)
	return xCubed.Add(xCubed, seven)
}

func (f *fieldElt) isEven() bool {
	return big.NewInt(0).Mod(f.int(), two).Cmp(big.NewInt(0)) == 0
}
