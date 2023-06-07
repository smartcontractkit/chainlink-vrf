package secp256k1

import (
	"database/sql/driver"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.dedis.ch/kyber/v3"
)

type PublicKey [CompressedPublicKeyLength]byte

const CompressedPublicKeyLength = 33

func init() {
	if CompressedPublicKeyLength != (&Secp256k1{}).Point().MarshalSize() {
		panic("disparity in expected public key lengths")
	}
}

func (k *PublicKey) Set(l PublicKey) {
	if copy(k[:], l[:]) != CompressedPublicKeyLength {
		panic(fmt.Errorf("failed to copy entire public key %x to %x", l, k))
	}
}

func (k *PublicKey) Point() (kyber.Point, error) {
	p := (&Secp256k1{}).Point()
	return p, p.UnmarshalBinary(k[:])
}

func NewPublicKeyFromHex(hex string) (PublicKey, error) {
	rawKey, err := hexutil.Decode(hex)
	if err != nil {
		return PublicKey{}, err
	}
	return NewPublicKeyFromBytes(rawKey)
}

func NewPublicKeyFromBytes(rawKey []byte) (PublicKey, error) {
	if l := len(rawKey); l != CompressedPublicKeyLength {
		return PublicKey{}, fmt.Errorf(
			"wrong length for public key: %s of length %d", rawKey, l)
	}
	var k PublicKey
	if c := copy(k[:], rawKey); c != CompressedPublicKeyLength {
		panic(fmt.Errorf("failed to copy entire key to return value"))
	}
	return k, nil
}

func (k *PublicKey) SetFromHex(hex string) error {
	nk, err := NewPublicKeyFromHex(hex)
	if err != nil {
		return err
	}
	k.Set(nk)
	return nil
}

func (k PublicKey) String() string {
	return hexutil.Encode(k[:])
}

func (k *PublicKey) StringUncompressed() (string, error) {
	p, err := k.Point()
	if err != nil {
		return "", err
	}
	return hexutil.Encode(LongMarshal(p)), nil
}

func (k *PublicKey) Hash() (common.Hash, error) {
	p, err := k.Point()
	if err != nil {
		return common.Hash{}, err
	}
	return MustHash(string(LongMarshal(p))), nil
}

func (k *PublicKey) MustHash() common.Hash {
	hash, err := k.Hash()
	if err != nil {
		panic(fmt.Sprintf("Failed to compute hash of public vrf key %v", k))
	}
	return hash
}

func (k *PublicKey) Address() common.Address {
	hash, err := k.Hash()
	if err != nil {
		return common.Address{}
	}
	return common.BytesToAddress(hash.Bytes()[12:])
}

func (k *PublicKey) IsZero() bool {
	return *k == PublicKey{}
}

func (k PublicKey) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
}

func (k *PublicKey) UnmarshalText(text []byte) error {
	if err := k.SetFromHex(string(text)); err != nil {
		return errors.Wrapf(err, "while parsing %s as public key", text)
	}
	return nil
}

func (k PublicKey) Value() (driver.Value, error) {
	return k.String(), nil
}

func (k *PublicKey) Scan(value interface{}) error {
	rawKey, ok := value.(string)
	if !ok {
		return errors.Wrap(fmt.Errorf("unable to convert %+v of type %T to PublicKey", value, value), "scan failure")
	}
	if err := k.SetFromHex(rawKey); err != nil {
		return errors.Wrapf(err, "while scanning %s as PublicKey", rawKey)
	}
	return nil
}

func MustHash(in string) common.Hash {
	out, err := Keccak256([]byte(in))
	if err != nil {
		panic(err)
	}
	return common.BytesToHash(out)
}

func Keccak256(in []byte) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write(in)
	return hash.Sum(nil), err
}
