package core

import (
	"bytes"
	"testing"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestVerifyTransactionWithTamper(t *testing.T) {
	tx := NewTransaction(nil)

	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()
	hackerPrivKey := crypto.GeneratePrivateKey()

	tx.From = *fromPrivKey.PublicKey()
	tx.To = *toPrivKey.PublicKey()
	tx.Value = 666

	assert.Nil(t, tx.Sign(fromPrivKey))
	tx.hash = types.Hash{}

	tx.To = *hackerPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func TestNativeTransferTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		To:    *toPrivKey.PublicKey(),
		Value: 666,
	}

	assert.Nil(t, tx.Sign(fromPrivKey))
}

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("marvin"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = *otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func TestTxEncodeDecode(t *testing.T) {
	tx := GenerateRandomTxWithSignature(t)
	buf := &bytes.Buffer{}
	assert.Nil(t, tx.Encode(NewTxEncoder(buf)))
	tx.hash = types.Hash{}

	txDecoded := new(Transaction)
	assert.Nil(t, txDecoded.Decode(NewTxDecoder(buf)))
	assert.Equal(t, tx, txDecoded)
}

func GenerateRandomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("marvin"),
	}
	assert.Nil(t, tx.Sign(privKey))

	return &tx
}
