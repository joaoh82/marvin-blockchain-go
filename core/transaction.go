package core

import (
	"crypto/sha256"
	"errors"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func SerializeTransaction(tx *proto.Transaction) ([]byte, error) {
	data, err := pb.Marshal(tx)
	if err != nil {
		return nil, errors.New("failed to marshal header")
	}

	return data, nil
}

func DeserializeTransaction(data []byte) (*proto.Transaction, error) {
	tx := &proto.Transaction{}
	if err := pb.Unmarshal(data, tx); err != nil {
		return nil, errors.New("failed to unmarshal header")
	}

	return tx, nil
}

// SignTransaction signs a transaction with a private key.
func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) error {
	hash, err := HashTransaction(tx)
	if err != nil {
		return err
	}

	sig, err := pk.Sign(hash)
	if err != nil {
		return err
	}

	tx.Signature = sig.Bytes()
	tx.Hash = hash
	// tx.From = pk.PublicKey().Bytes()

	return nil
}

// HashTransaction hashes a transaction.
func HashTransaction(tx *proto.Transaction) ([]byte, error) {
	b, err := pb.Marshal(tx)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)

	tx.Hash = hash[:]

	return hash[:], nil
}

// VerifyTransaction verifies the signature of a transaction.
func VerifyTransaction(tx *proto.Transaction) (bool, error) {
	// Temporarily remove the signature to calculate the hash
	tempSig := tx.Signature
	tempHash := tx.Hash
	tx.Signature = nil
	tx.Hash = nil

	// Calculate the hash without the signature and hash
	hash, err := HashTransaction(tx)
	if err != nil {
		return false, err
	}
	// Restore the signature
	tx.Signature = tempSig
	tx.Hash = tempHash

	signature := crypto.SignatureFromBytes(tx.Signature)
	publicKey := crypto.PublicKeyFromBytes(tx.From)
	isValid := signature.Verify(publicKey, hash)

	return isValid, nil
}
