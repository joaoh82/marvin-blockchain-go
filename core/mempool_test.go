package core

import (
	"testing"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/stretchr/testify/assert"
)

func TestMemPool(t *testing.T) {
	mempool := NewMempool()

	assert.Equal(t, 0, mempool.Len())
}

func TestAddTransaction(t *testing.T) {
	mempool := NewMempool()

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

	err = mempool.Add(tx)
	assert.Nil(t, err)
	assert.Equal(t, 1, mempool.Len())

	mempool.Flush()
	assert.Equal(t, 0, mempool.Len())
}
