package secp256k1

import (
	"crypto/cipher"
	"fmt"
	"io"
	"math/big"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/key"
	"golang.org/x/crypto/sha3"
)

type secp256k1Point struct {
	X *fieldElt
	Y *fieldElt
}

func newPoint() *secp256k1Point {
	return &secp256k1Point{newFieldZero(), newFieldZero()}
}

func (P *secp256k1Point) String() string {
	return fmt.Sprintf("Secp256k1{X: %s, Y: %s}", P.X, P.Y)
}

func (P *secp256k1Point) Equal(pPrime kyber.Point) bool {
	return P.X.Equal(pPrime.(*secp256k1Point).X) &&
		P.Y.Equal(pPrime.(*secp256k1Point).Y)
}

func (P *secp256k1Point) Null() kyber.Point {
	P.X = fieldEltFromInt(0)
	P.Y = fieldEltFromInt(0)
	return P
}

func (P *secp256k1Point) Base() kyber.Point {
	P.X.SetInt(s256.Gx)
	P.Y.SetInt(s256.Gy)
	return P
}

func (P *secp256k1Point) Pick(rand cipher.Stream) kyber.Point {
	for {

		P.X.Set(newFieldZero().Pick(rand))
		maybeRHS := rightHandSide(P.X)
		if maybeY := maybeSqrtInField(maybeRHS); maybeY != (*fieldElt)(nil) {
			P.Y.Set(maybeY)

			b := make([]byte, 1)
			rand.XORKeyStream(b, b)
			if b[0]&1 == 0 {
				P.Y.Neg(P.Y)
			}
			return P
		}
	}
}

func (P *secp256k1Point) Set(pPrime kyber.Point) kyber.Point {
	P.X.Set(pPrime.(*secp256k1Point).X)
	P.Y.Set(pPrime.(*secp256k1Point).Y)
	return P
}

func (P *secp256k1Point) Clone() kyber.Point {
	return &secp256k1Point{X: P.X.Clone(), Y: P.Y.Clone()}
}

func (*secp256k1Point) EmbedLen() int {

	return (255 - 8 - 8) / 8
}

func (P *secp256k1Point) Embed(data []byte, r cipher.Stream) kyber.Point {
	numEmbedBytes := P.EmbedLen()
	if len(data) > numEmbedBytes {
		panic("too much data to embed in a point")
	}
	numEmbedBytes = len(data)
	var x [32]byte
	randStart := 1
	if data != nil {
		x[0] = byte(numEmbedBytes)
		copy(x[1:1+numEmbedBytes], data)
		randStart = 1 + numEmbedBytes
	}
	maxAttempts := 10000

	for numAttempts := 0; numAttempts < maxAttempts; numAttempts++ {

		r.XORKeyStream(x[randStart:], x[randStart:])
		xOrdinate := newFieldZero().SetBytes(x)

		secp256k1RHS := rightHandSide(xOrdinate)
		if maybeY := maybeSqrtInField(secp256k1RHS); maybeY != (*fieldElt)(nil) {
			P.X = xOrdinate
			P.Y = maybeY
			return P
		}
	}

	panic("failed to find point satisfying all constraints")
}

func (P *secp256k1Point) Data() ([]byte, error) {
	b := P.X.Bytes()
	dataLength := int(b[0])
	if dataLength > P.EmbedLen() {
		return nil, fmt.Errorf("point specifies too much data")
	}
	return b[1 : dataLength+1], nil
}

func (P *secp256k1Point) Add(a, b kyber.Point) kyber.Point {
	X, Y := s256.Add(
		a.(*secp256k1Point).X.int(), a.(*secp256k1Point).Y.int(),
		b.(*secp256k1Point).X.int(), b.(*secp256k1Point).Y.int())
	P.X.SetInt(X)
	P.Y.SetInt(Y)
	return P
}

func (P *secp256k1Point) Sub(a, b kyber.Point) kyber.Point {
	X, Y := s256.Add(
		a.(*secp256k1Point).X.int(), a.(*secp256k1Point).Y.int(),
		b.(*secp256k1Point).X.int(),
		newFieldZero().Neg(b.(*secp256k1Point).Y).int())
	P.X.SetInt(X)
	P.Y.SetInt(Y)
	return P
}

func (P *secp256k1Point) Neg(a kyber.Point) kyber.Point {
	P.X = a.(*secp256k1Point).X.Clone()
	P.Y = newFieldZero().Neg(a.(*secp256k1Point).Y)
	return P
}

func (P *secp256k1Point) Mul(s kyber.Scalar, a kyber.Point) kyber.Point {
	sBytes, err := s.(*secp256k1Scalar).MarshalBinary()
	if err != nil {
		panic(fmt.Errorf("failure while marshaling multiplier: %w",
			err))
	}
	var X, Y *big.Int
	if a == (*secp256k1Point)(nil) || a == nil {
		X, Y = s256.ScalarBaseMult(sBytes)
	} else {
		X, Y = s256.ScalarMult(a.(*secp256k1Point).X.int(),
			a.(*secp256k1Point).Y.int(), sBytes)
	}
	P.X.SetInt(X)
	P.Y.SetInt(Y)
	return P
}

func (P *secp256k1Point) MarshalBinary() ([]byte, error) {
	maybeSqrt := maybeSqrtInField(rightHandSide(P.X))
	if maybeSqrt == (*fieldElt)(nil) {
		return nil, fmt.Errorf("x³+7 not a square")
	}
	minusMaybeSqrt := newFieldZero().Neg(maybeSqrt)
	if !P.Y.Equal(maybeSqrt) && !P.Y.Equal(minusMaybeSqrt) {
		return nil, fmt.Errorf(
			"y ≠ ±maybeSqrt(x³+7), so not a point on the curve")
	}
	rv := make([]byte, P.MarshalSize())
	signByte := P.MarshalSize() - 1
	xordinate := P.X.Bytes()
	copyLen := copy(rv[:signByte], xordinate[:])
	if copyLen != P.MarshalSize()-1 {
		return []byte{}, fmt.Errorf("marshal of x ordinate too short")
	}
	if P.Y.isEven() {
		rv[signByte] = 0
	} else {
		rv[signByte] = 1
	}
	return rv, nil
}

func (P *secp256k1Point) MarshalSize() int { return 33 }

func (P *secp256k1Point) MarshalID() [8]byte {
	return [8]byte{'s', 'p', '2', '5', '6', '.', 'p', 'o'}
}

func (P *secp256k1Point) UnmarshalBinary(buf []byte) error {
	var err error
	if len(buf) != P.MarshalSize() {
		err = fmt.Errorf("wrong length for marshaled point")
	}
	if err == nil && !(buf[32] == 0 || buf[32] == 1) {
		err = fmt.Errorf("bad sign byte (the last one)")
	}
	if err != nil {
		return err
	}
	var xordinate [32]byte
	copy(xordinate[:], buf[:32])
	P.X = newFieldZero().SetBytes(xordinate)
	secp256k1RHS := rightHandSide(P.X)
	maybeY := maybeSqrtInField(secp256k1RHS)
	if maybeY == (*fieldElt)(nil) {
		return fmt.Errorf("x ordinate does not correspond to a curve point")
	}
	isEven := maybeY.isEven()
	P.Y.Set(maybeY)
	if (buf[32] == 0 && !isEven) || (buf[32] == 1 && isEven) {
		P.Y.Neg(P.Y)
	} else {
		if buf[32] != 0 && buf[32] != 1 {
			return fmt.Errorf("parity byte must be 0 or 1")
		}
	}
	return nil
}

func (P *secp256k1Point) MarshalTo(w io.Writer) (int, error) {
	buf, err := P.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return w.Write(buf)
}

func (P *secp256k1Point) UnmarshalFrom(r io.Reader) (int, error) {
	buf := make([]byte, P.MarshalSize())
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return 0, err
	}
	return n, P.UnmarshalBinary(buf)
}

func EthereumAddress(p kyber.Point) (rv [20]byte) {

	h := sha3.NewLegacyKeccak256()
	if _, err := h.Write(LongMarshal(p)); err != nil {
		panic(err)
	}
	copy(rv[:], h.Sum(nil)[12:])
	return rv
}

func IsSecp256k1Point(p kyber.Point) bool {
	switch p.(type) {
	case *secp256k1Point:
		return true
	default:
		return false
	}
}

func Coordinates(p kyber.Point) (*big.Int, *big.Int) {
	return p.(*secp256k1Point).X.int(), p.(*secp256k1Point).Y.int()
}

func ValidPublicKey(p kyber.Point) bool {
	if p == (*secp256k1Point)(nil) || p == nil {
		return false
	}
	P, ok := p.(*secp256k1Point)
	if !ok {
		return false
	}
	maybeY := maybeSqrtInField(rightHandSide(P.X))
	return maybeY != nil && (P.Y.Equal(maybeY) || P.Y.Equal(maybeY.Neg(maybeY)))
}

func Generate(random cipher.Stream) *key.Pair {
	p := key.Pair{}
	for !ValidPublicKey(p.Public) {
		p.Private = (&Secp256k1{}).Scalar().Pick(random)
		p.Public = (&Secp256k1{}).Point().Mul(p.Private, nil)
	}
	return &p
}

func LongMarshal(p kyber.Point) []byte {
	xMarshal := p.(*secp256k1Point).X.Bytes()
	yMarshal := p.(*secp256k1Point).Y.Bytes()
	return append(xMarshal[:], yMarshal[:]...)
}

func LongUnmarshal(m []byte) (kyber.Point, error) {
	if len(m) != 64 {
		return nil, fmt.Errorf(
			"0x%x does not represent an uncompressed secp256k1Point. Should be length 64, but is length %d",
			m, len(m))
	}
	p := newPoint()
	p.X.SetInt(big.NewInt(0).SetBytes(m[:32]))
	p.Y.SetInt(big.NewInt(0).SetBytes(m[32:]))
	if !ValidPublicKey(p) {
		return nil, fmt.Errorf("%s is not a valid secp256k1 point", p)
	}
	return p, nil
}

func ScalarToPublicPoint(s kyber.Scalar) kyber.Point {
	publicPoint := (&Secp256k1{}).Point()
	return publicPoint.Mul(s, nil)
}

func SetCoordinates(x, y *big.Int) kyber.Point {
	rv := newPoint()
	rv.X.SetInt(x)
	rv.Y.SetInt(y)
	if !ValidPublicKey(rv) {
		panic("point requested from invalid coordinates")
	}
	return rv
}
