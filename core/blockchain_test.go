package core

import (
	"testing"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/joaoh82/marvinblockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain(NewMemorystore())
	assert.Equal(t, 0, bc.Height())
	assert.True(t, bc.HasBlock(0))
}

func TestHasBlock(t *testing.T) {
	bc := NewBlockchain(NewMemorystore())

	numBlocks := 100
	for i := 0; i < numBlocks; i++ {
		prevHash, err := types.HashHeader(bc.headers.Last())
		assert.NoError(t, err)
		block := GenerateRandomBlock(t, uint64(i+1), prevHash)
		assert.NoError(t, bc.AddBlock(block))
	}

	assert.True(t, bc.HasBlock(50))
	assert.False(t, bc.HasBlock(101))
	assert.True(t, bc.HasBlock(100))
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockchain(NewMemorystore())

	numBlocks := 100
	for i := 0; i < numBlocks; i++ {
		prevHash, err := types.HashHeader(bc.headers.Last())
		assert.NoError(t, err)
		block := GenerateRandomBlock(t, uint64(i+1), prevHash)
		assert.NoError(t, bc.AddBlock(block))
	}
	// Checking if the blockchain has the correct height
	assert.Equal(t, numBlocks, bc.Height())

	// Add a block with the same height
	existingBlock := GenerateRandomBlock(t, 50, []byte("hash"))
	err := bc.AddBlock(existingBlock)
	assert.Error(t, err)
}

// GenerateRandomBlock generates a random block with signature for testing purposes
func GenerateRandomBlock(t *testing.T, height uint64, prevBlockHash []byte) *proto.Block {
	mnemonic := "all wild paddle pride wheat menu task funny sign profit blouse hockey"
	addressString := "e15af3cd7d9c09ebaf20d1f97ea396c218b66037"

	privateKey, err := crypto.NewPrivateKeyfromMnemonic(mnemonic)
	assert.Nil(t, err)
	publicKey := privateKey.PublicKey()
	address := publicKey.Address()
	assert.Equal(t, addressString, address.String())

	toPrivKey, err := crypto.GeneratePrivateKey()
	assert.Nil(t, err)

	b := &proto.Block{
		Header: &proto.Header{
			PrevBlockHash: prevBlockHash,
			TxHash:        []byte("tx"),
			Version:       1,
			Height:        height,
			Timestamp:     1724695016265493000,
			Nonce:         1,
			Difficulty:    1,
		},
	}

	tx := &proto.Transaction{
		From:  privateKey.PublicKey().Bytes(),
		To:    toPrivKey.PublicKey().Bytes(),
		Value: 1,
		Data:  []byte("data"),
		Nonce: 1,
	}
	types.SignTransaction(&privateKey, tx)
	types.AddTransaction(b, tx)

	txHash, err := types.CalculateTxHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.TxHash = txHash
	sig, err := types.SignBlock(&privateKey, b)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, sig)

	return b
}
