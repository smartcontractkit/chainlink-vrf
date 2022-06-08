package vrf

import (
	"math/big"
	"sort"

	"github.com/pkg/errors"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

var _ vrf_types.JuelsPerFeeCoin = (*medianJulesPerFeeCoin)(nil)

type byValue []*big.Int

func (a byValue) Len() int           { return len(a) }
func (a byValue) Less(i, j int) bool { return a[i].Cmp(a[j]) < 0 }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type medianJulesPerFeeCoin struct {
	val *big.Int
}

func (m *medianJulesPerFeeCoin) SetVal(val *big.Int) {
	m.val = val
}
func (m medianJulesPerFeeCoin) JuelsPerFeeCoin() (*big.Int, error) {
	return m.val, nil
}

func (m medianJulesPerFeeCoin) AggregateValues(values []*big.Int) (*big.Int, error) {
	if len(values) == 0 {
		return nil, errors.Errorf("The array of values is empty")
	}
	sort.Sort(byValue(values))
	medianIndex := len(values) / 2
	return values[medianIndex], nil
}
