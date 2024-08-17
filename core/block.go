package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/types"
)

// Header represents the header of a block in the blockchain.
type Header struct {
	PrevBlockHash types.Hash // Hash of the previous block

	TxHash    types.Hash // Hash of the transactions in the block
	Version   uint32     // Version of the block
	Height    uint64     // Height of the block in the blockchain
	Timestamp int64      // Timestamp of the block

	Nonce      uint64 // Nonce used to mine the block
	Difficulty uint8  // Difficulty used to mine the block
}

// Bytes returns the byte representation of the header.
func (h *Header) Bytes() []byte {
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(h)

	return b.Bytes()
}

// String returns a string representation of the header.
func (h *Header) String() string {
	return fmt.Sprintf(`
Header {
	PrevBlockHash: %x,
	TxHash: %x,
	Version: %d,
	Height: %d,
	Timestamp: %d,
	Nonce: %d,
	Difficulty: %d
}`,
		h.PrevBlockHash, h.TxHash, h.Version, h.Height, h.Timestamp, h.Nonce, h.Difficulty)
}

// Block represents a block in the blockchain.
type Block struct {
	*Header

	Transactions []*Transaction
	PublicKey    crypto.PublicKey
	Signature    *crypto.Signature

	// Cached version of the header hash
	hash types.Hash
}

// NewBlock creates a new block with the given header and transactions.
func NewBlock(h *Header, txs []*Transaction) (*Block, error) {
	return &Block{
		Header:       h,
		Transactions: txs,
	}, nil
}

// NewBlockFromPrevHeader creates a new block with the given previous header and transactions.
func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, tx)
	hash, _ := CalculateTxHash(b.Transactions)
	b.TxHash = hash
}

// Sign signs the block with the given private key.
func (b *Block) Sign(privKey crypto.PrivateKey) error {
	signature, err := privKey.Sign(b.Header.Bytes())
	if err != nil {
		return err
	}

	b.PublicKey = *privKey.PublicKey()
	b.Signature = signature

	return nil
}

// Verify verifies the block.
func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(&b.PublicKey, b.Header.Bytes()) {
		return fmt.Errorf("block has invalid signature")
	}

	for _, tx := range b.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	txHash, err := CalculateTxHash(b.Transactions)
	if err != nil {
		return err
	}

	if txHash != b.TxHash {
		return fmt.Errorf("block (%s) has an invalid tx hash", b.Hash(BlockHasher{}))
	}

	return nil
}

// Decode decodes the block from the given decoder.
func (b *Block) Decode(dec Decoder[*Block]) error {
	return dec.Decode(b)
}

// Encode encodes the block with the given encoder.
func (b *Block) Encode(enc Encoder[*Block]) error {
	return enc.Encode(b)
}

// Hash returns the hash of the block.
func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}

	return b.hash
}

// String returns a string representation of the block.
func CalculateTxHash(txs []*Transaction) (types.Hash, error) {
	var hash types.Hash
	buf := &bytes.Buffer{}

	for _, tx := range txs {
		if err := gob.NewEncoder(buf).Encode(tx); err != nil {
			return hash, err
		}
	}

	hash = sha256.Sum256(buf.Bytes())

	return hash, nil
}
