package main

import (
	"encoding/hex"
	"fmt"

	"github.com/joaoh82/marvinblockchain/core"
	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	fmt.Println("Marvin Blockchain")

	BlockSerialization()
}

func BlockSerialization() {
	fmt.Println("Block Serialization")

	mnemonicTo := "all wild paddle pride wheat menu task funny sign profit blouse hockey"
	// Generate a new private key from the mnemonic
	privateKeyTo := crypto.NewPrivateKeyfromMnemonic(mnemonicTo)
	// fmt.Println("private key TO:", privateKeyTo)
	publicKeyTo := privateKeyTo.PublicKey()
	// fmt.Println("public key TO:", publicKeyTo)
	addressTo := publicKeyTo.Address()
	fmt.Println("address TO:", addressTo)

	mnemonicFrom := "hello wild paddle pride wheat menu task funny sign profit blouse hockey"
	// Generate a new private key from the mnemonic
	privateKeyFrom := crypto.NewPrivateKeyfromMnemonic(mnemonicFrom)
	// fmt.Println("private key FROM:", privateKeyFrom)
	publicKeyFrom := privateKeyFrom.PublicKey()
	// fmt.Println("public key FROM:", publicKeyFrom)
	addressFrom := publicKeyFrom.Address()
	fmt.Println("address FROM:", addressFrom)

	header := &proto.Header{
		PrevBlockHash: make([]byte, 32),
		TxHash:        make([]byte, 32),
		Version:       1,
		Height:        1,
		Timestamp:     1627483623,
		Nonce:         12345,
		Difficulty:    10,
	}

	block := &proto.Block{
		Header:       header,
		Transactions: []*proto.Transaction{},
		PublicKey:    publicKeyFrom.Bytes(),
		Signature:    []byte{},
		Hash:         []byte{},
	}

	tx := &proto.Transaction{
		From:      publicKeyFrom.Bytes(),
		To:        publicKeyTo.Bytes(),
		Value:     1000,
		Data:      []byte("Transaction data"),
		Signature: make([]byte, 64),
		Nonce:     123,
		Hash:      make([]byte, 32),
	}
	core.SignTransaction(&privateKeyFrom, tx)
	core.AddTransaction(block, tx)

	core.SignBlock(&privateKeyFrom, block)

	bBlock, _ := core.SerializeBlock(block)
	fmt.Println("GO: Block WITH TRANSACTIONS hex:", hex.EncodeToString(bBlock))

}

func GenerateKeyPairFullCycle() {
	// Generate a random entropy (128 bits for 12 words or 160 bits for 15 words, 256 bits for 24 words)
	entropy, err := bip39.NewEntropy(128) // 128 bits for 12 words, 160 bits for 15 words, 256 bits for 24 words
	if err != nil {
		fmt.Println("Error generating entropy:", err)
		return
	}
	mnemonic := crypto.GetMnemonicFromEntropy(entropy)

	fmt.Println("entropy:", entropy)
	fmt.Println("mnemonic:", mnemonic)

	// Generate a new private key from the mnemonic
	privateKey := crypto.NewPrivateKeyfromMnemonic(mnemonic)
	fmt.Println("private key:", privateKey)
	publicKey := privateKey.PublicKey()
	fmt.Println("public key:", publicKey)
	address := publicKey.Address()
	fmt.Println("address:", address)
}
