package vrf

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	kshare "go.dedis.ch/kyber/v3/share"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/util"
	"github.com/smartcontractkit/ocr2vrf/internal/vrf/protobuf"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

var _ types.ReportingPlugin = (*sigRequest)(nil)

func (s *sigRequest) Query(
	_ context.Context, _ types.ReportTimestamp,
) (types.Query, error) {
	return nil, nil
}

func (s *sigRequest) Observation(
	ctx context.Context, rts types.ReportTimestamp, _ types.Query,
) (types.Observation, error) {
	if err := s.ocrsSynced(ctx); err != nil {
		return nil, errors.Wrap(err, "could not construct observation")
	}
	pendingBlocks, pendingCallbacks, err := s.coordinator.ReportBlocks(
		ctx, s.period, s.confirmationDelays, s.retransmissionDelay, 100, 100,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Observation: could not list pending requests")
	}
	if len(pendingBlocks) == 0 && len(pendingCallbacks) == 0 {
		s.logger.Debug(
			"no observation required on this round",
			commontypes.LogFields{},
		)
		return nil, nil
	}
	currentHeight, err := s.blockhashes.CurrentHeight(ctx)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"could not determine current chain height for confirmation threshold",
		)
	}
	outputs := make([]*protobuf.VRFResponse, 0, len(pendingBlocks))
	s.proofLock.Lock()
	defer s.proofLock.Unlock()
	for _, b := range pendingBlocks {
		if _, present := s.confirmationDelays[b.ConfirmationDelay]; !present {
			s.logger.Error("unknown confirmation delay", commontypes.LogFields{
				"delay": b.ConfirmationDelay, "known delays": s.confirmationDelays,
				"block": b,
			})
			continue
		}
		if b.Height+uint64(b.ConfirmationDelay) >= currentHeight+1 {
			s.logger.Error(
				"ReportBlocks returned a block too early",
				commontypes.LogFields{"block": b, "currentHeight": currentHeight},
			)
			continue
		}
		if remainder := b.Height % uint64(s.period); remainder != 0 {
			s.logger.Error(
				"ReportBlocks returned a non-beacon height",
				commontypes.LogFields{"block": b, "period": s.period, "remainder": remainder},
			)
			continue
		}
		if _, present := s.blockProofs[b]; !present {

			blockProof, err := s.vrfOutput(b, s.keyProvider.KeyLookup(s.keyID))

			if err != nil {
				return nil, err
			}
			s.blockProofs[b] = blockProof
		}
		proofBytes, err := s.blockProofs[b].MarshalBinary()
		if err != nil {
			s.logger.Warn("could not marshal VRF proof", commontypes.LogFields{
				"oracleID": s.i, "error": err,
				"proof": fmt.Sprintf("0x%x", s.blockProofs[b]),
			})
			continue
		}
		var blockhash common.Hash
		copy(blockhash[:], b.Hash[:])
		outputs = append(outputs, &protobuf.VRFResponse{
			Height:    b.Height,
			Delay:     b.ConfirmationDelay,
			Blockhash: blockhash[:],
			Sig:       &protobuf.Signature{Sig: proofBytes[:]},
		})
	}

	callbacks := make([]*protobuf.CostedCallback, 0, len(pendingCallbacks))
	for _, c := range pendingCallbacks {
		var reqBlockHash common.Hash
		copy(reqBlockHash[:], c.RequestBlockHash[:])
		var requester common.Address
		copy(requester[:], c.Requester[:])
		pcb := protobuf.CostedCallback{
			Callback: &protobuf.Callback{
				RequestId:        c.RequestID,
				NumWords:         uint32(c.NumWords),
				Requester:        requester[:],
				Arguments:        c.Arguments,
				SubscriptionID:   c.SubscriptionID,
				Height:           c.BeaconHeight,
				ConfDelay:        c.ConfirmationDelay,
				RequestHeight:    c.RequestHeight,
				RequestBlockHash: reqBlockHash[:],
			},
			Price:        c.Price.Bytes(),
			GasAllowance: c.GasAllowance.Bytes(),
		}
		err := sanityCheckCallback(
			&pcb, s.logger, s.i.OracleID(), s.confirmationDelays, s.period,
		)
		if err != nil {
			continue
		}
		callbacks = append(callbacks, &pcb)
	}

	if (len(outputs) == 0) && (len(callbacks) == 0) {
		s.logger.Error("no valid data to include in report", nil)
		return nil, errors.Errorf("no valid data to include in report")
	}
	juelsPerFeeCoin, err := s.juelsPerFeeCoin.JuelsPerFeeCoin()
	if err != nil {
		return nil, errors.Wrap(err, "error while reading JulesPerFeeCoin")
	}
	if len(juelsPerFeeCoin.Bytes()) > (96 / 8) {
		return nil, errors.Errorf(
			"fee-coin exchange rate too large: %d", juelsPerFeeCoin,
		)
	}
	startHeight, blocks, err := s.blockhashes.OnchainVerifiableBlocks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get verifiable blocks")
	}
	lastVerifiableBlockHeight := startHeight + uint64(len(blocks)) - 1
	if lastVerifiableBlockHeight < currentHeight {
		return nil, errors.Errorf(
			"verifiable blocks don't include current block: %d < %d",
			lastVerifiableBlockHeight, currentHeight,
		)
	}
	recentHashes := make([]*protobuf.RecentBlockAndHash, 0, len(blocks))
	for i, blockhash := range blocks {
		var hash common.Hash
		copy(hash[:], blockhash[:])
		recentHashes = append(
			recentHashes,
			&protobuf.RecentBlockAndHash{
				Height:    startHeight + uint64(i),
				Blockhash: hash[:],
			},
		)
	}

	s.logger.Debug("initial observation", commontypes.LogFields{
		"JulesPerFeeCoin":   juelsPerFeeCoin,
		"RecentBlockHashes": recentHashes,
		"Proofs":            outputs,
		"Callbacks":         callbacks,
		"Raw blocks":        blocks,
	})

	observation := &protobuf.Observation{
		JuelsPerFeeCoin:   juelsPerFeeCoin.Bytes(),
		RecentBlockHashes: recentHashes,
		Proofs:            outputs,
		Callbacks:         callbacks,
	}
	rv, err := proto.Marshal(observation)
	if err != nil {
		return nil, errors.Errorf("Error while marshaling Observation")
	}
	return rv, nil
}

func (s *sigRequest) Report(
	ctx context.Context,
	ts types.ReportTimestamp,
	_ types.Query,
	obs []types.AttributedObservation,
) (bool, types.Report, error) {
	minObs := int(s.t) + 1
	if len(obs) < minObs {
		err := fmt.Errorf("got %d observations, need %d", len(obs), minObs)
		return false, nil, err
	}
	if err := s.ocrsSynced(ctx); err != nil {
		return false, nil, errors.Wrap(err, "Report: ocr is not synced")
	}
	type callback = vrf_types.AbstractCostedCallbackRequest

	callbacks := make(map[common.Hash]callback)

	callbackCounts := make(map[common.Hash]uint64)
	callbacksByBlock := make(map[heightDelay]map[common.Hash]struct{})
	vrfContributions := make(
		map[vrf_types.Block]map[commontypes.OracleID]kshare.PubShare,
	)
	kd := s.keyProvider.KeyLookup(s.keyID)
	players, err := player_idx.PlayerIdxs(s.n)
	if err != nil {
		errMsg := "could not construct players for tracking shares"
		return false, nil, errors.Wrap(err, errMsg)
	}
	juelsPerFeeCoinObs := make([]*big.Int, 0, len(obs))
	type heightHash struct {
		height uint64
		hash   common.Hash
	}
	recentBlockHashes := make(map[heightHash]int, 256*len(obs))
	for _, o := range obs {
		observation := protobuf.Observation{}
		err := proto.Unmarshal(o.Observation, &observation)
		if err != nil {
			s.logger.Warn("failed to parse observation", commontypes.LogFields{
				"oracleID": o.Observer, "observation": o.Observation, "error": err,
			})
			continue
		}

		s.storeCallbacksByBlocks(observation.Callbacks, callbacksByBlock, callbackCounts, callbacks, o.Observer)
		if s.n <= uint8(o.Observer) {
			s.logger.Error(
				"not enough players for observer index",
				commontypes.LogFields{"n": s.n, "oracleID": o.Observer},
			)
			continue
		}
		player := players[o.Observer]

		s.parseVRFProofs(observation.Proofs, vrfContributions, o.Observer, player, kd)
		juelsPerFeeCoinObs = append(juelsPerFeeCoinObs,
			big.NewInt(0).SetBytes(observation.JuelsPerFeeCoin),
		)

		seenHashes := make(
			map[heightHash]struct{}, len(observation.RecentBlockHashes),
		)

		for _, h := range observation.RecentBlockHashes {
			hh := heightHash{h.Height, common.BytesToHash(h.Blockhash)}
			if _, present := seenHashes[hh]; present {
				s.logger.Warn(
					"duplicate hash observed",
					commontypes.LogFields{"hash": hh},
				)
				continue
			} else {

				seenHashes[hh] = struct{}{}
				recentBlockHashes[hh]++
			}
		}
	}

	blocks := make(vrf_types.Blocks, 0, len(vrfContributions))
	for b := range vrfContributions {
		blocks = append(blocks, b)
	}
	sort.Sort(blocks)

	outputs, err := s.aggregateOutputs(
		blocks,
		vrfContributions,
		callbacksByBlock,
		callbackCounts,
		callbacks,
	)
	if err != nil {
		return false, nil, util.WrapError(err, "could not aggregate VRF outputs")
	}

	orphanBlocks := make(hds, 0, len(callbacksByBlock))
	for hd := range callbacksByBlock {
		orphanBlocks = append(orphanBlocks, hd)
	}
	sort.Sort(orphanBlocks)
	for _, hd := range orphanBlocks {
		chashes := make([]string, 0, len(callbacksByBlock[hd]))
		for ch := range callbacksByBlock[hd] {
			chashes = append(chashes, ch.Hex())
		}
		sort.Strings(chashes)
		ccallbacks := make(
			[]vrf_types.AbstractCostedCallbackRequest, 0, len(chashes))
		for _, chs := range chashes {
			ch := common.HexToHash(chs)

			if callbackCounts[ch] > uint64(s.t) {
				ccallbacks = append(ccallbacks, callbacks[ch])
			}
		}
		outputs = append(outputs, vrf_types.AbstractVRFOutput{
			BlockHeight:       hd.height,
			ConfirmationDelay: hd.delay,
			VRFProof:          [32]byte{},
			Callbacks:         ccallbacks,
		})
	}
	if len(outputs) == 0 {
		noFields := commontypes.LogFields{}
		s.logger.Debug(noOutputsRequiredNotTransmittingReport, noFields)
		return false, nil, nil
	}

	var mostRecentBlockHash heightHash
	var zeroHash common.Hash
	for hh, c := range recentBlockHashes {

		if c > int(s.t) {
			if (mostRecentBlockHash.hash == zeroHash) ||
				(hh.height > mostRecentBlockHash.height) ||

				((hh.height == mostRecentBlockHash.height) &&
					hh.hash.Big().Cmp(mostRecentBlockHash.hash.Big()) > 0) {
				mostRecentBlockHash = hh
			}
		}
	}
	if mostRecentBlockHash.hash == zeroHash {
		return false, nil, errors.Errorf(
			"no consensus achieved on most recent block hash",
		)
	}

	abstractReport := vrf_types.AbstractReport{
		outputs, medianBigInt(juelsPerFeeCoinObs), mostRecentBlockHash.height,
		mostRecentBlockHash.hash,
	}
	serializedReport, err := s.serializer.SerializeReport(abstractReport)
	if err != nil {
		s.logger.Error("could not construct serialized report",
			commontypes.LogFields{"err": err},
		)
		return false, types.Report{}, err
	}
	s.reportsLock.Lock()
	defer s.reportsLock.Unlock()
	s.reports[ts] = report{abstractReport, serializedReport}
	return len(outputs) > 0, serializedReport, nil
}

func (s *sigRequest) ShouldAcceptFinalizedReport(
	ctx context.Context, ts types.ReportTimestamp, r types.Report,
) (bool, error) {

	s.reportsLock.Lock()
	if or, present := s.reports[ts]; present && bytes.Equal(or.s, r) {
		if err := s.coordinator.ReportWillBeTransmitted(ctx, or.r); err != nil {
			return false, util.WrapError(err, "Error in ShouldAcceptFinalizedReport")
		}
		delete(s.reports, ts)
	}
	s.reportsLock.Unlock()
	return true, nil
}

func (s *sigRequest) ShouldTransmitAcceptedReport(
	ctx context.Context, ts types.ReportTimestamp, _ types.Report,
) (bool, error) {
	reportIsOnChain, err := s.coordinator.ReportIsOnchain(ctx, ts.Epoch, ts.Round)
	if err != nil {
		return false, util.WrapError(err, "coordinator ReportIsOnchain")
	}
	return !reportIsOnChain, nil
}

func (s *sigRequest) Start() error { return nil }

func (s *sigRequest) Close() error { return nil }

type heightDelay struct {
	height uint64
	delay  uint32
}

type hds []heightDelay

var _ sort.Interface = hds(nil)

func (h hds) Len() int { return len(h) }
func (h hds) Less(i, j int) bool {
	if h[i].height < h[j].height {
		return true
	}
	if h[i].height > h[j].height {
		return false
	}
	return h[i].delay < h[j].delay
}
func (h hds) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

const noOutputsRequiredNotTransmittingReport = "no VRF outputs required; not transmitting report"
const noEnoughContributions = "not enough contributions for block"
const wrongShare = "wrong share provided"
const outOfRangeObserver = "not enough players for observer index"
