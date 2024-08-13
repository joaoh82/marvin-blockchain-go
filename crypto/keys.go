package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"io"

	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/pbkdf2"
)

const (
	privateKeySize = ed25519.PrivateKeySize // 64
	publicKeySize  = ed25519.PublicKeySize  // 32
	signatureSize  = ed25519.SignatureSize  // 64
	seedSize       = 32
	addressSize    = 20
)

// PrivateKey represents a private key for the Ed25519 signature scheme.
type PrivateKey struct {
	key ed25519.PrivateKey
}

func GetMnemonicFromEntropy(entropy []byte) string {
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	return mnemonic
}

// NewPrivateKeyfromMenmonic creates a new private key from a given mnemonic.
func NewPrivateKeyfromMnemonic(mnemonic string) PrivateKey {
	seed := SeedFromMnemonic(mnemonic)
	return NewPrivateKeyFromSeed(seed)
}

// SeedFromMnemonic creates a new seed from a given mnemonic.
func SeedFromMnemonic(mnemonic string) []byte {
	seed := pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+"Secret Passphrase"), 2048, 32, sha512.New)
	return seed
}

// NewPrivateKeyfromString creates a new private key from a given hex string.
func NewPrivateKeyfromString(s string) PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

// NewPrivateKeyFromSeed creates a new private key from a given seed byte slice.
func NewPrivateKeyFromSeed(seed []byte) PrivateKey {
	if len(seed) != seedSize {
		panic("invalid seed size, must be 32 bytes")
	}

	return PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

// GeneratePrivateKey generates a new private key.
func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedSize)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

// Bytes creates a new private key from a byte slice.
func (p *PrivateKey) Bytes() []byte {
	return p.key
}

// Sign signs the data with the private key.
func (p *PrivateKey) Sign(data []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, data),
	}
}

// PublicKey returns the public key for the private key.
func (p *PrivateKey) PublicKey() *PublicKey {
	b := make([]byte, publicKeySize)
	copy(b, p.key[32:])
	return &PublicKey{
		key: b,
	}
}

// PublicKey represents a public key for the Ed25519 signature scheme.
type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[:addressSize],
	}
}

// Bytes creates a new public key from a byte slice.
func (p *PublicKey) Bytes() []byte {
	return p.key
}

// Signature represents a signature for the Ed25519 signature scheme.
type Signature struct {
	value []byte
}

// Bytes creates a new signature from a byte slice.
func (s *Signature) Bytes() []byte {
	return s.value
}

// Verify verifies the signature of the data with the public key.
// It returns true if the signature is valid, and false otherwise.
func (s *Signature) Verify(publicKey *PublicKey, data []byte) bool {
	return ed25519.Verify(publicKey.key, data, s.value)
}

// Address represents an address for the Ed25519 signature scheme.
type Address struct {
	value []byte
}

// Bytes creates a new address from a byte slice.
func (a *Address) Bytes() []byte {
	return a.value
}

// String returns the address as a hex string.
func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
