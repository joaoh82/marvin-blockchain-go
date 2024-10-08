package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	assert.Nil(t, err)

	assert.Equal(t, privateKeySize, len(privateKey.Bytes()))

	publicKey := privateKey.PublicKey()
	assert.Equal(t, PublicKeySize, len(publicKey.Bytes()))
}

func TestGetMnemonicFromEntropy(t *testing.T) {
	entropy := "067f627b554f9f167782f1c8557860b6"
	mnemonic := "all wild paddle pride wheat menu task funny sign profit blouse hockey"

	entroTest, _ := hex.DecodeString(entropy)
	generatedMnemonic, err := GetMnemonicFromEntropy(entroTest)
	assert.Nil(t, err)
	assert.Equal(t, mnemonic, generatedMnemonic)
}

func TestNewPrivateKeyFromMnemonic(t *testing.T) {
	// entropy := "067f627b554f9f167782f1c8557860b6"
	mnemonic := "all wild paddle pride wheat menu task funny sign profit blouse hockey"
	addressString := "e15af3cd7d9c09ebaf20d1f97ea396c218b66037"

	privateKey, err := NewPrivateKeyfromMnemonic(mnemonic)
	assert.Nil(t, err)
	assert.Equal(t, privateKeySize, len(privateKey.Bytes()))

	publicKey := privateKey.PublicKey()
	assert.Equal(t, PublicKeySize, len(publicKey.Bytes()))

	address := publicKey.Address()
	assert.Equal(t, addressString, address.String())
}

func TestNewPrivateKeyFromString(t *testing.T) {
	seed := "753bfa924576a230736e83589933ccb7aad8fd3934d7e9637df4912b58ac95d6"
	addressString := "339f9690596b35d909a8c47fe26c5e8697af034c"
	privateKey, err := NewPrivateKeyfromString(seed)
	assert.Nil(t, err)

	assert.Equal(t, privateKeySize, len(privateKey.Bytes()))

	address := privateKey.PublicKey().Address()
	assert.Equal(t, addressString, address.String())
}

func TestPrivateKeySign(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	assert.Nil(t, err)
	publicKey := privateKey.PublicKey()
	invalidPrivateKey, err := GeneratePrivateKey()
	assert.Nil(t, err)
	invalidPublicKey := invalidPrivateKey.PublicKey()

	data := []byte("hello, world")
	signature, _ := privateKey.Sign(data)

	// Check that the signature is the correct size
	assert.Equal(t, SignatureSize, len(signature.Value))

	// Check that the signature is valid
	assert.True(t, signature.Verify(publicKey, data))

	// Check that the signature is invalid
	assert.False(t, signature.Verify(invalidPublicKey, data))
}

func TestPublicKeyToAddress(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	assert.Nil(t, err)
	publicKey := privateKey.PublicKey()
	address := publicKey.Address()

	// fmt.Println(address)

	assert.Equal(t, addressSize, len(address.Bytes()))
}
