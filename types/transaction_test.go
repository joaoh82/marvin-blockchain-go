package types

import (
	"testing"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransactionV2(t *testing.T) {
	fromPrivKey, err := crypto.GeneratePrivateKey()
	assert.Nil(t, err)
	toPrivKey, err := crypto.GeneratePrivateKey()
	assert.Nil(t, err)

	tx := &proto.Transaction{
		From:  fromPrivKey.PublicKey().Bytes(),
		To:    toPrivKey.PublicKey().Bytes(),
		Value: 42,
		Data:  []byte("data"),
		Nonce: 1,
	}

	err = SignTransaction(fromPrivKey, tx)
	assert.Nil(t, err)
	assert.NotNil(t, tx.Signature)

	isValid, err := VerifyTransaction(tx)
	assert.Nil(t, err)
	assert.True(t, isValid)
}

func TestSignTransactionV2InvalidSignature(t *testing.T) {
	fromPrivKey, err := crypto.GeneratePrivateKey()
	assert.Nil(t, err)
	toPrivKey, err := crypto.GeneratePrivateKey()
	assert.Nil(t, err)

	tx := &proto.Transaction{
		From:  fromPrivKey.PublicKey().Bytes(),
		To:    toPrivKey.PublicKey().Bytes(),
		Value: 42,
		Data:  []byte("data"),
		Nonce: 1,
	}

	err = SignTransaction(fromPrivKey, tx)
	assert.Nil(t, err)
	assert.NotNil(t, tx.Signature)

	tx.Value = 43

	isValid, err := VerifyTransaction(tx)
	assert.Nil(t, err)
	assert.False(t, isValid)
}
