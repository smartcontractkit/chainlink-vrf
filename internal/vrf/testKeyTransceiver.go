package vrf

import (
	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
)

type keyTransceiver struct {
	keyID contract.KeyID
	kd    *dkg.KeyData
}

var _ dkg.KeyConsumer = (*keyTransceiver)(nil)
var _ KeyProvider = (*keyTransceiver)(nil)

func (kt *keyTransceiver) KeyInvalidated(kID contract.KeyID) {
	if kt.keyID == kID {
		kt.kd = nil
	}
}

func (kt *keyTransceiver) NewKey(kID contract.KeyID, kd *dkg.KeyData) {
	if kt.keyID == kID {
		kt.kd = kd.Clone()
	}
}

func (kt *keyTransceiver) KeyLookup(p contract.KeyID) dkg.KeyData {
	if p == kt.keyID && kt.kd != nil {
		return *kt.kd.Clone()
	}
	if kt.kd == nil {
		panic("key consumer is asking for unavailable key")
	}

	panic("key consumer is asking for unknown key ID")
}
