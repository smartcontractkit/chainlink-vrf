package product_group

import (
	"crypto/cipher"
	"io"
	"reflect"

	"go.dedis.ch/fixbuf"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/random"
	"go.dedis.ch/kyber/v3/xof/blake2xb"
)

func (p *ProductGroup) Write(w io.Writer, objs ...interface{}) error {
	return fixbuf.Write(w, objs)
}

func (p *ProductGroup) Read(r io.Reader, objs ...interface{}) error {
	return fixbuf.Read(r, p, objs)
}

var aScalar kyber.Scalar
var aPoint kyber.Point

var tScalar = reflect.TypeOf(&aScalar).Elem()
var tPoint = reflect.TypeOf(&aPoint).Elem()

func (p *ProductGroup) New(t reflect.Type) interface{} {
	switch t {
	case tScalar:
		return p.Scalar()
	case tPoint:
		return p.Point()
	}
	return nil
}

func (p *ProductGroup) XOF(seed []byte) kyber.XOF {
	return blake2xb.New(seed)
}

func (p *ProductGroup) RandomStream() cipher.Stream {
	if p.r != nil {
		return p.r
	}
	return random.New()
}
