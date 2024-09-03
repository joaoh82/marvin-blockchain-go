package core

import (
	"testing"

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
		prevHash, err := HashHeader(bc.headers.Last())
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
		prevHash, err := HashHeader(bc.headers.Last())
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
