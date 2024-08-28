package main

import (
	"fmt"
	"log"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/tyler-smith/go-bip39"

	pb "google.golang.org/protobuf/proto"
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
	fmt.Println("private key TO:", privateKeyTo)
	publicKeyTo := privateKeyTo.PublicKey()
	fmt.Println("public key TO:", publicKeyTo)
	addressTo := publicKeyTo.Address()
	fmt.Println("address TO:", addressTo)

	mnemonicFrom := "hello wild paddle pride wheat menu task funny sign profit blouse hockey"
	// Generate a new private key from the mnemonic
	privateKeyFrom := crypto.NewPrivateKeyfromMnemonic(mnemonicFrom)
	fmt.Println("private key FROM:", privateKeyFrom)
	publicKeyFrom := privateKeyFrom.PublicKey()
	fmt.Println("public key FROM:", publicKeyFrom)
	addressFrom := publicKeyFrom.Address()
	fmt.Println("address FROM:", addressFrom)

	block := &proto.Block{
		Header: &proto.Header{
			PrevBlockHash: []byte{},
			TxHash:        []byte{},
			Version:       1,
			Height:        1,
			Timestamp:     1627483623,
			Nonce:         12345,
			Difficulty:    10,
		},
		Transactions: []*proto.Transaction{
			{
				From:      publicKeyTo.Bytes(),
				To:        publicKeyFrom.Bytes(),
				Value:     1000,
				Data:      []byte("Transaction data"),
				Signature: []byte{},
				Nonce:     123,
				Hash:      []byte{},
			},
		},
		PublicKey: publicKeyFrom.Bytes(),
		Signature: []byte{},
		Hash:      []byte{},
	}

	// Serialize to binary format
	data, err := pb.Marshal(block)
	if err != nil {
		log.Fatalf("Failed to marshal block: %v", err)
	}

	fmt.Println("Serialized data:", data)

	// Deserialize from binary format
	var newBlock proto.Block
	err = pb.Unmarshal(data, &newBlock)
	if err != nil {
		log.Fatalf("Failed to unmarshal block: %v", err)
	}

	// Cast proto.Header to core.Header
	header := newBlock.GetHeader()
	fmt.Println("Header:", header)
	fmt.Println("PrevBlockHash:", header.GetPrevBlockHash())
	fmt.Println("TxHash:", header.GetTxHash())
	fmt.Println("Version:", header.GetVersion())
	fmt.Println("Height:", header.GetHeight())
	fmt.Println("Timestamp:", header.GetTimestamp())
	fmt.Println("Nonce:", header.GetNonce())
	fmt.Println("Difficulty:", header.GetDifficulty())

	// Print block public key address\
	publicKey := crypto.PublicKeyFromBytes(newBlock.GetPublicKey())
	fmt.Println("Public key address:", publicKey.Address())

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
