package vrf

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/util"
	"github.com/smartcontractkit/ocr2vrf/internal/vrf/protobuf"
	vrf_types "github.com/smartcontractkit/ocr2vrf/types"

	kshare "go.dedis.ch/kyber/v3/share"

	"google.golang.org/protobuf/proto"
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
		return nil, errors.Wrap(err, failedConstructObservation)
	}
	pendingBlocks, pendingCallbacks, err := s.coordinator.ReportBlocks(
		ctx, s.period, s.confirmationDelays, s.retransmissionDelay, 100, 100,
	)
	if err != nil {
		return nil, errors.Wrap(err, failedListPendingBlocks)
	}
	if len(pendingBlocks) == 0 && len(pendingCallbacks) == 0 {
		s.logger.Debug(
			noObservationInRound,
			commontypes.LogFields{},
		)
		return nil, nil
	}
	currentHeight, err := s.blockhashes.CurrentHeight(ctx)
	if err != nil {
		return nil, errors.Wrap(
			err,
			failedReadCurrentHeight,
		)
	}
	outputs := make([]*protobuf.VRFResponse, 0, len(pendingBlocks))
	s.proofLock.Lock()
	defer s.proofLock.Unlock()
	for _, b := range pendingBlocks {
		if _, present := s.confirmationDelays[b.ConfirmationDelay]; !present {
			s.logger.Error(unknownConfirmationDelay, commontypes.LogFields{
				"delay": b.ConfirmationDelay, "known delays": s.confirmationDelays,
				"block": b,
			})
			continue
		}
		if b.Height+uint64(b.ConfirmationDelay) >= currentHeight {
			s.logger.Error(
				earlyBlockReportBlocks,
				commontypes.LogFields{"block": b, "currentHeight": currentHeight},
			)
			continue
		}
		if remainder := b.Height % uint64(s.period); remainder != 0 {
			s.logger.Error(
				invalidBlockReportBlocks,
				commontypes.LogFields{"block": b, "period": s.period, "remainder": remainder},
			)
			continue
		}
		if _, present := s.blockProofs[b]; !present {

			blockProof, err2 := s.vrfOutput(b, s.keyProvider.KeyLookup(s.keyID))

			if err2 != nil {
				return nil, err2
			}
			s.blockProofs[b] = blockProof
		}
		proofBytes, err3 := s.blockProofs[b].MarshalBinary()
		if err3 != nil {
			s.logger.Warn(failedMarshalVRFProof, commontypes.LogFields{
				"oracleID": s.i, "error": err3,
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
		pcb := protobuf.CostedCallback{
			Callback: &protobuf.Callback{
				RequestId:      c.RequestID,
				NumWords:       uint32(c.NumWords),
				Requester:      append([]byte{}, c.Requester[:]...),
				Arguments:      append([]byte{}, c.Arguments...),
				SubscriptionID: c.SubscriptionID,
				Height:         c.BeaconHeight,
				ConfDelay:      c.ConfirmationDelay,
			},
			Price:        c.Price.Bytes(),
			GasAllowance: c.GasAllowance.Bytes(),
		}
		err2 := sanityCheckCallback(
			&pcb, s.logger, s.i.OracleID(), s.confirmationDelays, s.period,
		)
		if err2 != nil {
			s.logger.Debug(skipErrMsg, commontypes.LogFields{
				"callback": c,
				"error":    err2,
			})
			continue
		}
		callbacks = append(callbacks, &pcb)
	}

	if (len(outputs) == 0) && (len(callbacks) == 0) {
		s.logger.Error(noValidDataToIncludeInReport, nil)
		return nil, errors.Errorf(noValidDataToIncludeInReport)
	}
	juelsPerFeeCoin, err := s.juelsPerFeeCoin.JuelsPerFeeCoin()
	if err != nil {
		return nil, errors.Wrap(err, failedReadJulesPerFeeCoin)
	}
	if len(juelsPerFeeCoin.Bytes()) > (96 / 8) {
		return nil, errors.Errorf(
			largeFeeCoin+" %d", juelsPerFeeCoin,
		)
	}
	startHeight, blocks, err := s.blockhashes.OnchainVerifiableBlocks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, failedReadVerifiableBlocks)
	}
	lastVerifiableBlockHeight := startHeight + uint64(len(blocks)) - 1
	if lastVerifiableBlockHeight < currentHeight {
		return nil, errors.Errorf(
			currentBlockIsNotInVerifiableBlocks+" %d < %d",
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

	s.logger.Debug(initialObservation, commontypes.LogFields{
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
		return nil, util.WrapError(err, failedMarshalObservation)
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
		err2 := proto.Unmarshal(o.Observation, &observation)
		if err2 != nil {
			s.logger.Warn(failedParseObservation, commontypes.LogFields{
				"oracleID": o.Observer, "observation": o.Observation, "error": err2,
			})
			continue
		}
		s.storeCallbacksByBlocks(
			observation.Callbacks,
			callbacksByBlock,
			callbackCounts,
			callbacks,
			o.Observer,
		)
		if s.n <= uint8(o.Observer) {
			s.logger.Error(
				outOfRangeObserver,
				commontypes.LogFields{"n": s.n, "oracleID": o.Observer},
			)
			continue
		}
		player := players[o.Observer]

		proofs := observation.Proofs
		s.parseVRFProofs(proofs, vrfContributions, o.Observer, player, kd)
		juelsPerFeeCoin := big.NewInt(0).SetBytes(observation.JuelsPerFeeCoin)
		juelsPerFeeCoinObs = append(juelsPerFeeCoinObs, juelsPerFeeCoin)

		type hashes = map[heightHash]struct{}
		seenHashes := make(hashes, len(observation.RecentBlockHashes))

		for _, h := range observation.RecentBlockHashes {
			hh := heightHash{h.Height, common.BytesToHash(h.Blockhash)}
			if _, present := seenHashes[hh]; !present {

				seenHashes[hh] = struct{}{}
				recentBlockHashes[hh]++
			} else {
				fields := commontypes.LogFields{"hash": hh}
				s.logger.Warn(duplicateHashErr, fields)
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
			} else {
				s.logger.Warn(
					notEnoughAppearancesCallback,
					commontypes.LogFields{
						"callback hash": ch, "t": s.t, "count": callbackCounts[ch],
					},
				)
			}
		}
		if len(ccallbacks) == 0 {
			s.logger.Warn(
				noConsensusOnOrphanBlockCallbacksMsg,
				commontypes.LogFields{
					"heightAndDelay": hd, "callbacksHashes": chashes,
				},
			)
			continue
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
			noConsensusOnRecentBlockhash,
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
		return false, nil, err
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
	defer s.reportsLock.Unlock()
	if or, present := s.reports[ts]; present && bytes.Equal(or.s, r) {
		if err := s.coordinator.ReportWillBeTransmitted(ctx, or.r); err != nil {
			return false, util.WrapError(err, "Error in ShouldAcceptFinalizedReport")
		}
		delete(s.reports, ts)
	}
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

const (
	failedConstructObservation             = "could not construct observation"
	failedListPendingBlocks                = "Observation: could not list pending requests"
	noOutputsRequiredNotTransmittingReport = "no VRF outputs required; not transmitting report"
	notEnoughContributions                 = "not enough contributions for block"
	wrongShare                             = "wrong share provided"
	outOfRangeObserver                     = "not enough players for observer index"
	noObservationInRound                   = "no observation required on this round"
	failedReadJulesPerFeeCoin              = "error while reading JulesPerFeeCoin"
	failedReadVerifiableBlocks             = "could not get verifiable blocks"
	failedMarshalObservation               = "Error while marshaling Observation"
	unknownConfirmationDelay               = "unknown confirmation delay"
	earlyBlockReportBlocks                 = "ReportBlocks returned a block too early"
	invalidBlockReportBlocks               = "ReportBlocks returned a non-beacon height"
	failedMarshalVRFProof                  = "could not marshal VRF proof"
	noValidDataToIncludeInReport           = "no valid data to include in report"
	currentBlockIsNotInVerifiableBlocks    = "verifiable blocks don't include current block:"
	largeFeeCoin                           = "fee-coin exchange rate too large:"
	failedReadCurrentHeight                = "could not determine current chain height for confirmation threshold"
	initialObservation                     = "initial observation"
	failedVerifyVRFOutput                  = "could not verify distributed VRF output"
	failedParseObservation                 = "failed to parse observation"
	duplicateHashErr                       = "duplicate hash observed"
	noConsensusOnRecentBlockhash           = "no consensus achieved on most recent block hash"
	notEnoughAppearancesCallback           = "insufficient number of appearances for a callback"
	skipErrMsg                             = "skipping callback due to error"
	noConsensusOnOrphanBlockCallbacksMsg   = "there is no consensus on any of the callbacks of an orphan block"
)
