package recovery

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/player_idx"
	"github.com/smartcontractkit/ocr2vrf/internal/crypto/point_translation"
	"github.com/smartcontractkit/ocr2vrf/util"
	"github.com/twystd/tweetnacl-go/tweetnacl"

	"github.com/stretchr/testify/require"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/sign/anon"
)

type EncryptedItem struct {
	ephemeralPublicKey []byte
	cipher             []byte
	nonce              []byte
}

var PairingSuite pairing.Suite = &altbn_128.PairingSuite{}
var ParingTranslation = &point_translation.PairingTranslation{
	Suite: &altbn_128.PairingSuite{},
}

type randomStream rand.Rand

func NewStream(t *testing.T, seed int64) *randomStream {
	return (*randomStream)(rand.New(rand.NewSource(seed)))
}

func (s *randomStream) XORKeyStream(dst, src []byte) {
	(*rand.Rand)(s).Read(dst)
}

func TestAccountEnrollment(t *testing.T) {
	r := NewStream(t, 0)

	// We are using the G1 suite.
	suite := (PairingSuite).G1().(anon.Suite)
	_, thresh := 31, 10

	/**
	 * Start: work done by the user.
	 */

	// Create key for public-key encryption.
	recipientKeyPair := util.UserPublicKeyEncryptionKeyPair

	// Create bytes from the encrypting public key. Transmit the public key bytes on-chain.
	pubKeyBytes := recipientKeyPair.PublicKey

	// Make a G1 point that represents an account address. Then, convert it to bytes and transmit it on-chain.
	accountPoint := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853")).HashPoint
	accountPointBytes, err := accountPoint.MarshalBinary()
	require.NoError(t, err)

	/**
	 * Start: work done by the committe.
	 */

	// Construct a private polynomial; we'll pretend it's been split up into secret shares securely via a DKG.
	f := share.NewPriPoly(suite, thresh, nil, r)
	secret := f.Secret()

	// Each member of the committe multiplies the account point by their secret share,
	// then encrypts the result to ciphers.
	var shares []*EncryptedItem
	for i := 0; i < thresh; i++ {
		// Retrieve secret share.
		pri := f.Eval(i + 1)

		// Get the account point from the on-chain bytes.
		recoveredAccountPoint := suite.Point()
		err = recoveredAccountPoint.UnmarshalBinary(accountPointBytes)
		require.NoError(t, err)
		require.Equal(t, accountPoint, recoveredAccountPoint)

		// Construct committment by multiplying the secret share by the account point.
		commitment := suite.Point().Mul(pri.V, recoveredAccountPoint)

		// Create an ephemeral key for public-key encryption.
		ephemeralKey, err := tweetnacl.CryptoBoxKeyPair()
		require.NoError(t, err)

		// Encrypt the resulting point for on-chain transmission.
		cipher, ephemeralPublicKey, nonce := util.GetCipherFromPoint(commitment, pubKeyBytes, ephemeralKey)
		shares = append(shares, &EncryptedItem{
			cipher:             cipher,
			ephemeralPublicKey: ephemeralPublicKey,
			nonce:              nonce,
		})
	}

	/**
	 * At this point, all shares have been safely constructed and encrypted into binary.
	 * These shares could be transmitted on-chain, or between nodes in OCR communication.
	 */

	/**
	 * Start: work done by the user.
	 */

	// Recover the public shares from the encrypted points given by the committee.
	var pubShares []*share.PubShare
	for i := 0; i < thresh; i++ {
		// Retrieve the point from the encrypted cipher.
		shareItem := shares[i]
		point := util.GetPointFromCipher(shareItem.cipher, shareItem.nonce, recipientKeyPair, shareItem.ephemeralPublicKey)

		// Construct a public share and append it.
		s := &share.PubShare{
			I: i + 1,
			V: point,
		}
		pubShares = append(pubShares, s)
	}

	// Construct the secret point, which is the account point multiplied by the secret key.
	commit, err := share.RecoverCommit((&altbn_128.PairingSuite{}).G1(), pubShares, thresh, len(shares))
	require.NoError(t, err)

	// Ensure that the commit was successfully constructed,
	// which is the accountPoint times the distributed secret key.
	require.True(t, commit.Equal(suite.Point().Mul(secret, accountPoint)))
}

func TestFullRecovery(t *testing.T) {
	r := NewStream(t, 0)

	// We are using the G1 suite.
	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	_, thresh := 31, 10

	// Simulate T+1 public encryption keys for the committee.
	// In practice the number being used could be larger than T+1.
	var keys []*tweetnacl.KeyPair
	for i := 0; i < thresh+1; i++ {
		// Create key for public-key encryption.
		key, err := tweetnacl.CryptoBoxKeyPair()
		require.NoError(t, err)
		keys = append(keys, key)
	}

	// Construct private polynomial.
	f := share.NewPriPoly(suite, thresh, nil, r)
	secret := f.Secret()
	distributedPublicKey, err := ParingTranslation.TranslateKey(secret)
	require.NoError(t, err)

	// Make a recovery key for account 1, that is its account, the secret key, and the answer.
	acct1Point := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853"))
	answerHash := crypto.Keccak256([]byte("Robert"))
	answerScalar := suite.Scalar().SetBytes(answerHash)
	accountKey := suite.Point().Mul(answerScalar, suite.Point().Mul(secret, acct1Point.HashPoint))
	_, err = accountKey.MarshalBinary()
	require.NoError(t, err)

	/**
	 * Start: work done by the user.
	 */

	// Craft a recovery attempt for account 2, that is the account, the new one, and an answer.
	answer2Hash := crypto.Keccak256([]byte("Robert"))
	answer2Scalar := suite.Scalar().SetBytes(answer2Hash)
	recoverKey := suite.Point().Mul(answer2Scalar, acct1Point.HashPoint)

	var recoveryKeyShares []*EncryptedItem
	for i := 0; i < thresh+1; i++ {
		ephemeralKeyPair, err := tweetnacl.CryptoBoxKeyPair()
		require.NoError(t, err)

		cipher, pubKeyBytes, nonce := util.GetCipherFromPoint(recoverKey, keys[i].PublicKey, ephemeralKeyPair)

		fmt.Printf("LENGTH OF Ciphers: %d", len(cipher))
		recoveryKeyShares = append(recoveryKeyShares, &EncryptedItem{
			cipher:             cipher,
			ephemeralPublicKey: pubKeyBytes,
			nonce:              nonce,
		})
	}

	/**
	 * Start: work done by the committee.
	 */

	// Simulate a DKG committee giving shares.
	shares := []*share.PubShare{}
	for i := 0; i < thresh+1; i++ {
		// Decrypt the recovery key from the posted shares.
		rShare := recoveryKeyShares[i]
		rKey := util.GetPointFromCipher(rShare.cipher, rShare.nonce, keys[i], rShare.ephemeralPublicKey)
		// Construct a public share and append it.
		pri := f.Eval(i + 1)
		res := suite.Point().Mul(pri.V, rKey)
		shares = append(shares, &share.PubShare{I: i + 1, V: res})
	}

	// Recover the combined account point via lagrange interpolation.
	commit, err := share.RecoverCommit(util.PairingSuite.G1(), shares, thresh+1, len(shares))
	require.NoError(t, err)

	// Verify that the signature is correct with the secret key.
	commitIsValid := commit.Equal(suite.Point().Mul(secret, recoverKey))
	require.True(t, commitIsValid)

	// Validate that the signature is correct with the distributed public key.
	validateSignature(PairingSuite, recoverKey, distributedPublicKey, commit)

	// Verify that the answer is correct.
	correct := commit.Equal(accountKey)
	require.True(t, correct)
}

// Tests the encryption and decryption of a G1 Point.
func TestEncryptDecryptPoint(t *testing.T) {
	// We are using the G1 suite.
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	point := (&altbn_128.PairingSuite{}).G1().Point()

	// Marshal G1 point to binary.
	scalar := suite.Scalar().SetBytes([]byte("random"))
	inputPoint := point.Mul(scalar, point.Base())
	inputPointClone := inputPoint.Clone()

	// Create recipient key pair.
	recipientKeyPair := util.UserPublicKeyEncryptionKeyPair

	// Create ephemeral, sending key pair.
	ephemeralKeyPair, err := tweetnacl.CryptoBoxKeyPair()
	require.NoError(t, err)

	// Create bytes from the public key (and maybe put them on-chain).
	pubKeyBytes := recipientKeyPair.PublicKey

	// Create two encrypted ciphers from the G1 point.
	cipher, ephemeralPublicKey, nonce := util.GetCipherFromPoint(inputPoint, pubKeyBytes, ephemeralKeyPair)

	/**
	 * At this point, these ciphers are in binary form. They can be easily written on-chain
	 * or sent between nodes in OCR.
	 */

	// Reconstruct the point from its ciphers.
	outputPoint := util.GetPointFromCipher(cipher, nonce, recipientKeyPair, ephemeralPublicKey)

	// Ensure point has been reconstrcuted correctly.
	require.True(t, inputPoint.Equal(outputPoint))
	require.True(t, inputPointClone.Equal(outputPoint))
}

func TestTweetNacl(t *testing.T) {

	recipientKeyPair := util.UserPublicKeyEncryptionKeyPair

	ephemeralKeyPair, err := tweetnacl.CryptoBoxKeyPair()
	require.NoError(t, err)

	msg := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853")).HashPoint

	cipher, ephermeralPubKey, nonce := util.GetCipherFromPoint(msg, recipientKeyPair.PublicKey, ephemeralKeyPair)
	require.NoError(t, err)

	point := util.GetPointFromCipher(cipher, nonce, recipientKeyPair, ephermeralPubKey)
	require.True(t, point.Equal(msg))
}

func TestCipherWithString(t *testing.T) {
	r := NewStream(t, 0)
	// create ciphertext
	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	stringScalar := suite.Scalar().SetBytes([]byte("hi there!"))
	n, thresh := 31, 10
	f := share.NewPriPoly(suite, thresh, stringScalar, r)
	_, err := player_idx.PlayerIdxs(player_idx.Int(n))

	shares := []*share.PriShare{}
	for i := 0; i < thresh+1; i++ {
		shares = append(shares, f.Eval(i+1))
	}

	secret, err := share.RecoverSecret((&altbn_128.PairingSuite{}).G1(), shares, thresh, len(shares))
	require.NoError(t, err)

	secretsAreEqual := secret.Equal(stringScalar)
	require.True(t, secretsAreEqual)
	b, err := secret.MarshalBinary()
	require.NoError(t, err)
	fmt.Println(string(b))
}

func TestCipherMakeAccountSecretKeyRaw(t *testing.T) {
	r := NewStream(t, 0)
	// create ciphertext
	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	acctBytes := crypto.Keccak256(common.HexToHash("0xD883a6A1C22fC4AbFE938a5aDF9B2Cc31b1BF18B").Bytes())
	acctScalar := suite.Scalar().SetBytes(acctBytes)
	n, thresh := 31, 10
	f := share.NewPriPoly(suite, thresh, nil, r)
	_, err := player_idx.PlayerIdxs(player_idx.Int(n))

	shares := []*share.PriShare{}
	for i := 0; i < thresh+1; i++ {
		pri := f.Eval(i + 1)
		pri.V = suite.Scalar().Add(pri.V, acctScalar)
		shares = append(shares, pri)
	}

	secret, err := share.RecoverSecret((&altbn_128.PairingSuite{}).G1(), shares, thresh, len(shares))
	require.NoError(t, err)

	secretsAreEqual := secret.Equal(suite.Scalar().Add(f.Secret(), acctScalar))
	require.True(t, secretsAreEqual)
}

func TestCipherMakeAccountSecretKeyWithPoint(t *testing.T) {
	r := NewStream(t, 0)

	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	_, thresh := 31, 10

	f := share.NewPriPoly(suite, thresh, nil, r)

	acctPoint := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853"))
	secret := f.Secret()

	shares := []*share.PubShare{}
	for i := 0; i < thresh+1; i++ {
		pri := f.Eval(i + 1)
		res := suite.Point().Mul(pri.V, acctPoint.HashPoint)
		shares = append(shares, &share.PubShare{I: i + 1, V: res})
	}

	commit, err := share.RecoverCommit((&altbn_128.PairingSuite{}).G1(), shares, thresh, len(shares))
	require.NoError(t, err)

	commitIsValid := commit.Equal(suite.Point().Mul(secret, acctPoint.HashPoint))
	require.True(t, commitIsValid)
}

func TestCipherMakeAccountKeyWithRequesterKey(t *testing.T) {
	r := NewStream(t, 0)

	// initialize suite.
	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	_, thresh := 31, 10

	// construct private polynomial
	f := share.NewPriPoly(suite, thresh, nil, r)
	secret := f.Secret()

	// make account_1 a point, and scalar multiply it by account_2
	acctPoint := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853"))
	acctBytes := crypto.Keccak256(common.HexToHash("0xD883a6A1C22fC4AbFE938a5aDF9B2Cc31b1BF18B").Bytes())
	acctScalar := suite.Scalar().SetBytes(acctBytes)
	combinedAccountPoint := suite.Point().Mul(acctScalar, acctPoint.HashPoint)

	// simulate a DKG committee giving shares
	shares := []*share.PubShare{}
	for i := 0; i < thresh+1; i++ {
		pri := f.Eval(i + 1)
		res := suite.Point().Mul(pri.V, combinedAccountPoint)
		shares = append(shares, &share.PubShare{I: i + 1, V: res})
	}

	// recover the combined account point via lagrange
	commit, err := share.RecoverCommit((&altbn_128.PairingSuite{}).G1(), shares, thresh, len(shares))
	require.NoError(t, err)

	// verify that the retrieved point is (account_1 x the secret key) x account_2
	innerAccountKeyPoint := suite.Point().Mul(secret, acctPoint.HashPoint)
	commitIsValid := commit.Equal(suite.Point().Mul(acctScalar, innerAccountKeyPoint))
	require.True(t, commitIsValid)
}

func TestAccountPublicKeySSS(t *testing.T) {
	r := NewStream(t, 0)
	// create ciphertext
	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)

	// instantiate new secret polynomial, with a public account key as the secret scalar
	acctBytes := common.HexToHash("0xD883a6A1C22fC4AbFE938a5aDF9B2Cc31b1BF18B").Bytes()
	acctScalar := suite.Scalar().SetBytes(acctBytes)
	n, thresh := 31, 10
	f := share.NewPriPoly(suite, thresh, acctScalar, r)
	_, err := player_idx.PlayerIdxs(player_idx.Int(n))

	// append all observations.
	shares := []*share.PriShare{}
	for i := 0; i < thresh+1; i++ {
		pri := f.Eval(i + 1)
		shares = append(shares, pri)
	}

	// recover the secret key using lagrange interpolation
	secret, err := share.RecoverSecret((&altbn_128.PairingSuite{}).G1(), shares, thresh, len(shares))
	require.NoError(t, err)

	// ensure the secret is correct
	secretsAreEqual := secret.Equal(acctScalar)
	require.True(t, secretsAreEqual)

	// convert the secret back into a hash
	secretBytes, err := secret.MarshalBinary()
	secretHash := common.BytesToHash(secretBytes)
	require.NoError(t, err)

	// ensure the hash is the original public key
	secretHashIsCorrect := secretHash == common.HexToHash("0xD883a6A1C22fC4AbFE938a5aDF9B2Cc31b1BF18B")
	require.True(t, secretHashIsCorrect)
}

func TestCreateSecret(t *testing.T) {
	r := NewStream(t, 0)

	// initialize suite.
	domainSep := make([]byte, 32)
	r.XORKeyStream(domainSep, nil)
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)
	_, thresh := 31, 10

	// construct private polynomial
	f := share.NewPriPoly(suite, thresh, nil, r)
	secret := f.Secret()

	// make a recovery key for account 1, that is its account, the secret key, and the answer,
	acct1Point := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853"))
	answerHash := crypto.Keccak256([]byte("Robert"))
	answerScalar := suite.Scalar().SetBytes(answerHash)
	accountKey := suite.Point().Mul(answerScalar, suite.Point().Mul(secret, acct1Point.HashPoint))

	// simulate a DKG committee giving shares
	shares := []*share.PubShare{}
	for i := 0; i < thresh+1; i++ {
		pri := f.Eval(i + 1)
		res := suite.Point().Mul(pri.V, acct1Point.HashPoint)
		shares = append(shares, &share.PubShare{I: i + 1, V: res})
	}

	// recover the combined account point via lagrange
	commit, err := share.RecoverCommit((&altbn_128.PairingSuite{}).G1(), shares, thresh, len(shares))
	require.NoError(t, err)

	// verify that the retrieved point is correctly signed by the secret key
	commitIsValid := accountKey.Equal(suite.Point().Mul(answerScalar, commit))
	require.True(t, commitIsValid)
}

func TestSimplePointMultiplication(t *testing.T) {
	suite := (&altbn_128.PairingSuite{}).G1().(anon.Suite)

	acct1Point := altbn_128.NewHashProof(common.HexToHash("0xeFF41C8725be95e66F6B10489B6bF34b08055853")).HashPoint
	scalar1 := suite.Scalar().SetBytes([]byte{1, 2, 3, 4})
	scalar2 := suite.Scalar().SetBytes([]byte{4, 6, 8, 3})

	a := acct1Point.Mul(scalar1, acct1Point)
	a = a.Mul(scalar2, a)

	b := acct1Point.Mul(scalar2, acct1Point)
	b = a.Mul(scalar1, b)

	require.True(t, a.Equal(b))
}

func validateSignature(p pairing.Suite, msg, pk, sig kyber.Point) bool {
	return p.Pair(msg, pk).Equal(p.Pair(sig, p.G2().Point().Base()))
}
