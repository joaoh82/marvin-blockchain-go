package core

import (
	"crypto/sha256"
	"errors"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func SerializeHeader(h *proto.Header) ([]byte, error) {
	data, err := pb.Marshal(h)
	if err != nil {
		return nil, errors.New("failed to marshal header")
	}

	return data, nil
}

func DeserializeHeader(data []byte) (*proto.Header, error) {
	h := &proto.Header{}
	if err := pb.Unmarshal(data, h); err != nil {
		return nil, errors.New("failed to unmarshal header")
	}

	return h, nil
}

func SerializeBlock(b *proto.Block) ([]byte, error) {
	data, err := pb.Marshal(b)
	if err != nil {
		return nil, errors.New("failed to marshal block")
	}

	return data, nil
}

func DeserializeBlock(data []byte) (*proto.Block, error) {
	b := &proto.Block{}
	if err := pb.Unmarshal(data, b); err != nil {
		return nil, errors.New("failed to unmarshal block")
	}

	return b, nil
}

func SignBlock(privateKey *crypto.PrivateKey, b *proto.Block) (*crypto.Signature, error) {
	hash, err := HashBlock(b)
	if err != nil {
		return nil, errors.New("failed to hash block")
	}
	signature, err := privateKey.Sign(hash)
	if err != nil {
		return nil, errors.New("failed to sign block")
	}
	b.Signature = signature.Bytes()
	b.PublicKey = privateKey.PublicKey().Bytes()
	b.Hash = hash

	return signature, nil
}

func VerifyBlock(b *proto.Block) (bool, error) {
	// Verify the transactions
	for _, tx := range b.Transactions {
		isValid, err := VerifyTransaction(tx)
		if err != nil {
			return false, err
		}
		if !isValid {
			return false, errors.New("invalid transaction")
		}
	}

	if b.Signature == nil || len(b.Signature) != crypto.SignatureSize {
		return false, errors.New("invalid block signature")
	}

	if b.PublicKey == nil || len(b.PublicKey) != crypto.PublicKeySize {
		return false, errors.New("invalid block public key")
	}

	signature, err := crypto.SignatureFromBytes(b.Signature)
	if err != nil {
		return false, err
	}
	publicKey, err := crypto.PublicKeyFromBytes(b.PublicKey)
	if err != nil {
		return false, err
	}
	hash, err := HashBlock(b)
	if err != nil {
		return false, errors.New("failed to hash block")
	}
	isValid := signature.Verify(publicKey, hash)

	return isValid, nil
}

// HashHeader hashes the header of a block.
func HashHeader(h *proto.Header) ([]byte, error) {
	b, err := pb.Marshal(h)
	if err != nil {
		return nil, errors.New("failed to hash header")
	}

	hash := sha256.Sum256(b)

	return hash[:], nil
}

// HashBlock returns the hash of the Block Header.
func HashBlock(b *proto.Block) ([]byte, error) {
	return HashHeader(b.Header)
}

// AddTransaction adds a transaction to a block.
func AddTransaction(b *proto.Block, tx *proto.Transaction) error {
	b.Transactions = append(b.Transactions, tx)
	hash, err := CalculateTxHash(b.Transactions)
	if err != nil {
		return err
	}
	b.Header.TxHash = hash

	return nil
}

// CalculateTxHash calculates the hash of the transactions in a block.
func CalculateTxHash(txs []*proto.Transaction) ([]byte, error) {
	hasher := sha256.New()
	for _, tx := range txs {
		hash, err := HashTransaction(tx)
		if err != nil {
			return nil, err
		}
		hasher.Write(hash)
	}

	return hasher.Sum(nil), nil
}
