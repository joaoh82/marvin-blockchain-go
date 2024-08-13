package main

import (
	"fmt"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	fmt.Println("Marvin Blockchain")

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
