package secp256k1

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"golang.org/x/crypto/sha3"

	"go.dedis.ch/fixbuf"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/random"
	"go.dedis.ch/kyber/v3/xof/blake2xb"
)

type SuiteSecp256k1 struct {
	Secp256k1
	r cipher.Stream
}

func (s *SuiteSecp256k1) Hash() hash.Hash {
	return sha3.NewLegacyKeccak256()
}

func (s *SuiteSecp256k1) XOF(key []byte) kyber.XOF {
	return blake2xb.New(key)
}

func (s *SuiteSecp256k1) Read(r io.Reader, objs ...interface{}) error {
	return fixbuf.Read(r, s, objs...)
}

func (s *SuiteSecp256k1) Write(w io.Writer, objs ...interface{}) error {
	return fixbuf.Write(w, objs)
}

var aScalar kyber.Scalar
var tScalar = reflect.TypeOf(aScalar)
var aPoint kyber.Point
var tPoint = reflect.TypeOf(aPoint)

func (s *SuiteSecp256k1) New(t reflect.Type) interface{} {
	switch t {
	case tScalar:
		return s.Scalar()
	case tPoint:
		return s.Point()
	}
	return nil
}

func (s *SuiteSecp256k1) RandomStream() cipher.Stream {
	if s.r != nil {
		return s.r
	}
	return random.New()
}

func NewBlakeKeccackSecp256k1() *SuiteSecp256k1 {
	return new(SuiteSecp256k1)
}
