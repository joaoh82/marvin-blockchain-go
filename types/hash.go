package types

import (
	"encoding/hex"
	"fmt"
)

const (
	// HashSize represents the size of a hash in bytes.
	HashSize = 32
)

// Hash represents the 32-byte hash of a block.
type Hash [HashSize]byte

// IsZero returns true if the hash is zero.
func (h Hash) IsZero() bool {
	for _, b := range h {
		if b != 0 {
			return false
		}
	}
	return true
}

// Bytes returns the hash as a byte slice.
func (h Hash) Bytes() []byte {
	b := make([]byte, HashSize)
	for i := 0; i < HashSize; i++ {
		b[i] = h[i]
	}
	return b
}

// String returns the hash as a hex-encoded string.
func (h Hash) String() string {
	return hex.EncodeToString(h.Bytes())
}

// HashFromBytes returns a Hash from a byte slice.
func HashFromBytes(b []byte) Hash {
	if len(b) != HashSize {
		panic(fmt.Sprintf("invalid hash length, should be: %d", len(b)))
	}

	var h Hash
	for i := 0; i < HashSize; i++ {
		h[i] = b[i]
	}
	return h
}
