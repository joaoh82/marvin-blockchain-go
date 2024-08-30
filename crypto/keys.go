package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
)

const (
	privateKeySize = ed25519.PrivateKeySize // 64
	PublicKeySize  = ed25519.PublicKeySize  // 32
	SignatureSize  = ed25519.SignatureSize  // 64
	seedSize       = 32
	addressSize    = 20
)

// PrivateKey represents a private key for the Ed25519 signature scheme.
type PrivateKey struct {
	Key ed25519.PrivateKey
}

// GetMnemonicFromEntropy generates a new mnemonic from a given entropy.
func GetMnemonicFromEntropy(entropy []byte) (string, error) {
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// NewPrivateKeyfromMenmonic creates a new private key from a given mnemonic.
func NewPrivateKeyfromMnemonic(mnemonic string) (PrivateKey, error) {
	seed := SeedFromMnemonic(mnemonic)
	pk, err := NewPrivateKeyFromSeed(seed)
	if err != nil {
		return PrivateKey{}, err
	}
	return pk, nil
}

// SeedFromMnemonic creates a new seed from a given mnemonic.
func SeedFromMnemonic(mnemonic string) []byte {
	seed := pbkdf2.Key([]byte(mnemonic), []byte("mnemonicSecret Passphrase"), 1024, 32, sha512.New)
	return seed
}

// NewPrivateKeyfromString creates a new private key from a given hex string.
func NewPrivateKeyfromString(s string) (PrivateKey, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return PrivateKey{}, err
	}
	return NewPrivateKeyFromSeed(b)
}

// NewPrivateKeyFromSeed creates a new private key from a given seed byte slice.
func NewPrivateKeyFromSeed(seed []byte) (PrivateKey, error) {
	if len(seed) != seedSize {
		return PrivateKey{}, fmt.Errorf("invalid seed size, must be 32 bytes")
	}

	return PrivateKey{
		Key: ed25519.NewKeyFromSeed(seed),
	}, nil
}

// GeneratePrivateKey generates a new private key.
func GeneratePrivateKey() (*PrivateKey, error) {
	seed := make([]byte, seedSize)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		return nil, err
	}
	return &PrivateKey{
		Key: ed25519.NewKeyFromSeed(seed),
	}, nil
}

// Bytes creates a new private key from a byte slice.
func (p *PrivateKey) Bytes() []byte {
	return p.Key
}

// Sign signs the data with the private key.
func (p *PrivateKey) Sign(data []byte) (*Signature, error) {
	return &Signature{
		Value: ed25519.Sign(p.Key, data),
	}, nil
}

// PublicKey returns the public key for the private key.
func (p *PrivateKey) PublicKey() *PublicKey {
	b := make([]byte, PublicKeySize)
	copy(b, p.Key[32:])
	return &PublicKey{
		Key: b,
	}
}

// String returns the private key as a hex string.
func (p *PrivateKey) String() string {
	return hex.EncodeToString(p.Key)
}

// PublicKey represents a public key for the Ed25519 signature scheme.
type PublicKey struct {
	Key ed25519.PublicKey
}

func PublicKeyFromBytes(b []byte) (*PublicKey, error) {
	if len(b) != PublicKeySize {
		return nil, fmt.Errorf("invalid public key length")
	}
	return &PublicKey{
		Key: ed25519.PublicKey(b),
	}, nil
}

// Address returns the address for the public key.
func (p *PublicKey) Address() Address {
	return Address{
		Value: p.Key[:addressSize],
	}
}

// Bytes creates a new public key from a byte slice.
func (p *PublicKey) Bytes() []byte {
	return p.Key
}

func (p *PublicKey) String() string {
	return hex.EncodeToString(p.Key)
}

// Signature represents a signature for the Ed25519 signature scheme.
type Signature struct {
	Value []byte
}

func SignatureFromBytes(b []byte) (*Signature, error) {
	if len(b) != SignatureSize {
		return nil, fmt.Errorf("invalid signature length")
	}
	return &Signature{
		Value: b,
	}, nil
}

// Bytes creates a new signature from a byte slice.
func (s *Signature) Bytes() []byte {
	return s.Value
}

// String returns the signature as a hex string.
func (s Signature) String() string {
	return hex.EncodeToString(s.Value)
}

// Verify verifies the signature of the data with the public key.
// It returns true if the signature is valid, and false otherwise.
func (s *Signature) Verify(publicKey *PublicKey, data []byte) bool {
	return ed25519.Verify(publicKey.Key, data, s.Value)
}

// Address represents an address for the Ed25519 signature scheme.
type Address struct {
	Value []byte
}

func AddressFromBytes(b []byte) (Address, error) {
	if len(b) != addressSize {
		return Address{}, fmt.Errorf("length of the (address) bytes not equal to 20")
	}
	return Address{
		Value: b,
	}, nil
}

// Bytes creates a new address from a byte slice.
func (a *Address) Bytes() []byte {
	return a.Value
}

// String returns the address as a hex string.
func (a Address) String() string {
	return hex.EncodeToString(a.Value)
}
