package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"github.com/joaoh82/marvinblockchain/types"
)

// Hasher is an interface for hashing objects.
type Hasher[T any] interface {
	Hash(T) types.Hash
}

// BlockHasher is a hasher for blocks.
type BlockHasher struct{}

// Hash returns the hash of a block.
func (BlockHasher) Hash(b *Header) types.Hash {
	h := sha256.Sum256(b.Bytes())
	return types.Hash(h)
}

// TxHasher is a hasher for transactions.
type TxHasher struct{}

// Hash returns the hash of a transaction.
func (TxHasher) Hash(tx *Transaction) types.Hash {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, tx.From.Bytes())
	binary.Write(b, binary.LittleEndian, tx.To.Bytes())
	binary.Write(b, binary.LittleEndian, tx.Value)
	binary.Write(b, binary.LittleEndian, tx.Data)
	binary.Write(b, binary.LittleEndian, tx.Nonce)

	return types.Hash(sha256.Sum256(b.Bytes()))
}
