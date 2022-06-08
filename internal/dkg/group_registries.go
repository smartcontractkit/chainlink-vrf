package dkg

import (
	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"go.dedis.ch/kyber/v3/sign/anon"
)

var translatorRegistry = point_translation.TranslatorRegistry

var altBN128Pairing = &altbn_128.PairingSuite{}

var testAltBN128Pairing = &altbn_128.XXX_TestPairingSuite{
	PairingSuite:          *altBN128Pairing,
	XXX_NeverUseThisField: xxxTESTONLYTAG,
	Seed:                  10,
}

var encryptionGroupRegistry = map[string]anon.Suite{
	"AltBN-128 G₁":         altBN128Pairing.G1().(anon.Suite),
	"AltBN-128 G₁ Testing": testAltBN128Pairing.G1().(anon.Suite),
}
