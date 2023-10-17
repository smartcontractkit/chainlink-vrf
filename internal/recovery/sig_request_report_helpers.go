package recovery

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/mod"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
)

func m(x int64) *mod.Int { return mod.NewInt64(x, bn256.P) }

var three = m(3)

func affineCoordinates(p kyber.Point) (*big.Int, *big.Int) {
	b := altbn_128.LongMarshal(p)
	if len(b) != 64 {
		panic("wrong length for marshaled point")
	}
	return big.NewInt(0).SetBytes(b[:32]), big.NewInt(0).SetBytes(b[32:])
}

const (
	failedVerifyOwnContributionMsg = "could not verify own contribution to signature"
)
