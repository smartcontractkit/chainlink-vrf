package secp256k1

import (
	"crypto/cipher"
	"fmt"
	"io"
	"math/big"

	secp256k1BTCD "github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/random"
)

var GroupOrder = secp256k1BTCD.S256().N
var FieldSize = secp256k1BTCD.S256().P

type secp256k1Scalar big.Int

func (s *secp256k1Scalar) AllowVarTime(varTimeAllowed bool) {

	if !varTimeAllowed {
		panic("implementation is not constant-time!")
	}
}

func newScalar(v *big.Int) kyber.Scalar {
	return (*secp256k1Scalar)(zero().Mod(v, GroupOrder))
}

func zero() *big.Int { return big.NewInt(0) }

func ToInt(s kyber.Scalar) *big.Int { return (*big.Int)(s.(*secp256k1Scalar)) }

func (s *secp256k1Scalar) int() *big.Int { return (*big.Int)(s) }

func (s *secp256k1Scalar) modG() kyber.Scalar {

	s.int().Mod(s.int(), GroupOrder)
	return s
}

func (s *secp256k1Scalar) String() string {
	return fmt.Sprintf("scalar{%x}", (*big.Int)(s))
}

var scalarZero = zero()

func (s *secp256k1Scalar) Equal(sPrime kyber.Scalar) bool {
	difference := zero().Sub(s.int(), ToInt(sPrime))
	return scalarZero.Cmp(difference.Mod(difference, GroupOrder)) == 0
}

func (s *secp256k1Scalar) Set(sPrime kyber.Scalar) kyber.Scalar {
	return (*secp256k1Scalar)(s.int().Mod(ToInt(sPrime), GroupOrder))
}

func (s *secp256k1Scalar) Clone() kyber.Scalar {
	return (*secp256k1Scalar)(zero().Mod(s.int(), GroupOrder))
}

func (s *secp256k1Scalar) SetInt64(v int64) kyber.Scalar {
	return (*secp256k1Scalar)(s.int().SetInt64(v)).modG()
}

func (s *secp256k1Scalar) Zero() kyber.Scalar {
	return s.SetInt64(0)
}

func (s *secp256k1Scalar) Add(a, b kyber.Scalar) kyber.Scalar {
	s.int().Add(ToInt(a), ToInt(b))
	return s.modG()
}

func (s *secp256k1Scalar) Sub(a, b kyber.Scalar) kyber.Scalar {
	s.int().Sub(ToInt(a), ToInt(b))
	return s.modG()
}

func (s *secp256k1Scalar) Neg(a kyber.Scalar) kyber.Scalar {
	s.int().Neg(ToInt(a))
	return s.modG()
}

func (s *secp256k1Scalar) One() kyber.Scalar {
	return s.SetInt64(1)
}

func (s *secp256k1Scalar) Mul(a, b kyber.Scalar) kyber.Scalar {

	s.int().Mul(ToInt(a), ToInt(b))
	return s.modG()
}

func (s *secp256k1Scalar) Div(a, b kyber.Scalar) kyber.Scalar {
	if ToInt(b).Cmp(scalarZero) == 0 {
		panic("attempt to divide by zero")
	}

	s.int().Mul(ToInt(a), zero().ModInverse(ToInt(b), GroupOrder))
	return s.modG()
}

func (s *secp256k1Scalar) Inv(a kyber.Scalar) kyber.Scalar {
	if ToInt(a).Cmp(scalarZero) == 0 {
		panic("attempt to divide by zero")
	}
	s.int().ModInverse(ToInt(a), GroupOrder)
	return s
}

func (s *secp256k1Scalar) Pick(rand cipher.Stream) kyber.Scalar {
	return s.Set((*secp256k1Scalar)(random.Int(GroupOrder, rand)))
}

func (s *secp256k1Scalar) MarshalBinary() ([]byte, error) {
	b := ToInt(s.modG()).Bytes()

	rv := append(make([]byte, s.MarshalSize()-len(b)), b...)
	if len(rv) != s.MarshalSize() {
		return nil, fmt.Errorf("marshalled scalar to wrong length")
	}
	return rv, nil
}

func (s *secp256k1Scalar) MarshalSize() int { return 32 }

func (s *secp256k1Scalar) MarshalID() [8]byte {
	return [8]byte{'s', 'p', '2', '5', '6', '.', 's', 'c'}
}

func (s *secp256k1Scalar) UnmarshalBinary(buf []byte) error {
	if len(buf) != s.MarshalSize() {
		return fmt.Errorf("cannot unmarshal to scalar: wrong length")
	}
	s.int().Mod(s.int().SetBytes(buf), GroupOrder)
	return nil
}

func (s *secp256k1Scalar) MarshalTo(w io.Writer) (int, error) {
	buf, err := s.MarshalBinary()
	if err != nil {
		return 0, fmt.Errorf("cannot marshal binary: %w", err)
	}
	return w.Write(buf)
}

func (s *secp256k1Scalar) UnmarshalFrom(r io.Reader) (int, error) {
	buf := make([]byte, s.MarshalSize())
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return n, err
	}
	return n, s.UnmarshalBinary(buf)
}

func (s *secp256k1Scalar) SetBytes(a []byte) kyber.Scalar {
	return ((*secp256k1Scalar)(s.int().SetBytes(a))).modG()
}

func IsSecp256k1Scalar(s kyber.Scalar) bool {
	switch s := s.(type) {
	case *secp256k1Scalar:
		s.modG()
		return true
	default:
		return false
	}
}

func IntToScalar(i *big.Int) kyber.Scalar {
	return ((*secp256k1Scalar)(i)).modG()
}

func ScalarToHash(s kyber.Scalar) common.Hash {
	return common.BigToHash(ToInt(s.(*secp256k1Scalar)))
}

func RepresentsScalar(i *big.Int) bool {
	return i.Cmp(GroupOrder) == -1
}
