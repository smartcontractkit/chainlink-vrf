package recovery

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"go.dedis.ch/kyber/v3/pairing"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	dkg_contract "github.com/smartcontractkit/ocr2vrf/internal/dkg/contract"
	recovery_types "github.com/smartcontractkit/ocr2vrf/types"
)

type sigRequest struct {
	keyID        dkg_contract.KeyID
	keyProvider  KeyProvider
	n            player_idx.Int
	t            player_idx.Int
	configDigest common.Hash
	i            player_idx.PlayerIdx
	pairing      pairing.Suite
	serializer   recovery_types.ReportSerializer
	proofLock    sync.RWMutex

	logger commontypes.Logger

	retransmissionDelay time.Duration
	coordinator         recovery_types.CoordinatorInterface
	reports             map[types.ReportTimestamp]report
	reportsLock         sync.RWMutex
}

func newSigRequest(
	keyID dkg_contract.KeyID,
	keyProvider KeyProvider,
	n player_idx.Int,
	t player_idx.Int,
	configDigest common.Hash,
	i player_idx.PlayerIdx,
	pairing pairing.Suite,
	serializer recovery_types.ReportSerializer,
	retransmissionDelay time.Duration,
	logger commontypes.Logger,
	coordinator recovery_types.CoordinatorInterface,
) (*sigRequest, error) {
	if n <= t {
		return nil, errors.Errorf(
			"committee size must be larger than the fault-tolerance threshold",
		)
	}
	return &sigRequest{
		keyID,
		keyProvider,
		n,
		t,
		configDigest,
		i,
		pairing,
		serializer,
		sync.RWMutex{},
		logger,
		retransmissionDelay,
		coordinator,
		make(map[types.ReportTimestamp]report),
		sync.RWMutex{},
	}, nil
}

type report struct {
	r recovery_types.AbstractReport
	s []byte
}
