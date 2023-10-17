package util

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/twystd/tweetnacl-go/tweetnacl"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/anon"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
)

var accountKey string = os.Getenv("ACCOUNT_KEY")
var ChainIDString string = os.Getenv("ETH_CHAIN_ID")
var ethURL string = os.Getenv("OTHER_ETH_URL")
var DomainSep = []byte{00}

var BeaconAddress string = "0xa34FD46a5f142Ab4A925ECDB099FABB7566Ade74"

var Suite = (&altbn_128.PairingSuite{}).G1().(anon.Suite)
var G1Point = (&altbn_128.PairingSuite{}).G1().Point()
var G2Point = (&altbn_128.PairingSuite{}).G2().Point()
var PairingSuite pairing.Suite = &altbn_128.PairingSuite{}

var UserPublicKeyEncryptionKeyPair = &tweetnacl.KeyPair{
	PublicKey: common.HexToHash("0x8e21e21d7535e0645529678fcdd577ed0b8ad6e9febc0f1f6fdad054f494fb10").Bytes(),
	SecretKey: common.HexToHash("0x51c0fc873712ce35347b4c88c09e2812eea77f50c270cd0211f2a7e8ddc7fa5b").Bytes(),
}

var NodePublicKeyEncryptionKeyPair, _ = tweetnacl.CryptoBoxKeyPair()

func GetKeys() (ecdsa.PrivateKey, common.Address) {
	b, err := hex.DecodeString(accountKey)
	PanicErr(err)
	d := new(big.Int).SetBytes(b)

	pkX, pkY := crypto.S256().ScalarBaseMult(b)
	privateKey := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: crypto.S256(),
			X:     pkX,
			Y:     pkY,
		},
		D: d,
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return privateKey, address
}

func GetCipherFromPoint(
	point kyber.Point,
	recipientPubKey []byte,
	ephemeralKeyPair *tweetnacl.KeyPair,
) ([]byte, []byte, []byte) {
	msg, err := point.MarshalBinary()
	PanicErr(err)

	if len(msg) != 32 {
		panic(errors.New("length of message is not 32 bytes"))
	}

	source := rand.NewSource(time.Now().Unix())
	rand := rand.New(source)
	nonce := []byte{}
	for i := 0; i < 3; i++ {
		nonce = binary.BigEndian.AppendUint64(nonce, rand.Uint64())
	}

	cipher, err := tweetnacl.CryptoBox(msg, nonce, recipientPubKey, ephemeralKeyPair.SecretKey)
	PanicErr(err)

	return cipher, ephemeralKeyPair.PublicKey, nonce
}

func GetPointFromCipher(
	cipher []byte,
	nonce []byte,
	recipientKeyPair *tweetnacl.KeyPair,
	ephemeralKeyPublicKey []byte,
) kyber.Point {
	msg, err := tweetnacl.CryptoBoxOpen(cipher, nonce, ephemeralKeyPublicKey, recipientKeyPair.SecretKey)
	PanicErr(err)

	outputPoint := (&altbn_128.PairingSuite{}).G1().Point()
	err = outputPoint.UnmarshalBinary(msg)
	PanicErr(err)

	return outputPoint
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetOwnerAndClient() (*bind.TransactOpts, *ethclient.Client) {

	ec, err := ethclient.Dial(ethURL)
	PanicErr(err)

	privateKey, _ := GetKeys()

	chainID, err := strconv.Atoi(ChainIDString)
	PanicErr(err)

	owner, err := bind.NewKeyedTransactorWithChainID(&privateKey, big.NewInt(int64(chainID)))
	if false {
		owner.Value = big.NewInt(int64(1e18))
		block, _ := ec.BlockNumber(context.Background())

		nonce, _ := ec.NonceAt(context.Background(), owner.From, big.NewInt(int64(block)))
		owner.Nonce = big.NewInt(int64(nonce))
	}

	gp, err := ec.SuggestGasPrice(context.Background())
	owner.GasPrice = gp.Mul(gp, big.NewInt(2))
	owner.GasLimit = 2_000_000

	return owner, ec
}

func AssertEqual(a interface{}, b interface{}) {
	if a != b {
		panic("assert equal failed")
	}
}

func AssertBytesEqual(a []byte, b []byte) {
	if bytes.Compare(a, b) != 0 {
		panic("assert equal bytes failed")
	}
}
