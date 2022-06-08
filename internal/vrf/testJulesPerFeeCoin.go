package vrf

import (
	"math/big"

	"github.com/pkg/errors"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

var _ vrf_types.JuelsPerFeeCoin = (*JulesPerFeeCoinWithError)(nil)

type JulesPerFeeCoinWithError struct {
	val *big.Int
}

func (m *JulesPerFeeCoinWithError) SetVal(val *big.Int) {
	m.val = val
}
func (m JulesPerFeeCoinWithError) JuelsPerFeeCoin() (*big.Int, error) {
	return m.val, errors.Errorf("error in JulesPerFeeCoin")
}

func (m JulesPerFeeCoinWithError) AggregateValues(values []*big.Int) (*big.Int, error) {

	return big.NewInt(100), errors.Errorf("error in AggregateValues")

}
