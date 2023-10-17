package ocr2recovery

import (
	"sync"

	"github.com/smartcontractkit/ocr2vrf/internal/dkg"
	"github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	"github.com/smartcontractkit/ocr2vrf/internal/recovery"
)

type keyTransceiver struct {
	keyID contract.KeyID
	kd    *dkg.KeyData
	mu    sync.RWMutex
}

var _ dkg.KeyConsumer = (*keyTransceiver)(nil)
var _ recovery.KeyProvider = (*keyTransceiver)(nil)

func (kt *keyTransceiver) KeyInvalidated(kID contract.KeyID) {

	kt.mu.Lock()
	defer kt.mu.Unlock()

	if kt.keyID == kID {
		kt.kd = nil
	}
}

func (kt *keyTransceiver) NewKey(kID contract.KeyID, kd *dkg.KeyData) {

	kt.mu.Lock()
	defer kt.mu.Unlock()

	if kt.keyID == kID {
		kt.kd = kd.Clone()
	}
}

func (kt *keyTransceiver) KeyLookup(p contract.KeyID) dkg.KeyData {

	kt.mu.RLock()
	defer kt.mu.RUnlock()

	if p == kt.keyID {
		if kt.kd != nil {
			return *kt.kd.Clone()
		}
		return dkg.KeyData{nil, nil, nil, 0, false}
	}

	panic("key consumer is asking for unknown key ID")
}
