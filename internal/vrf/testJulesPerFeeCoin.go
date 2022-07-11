package vrf

import (
	"math/big"

	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

var _ vrf_types.JuelsPerFeeCoin = (*TestJulesPerFeeCoin)(nil)

type TestJulesPerFeeCoin big.Int

func (m TestJulesPerFeeCoin) JuelsPerFeeCoin() (*big.Int, error) {
	r := big.Int(m)
	return &r, nil
}
