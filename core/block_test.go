package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/types"
	"github.com/stretchr/testify/assert"
)

// TestSignBlock tests signing a block
func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := GenerateRandomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(*privKey))
	assert.NotNil(t, b.Signature)
}

// TestVerifyBlock tests verifying a block
func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := GenerateRandomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(*privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.PublicKey = *otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

// TestEncodeDecodeBlock tests encoding and decoding a block
func TestEncodeDecodeBlock(t *testing.T) {
	b := GenerateRandomBlock(t, 1, types.Hash{})
	buf := &bytes.Buffer{}
	assert.Nil(t, b.Encode(NewBlockEncoder(buf)))

	bDecode := new(Block)
	assert.Nil(t, bDecode.Decode(NewBlockDecoder(buf)))

	assert.Equal(t, b.Header, bDecode.Header)

	for i := 0; i < len(b.Transactions); i++ {
		b.Transactions[i].hash = types.Hash{}
		assert.Equal(t, b.Transactions[i], bDecode.Transactions[i])
	}

	assert.Equal(t, b.PublicKey, bDecode.PublicKey)
	assert.Equal(t, b.Signature, bDecode.Signature)
}

// GenerateRandomBlock generates a random block for testing purposes
func GenerateRandomBlock(t *testing.T, height uint64, prevBlockHash types.Hash) *Block {
	privateKey := crypto.GeneratePrivateKey()
	// tx := GenerateRandomTxWithSignature(t)

	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	b, err := NewBlock(header, []*Transaction{})
	assert.Nil(t, err)
	txHash, err := CalculateTxHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.TxHash = txHash
	assert.Nil(t, b.Sign(*privateKey))

	return b
}
