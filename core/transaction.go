package core

import (
	"errors"

	"math/rand"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/types"
)

var (
	ErrorInvalidSignature = errors.New("invalid transaction signature")
)

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	From      crypto.PublicKey  // Public key of the sender
	To        crypto.PublicKey  // Public key of the receiver
	Value     uint64            // Amount to transfer
	Data      []byte            // Arbitrary data
	Signature *crypto.Signature // Signature of the transaction
	Nonce     int64             // Nonce of the transaction

	// hash of the transaction
	hash types.Hash
}

// NewTransaction creates a new transaction.
func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data:  data,
		Nonce: rand.Int63n(1000000000000000),
	}
}

// Hash calculates the hash of the transaction.
func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}

	return tx.hash
}

// Sign signs the transaction with the given private key.
func (tx *Transaction) Sign(privateKey *crypto.PrivateKey) error {
	hash := tx.Hash(TxHasher{})
	signature, err := privateKey.Sign(hash.Bytes())
	if err != nil {
		return err
	}

	tx.Signature = signature
	tx.From = *privateKey.PublicKey()

	return nil
}

// Verify verifies the signature of the transaction.
func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return errors.New("missing signature")
	}

	hash := tx.Hash(TxHasher{})
	if !tx.Signature.Verify(&tx.From, hash.Bytes()) {
		return ErrorInvalidSignature
	}

	return nil
}

// Decode decodes the transaction from the given decoder.
func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

// Encode encodes the transaction with the given encoder.
func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}
