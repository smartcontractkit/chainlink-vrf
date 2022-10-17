package vrf

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	kshare "go.dedis.ch/kyber/v3/share"
)

func recoverSignature(
	ps pairing.Suite, pk, msg kyber.Point, partialSignatures []kyber.Point,
	shareIdxs []player_idx.PlayerIdx,
	sharePublicKeys []kyber.Point,
	t, n player_idx.Int,
) (kyber.Point, error) {
	var shares []sigShare
	for pIdx, idx := range shareIdxs {
		shares = append(shares, sigShare{
			idx, partialSignatures[pIdx], sharePublicKeys[pIdx],
		})
	}
	a := aggregatedPartialSigs{pk, msg, shares, ps, t, n}
	if err := a.checkTypes(); err != nil {
		return nil, err
	}
	if err := a.checkUniqueShares(); err != nil {
		return nil, err
	}
	return a.recoverSignature()
}

type aggregatedPartialSigs struct {
	pk     kyber.Point
	msg    kyber.Point
	shares []sigShare
	ps     pairing.Suite
	t, n   player_idx.Int
}

func (a aggregatedPartialSigs) recoverSignature() (kyber.Point, error) {
	var shares []*kshare.PubShare
	for _, s := range a.shares {
		share := s.idx.PubShare(s.sig)
		shares = append(shares, &share)
	}
	rv, err := kshare.RecoverCommit(a.ps.G1(), shares, int(a.t), int(a.n))
	if err != nil {
		return nil, err
	}
	if !validateSignature(a.ps, a.msg, a.pk, rv) {
		return nil, a.badShares(errors.Errorf("invalid signature generated"))
	}
	return rv, nil
}

func (a aggregatedPartialSigs) badShares(err error) error {
	var badShareList []player_idx.PlayerIdx
	for _, share := range a.shares {
		if err2 := share.checkPartialSig(a.ps, a.msg); err2 != nil {
			badShareList = append(badShareList, share.idx)
		}
	}
	if len(badShareList) > 0 {
		return badShares{err, badShareList}
	}
	return err
}

func (a aggregatedPartialSigs) checkUniqueShares() error {
	var seenIdxs map[player_idx.PlayerIdx]bool
	for _, share := range a.shares {
		if seenIdxs[share.idx] {
			return errors.Errorf("duplicated index, %d", share.idx)
		}
	}
	return nil
}

func (a aggregatedPartialSigs) checkTypes() error {
	for _, s := range a.shares {
		if err := s.checkTypes(a.ps); err != nil {
			return err
		}
	}
	if len(a.shares) == 0 {
		return errors.Errorf("no shares provided")
	}
	if reflect.TypeOf(a.msg) != reflect.TypeOf(a.shares[0].sig) {
		return errors.Errorf("wrong type for point to be signed")
	}
	return nil
}

type sigShare struct {
	idx player_idx.PlayerIdx
	sig kyber.Point
	pk  kyber.Point
}

func (s sigShare) checkTypes(ps pairing.Suite) error {
	scalarExample := ps.G1().Scalar()
	scalarType := reflect.TypeOf(scalarExample)
	if reflect.TypeOf(s.idx) != scalarType {
		return errors.Errorf("wrong type for share index: got %T, need %T", s.idx, scalarExample)
	}
	g1Example := ps.G1().Point()
	g1Type := reflect.TypeOf(g1Example)
	if reflect.TypeOf(s.sig) != g1Type {
		return errors.Errorf("wrong type for partial signature: got %T, need %T", s.sig, g1Example)
	}
	g2Example := ps.G2().Point()
	g2Type := reflect.TypeOf(g2Example)
	if reflect.TypeOf(s.pk) != g2Type {
		return errors.Errorf("wrong type for public key: got %T, need %T", s.pk, g2Example)
	}
	return nil
}

func (s sigShare) checkPartialSig(ps pairing.Suite, msg kyber.Point) error {
	if !validateSignature(ps, msg, s.pk, s.sig) {
		return errors.Errorf(
			"bad partial signature %s from base %s, given public key %s",
			s.sig, msg, s.pk,
		)
	}
	return nil
}

type badShares struct {
	err       error
	badShares []player_idx.PlayerIdx
}

var _ error = badShares{}

func (b badShares) Error() string {
	return fmt.Sprintf("%s: PVSS indexes with bad shares: %v", b.err, b.badShares)
}

func validateSignature(p pairing.Suite, msg, pk, sig kyber.Point) bool {
	return p.Pair(msg, pk).Equal(p.Pair(sig, p.G2().Point().Base()))
}
