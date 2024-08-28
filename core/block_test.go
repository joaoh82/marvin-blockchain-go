package core

import (
	"testing"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/stretchr/testify/assert"
)

func TestSerializeDeserializeHeader(t *testing.T) {
	h := &proto.Header{
		PrevBlockHash: []byte("prev"),
		TxHash:        []byte("tx"),
		Version:       1,
		Height:        1,
		Timestamp:     1,
		Nonce:         1,
		Difficulty:    1,
	}

	data, err := SerializeHeader(h)
	if err != nil {
		t.Fatal(err)
	}

	h2, err := DeserializeHeader(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, h.PrevBlockHash, h2.PrevBlockHash)
	assert.Equal(t, h.TxHash, h2.TxHash)
	assert.Equal(t, h.Version, h2.Version)
	assert.Equal(t, h.Height, h2.Height)
	assert.Equal(t, h.Timestamp, h2.Timestamp)
	assert.Equal(t, h.Nonce, h2.Nonce)
	assert.Equal(t, h.Difficulty, h2.Difficulty)
}

func TestSerializeDeserializeBlock(t *testing.T) {
	b := &proto.Block{
		Header: &proto.Header{
			PrevBlockHash: []byte("prev"),
			TxHash:        []byte("tx"),
			Version:       1,
			Height:        1,
			Timestamp:     1,
			Nonce:         1,
			Difficulty:    1,
		},
		Transactions: []*proto.Transaction{
			{
				From:      []byte("from"),
				To:        []byte("to"),
				Value:     1,
				Data:      []byte("data"),
				Signature: []byte("sig"),
				Nonce:     1,
				Hash:      []byte("hash"),
			},
		},
	}

	data, err := SerializeBlock(b)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := DeserializeBlock(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, b.Header.PrevBlockHash, b2.Header.PrevBlockHash)
	assert.Equal(t, b.Header.TxHash, b2.Header.TxHash)
	assert.Equal(t, b.Header.Version, b2.Header.Version)
	assert.Equal(t, b.Header.Height, b2.Header.Height)
	assert.Equal(t, b.Header.Timestamp, b2.Header.Timestamp)
	assert.Equal(t, b.Header.Nonce, b2.Header.Nonce)
	assert.Equal(t, b.Header.Difficulty, b2.Header.Difficulty)

	assert.Equal(t, b.Transactions[0].From, b2.Transactions[0].From)
	assert.Equal(t, b.Transactions[0].To, b2.Transactions[0].To)
	assert.Equal(t, b.Transactions[0].Value, b2.Transactions[0].Value)
	assert.Equal(t, b.Transactions[0].Data, b2.Transactions[0].Data)
}

func TestSignBlockV2(t *testing.T) {
	mnemonic := "all wild paddle pride wheat menu task funny sign profit blouse hockey"
	addressString := "e15af3cd7d9c09ebaf20d1f97ea396c218b66037"

	privateKey := crypto.NewPrivateKeyfromMnemonic(mnemonic)
	publicKey := privateKey.PublicKey()
	address := publicKey.Address()
	assert.Equal(t, addressString, address.String())

	b := &proto.Block{
		Header: &proto.Header{
			PrevBlockHash: []byte("prev"),
			TxHash:        []byte("tx"),
			Version:       1,
			Height:        100,
			Timestamp:     1724695016265493000,
			Nonce:         1,
			Difficulty:    1,
		},
		Transactions: []*proto.Transaction{
			{
				From:      []byte("from"),
				To:        []byte("to"),
				Value:     1,
				Data:      []byte("data"),
				Signature: []byte("sig"),
				Nonce:     1,
				Hash:      []byte("hash"),
			},
		},
	}

	sig, err := SignBlock(&privateKey, b)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, sig)
}

func TestVerifyBlockV2(t *testing.T) {
	b := GenerateRandomBlockV2(t, 100, []byte("prev"))
	isValid, err := VerifyBlock(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isValid)

	// Test with invalid keypair
	invalidPrivateKey := crypto.GeneratePrivateKey()
	invalidPublicKey := invalidPrivateKey.PublicKey()
	b.PublicKey = invalidPublicKey.Bytes()
	isValid, err = VerifyBlock(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.False(t, isValid)

}

// GenerateRandomBlock generates a random block for testing purposes
func GenerateRandomBlockV2(t *testing.T, height uint64, prevBlockHash []byte) *proto.Block {
	mnemonic := "all wild paddle pride wheat menu task funny sign profit blouse hockey"
	addressString := "e15af3cd7d9c09ebaf20d1f97ea396c218b66037"

	privateKey := crypto.NewPrivateKeyfromMnemonic(mnemonic)
	publicKey := privateKey.PublicKey()
	address := publicKey.Address()
	assert.Equal(t, addressString, address.String())

	toPrivKey := crypto.GeneratePrivateKey()

	b := &proto.Block{
		Header: &proto.Header{
			PrevBlockHash: []byte("prev"),
			TxHash:        []byte("tx"),
			Version:       1,
			Height:        100,
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
	SignTransaction(&privateKey, tx)
	AddTransaction(b, tx)

	txHash, err := CalculateTxHashV2(b.Transactions)
	assert.Nil(t, err)
	b.Header.TxHash = txHash
	sig, err := SignBlock(&privateKey, b)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, sig)

	return b
}
