package core

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/joaoh82/marvinblockchain/proto"
)

// TODO: Pass the mnemonic as a parameter in a configuration file or as an environment variable
// chainMnemonic is the mnemonic for the blockchain private key composed of 12 words
const chainMnemonic = "velvet echo quill jungle nimbus crescent whisk anchor harbor tangle mosaic horizon"

type Blockchain struct {
	headers *HeaderList
	store   Storage
}

// NewBlockchain creates a new blockchain
func NewBlockchain(store Storage) *Blockchain {
	bc := &Blockchain{
		headers: NewHeaderList(),
		store:   store,
	}

	// Genesis block - Genesis block is the first block in the blockchain and has the height 0
	genesisBlock, err := createGenesisBlock()
	if err != nil {
		panic(err)
	}
	bc.addBlock(genesisBlock)

	return bc
}

// AddBlock adds a block to the blockchain
func (bc *Blockchain) AddBlock(b *proto.Block) error {
	// Validate the block before adding it to the blockchain
	if err := bc.ValidateBlock(b); err != nil {
		return err
	}
	return bc.addBlock(b)
}

// addBlock adds a block to the blockchain without validation
func (bc *Blockchain) addBlock(b *proto.Block) error {
	bc.headers.Add(b.Header)

	// Log the block added to the blockchain
	log.Info().Fields(map[string]interface{}{
		"height": b.Header.Height,
		"hash":   hex.EncodeToString(b.Hash),
	}).Msg("block added to blockchain")

	// Store the block in the storage
	return bc.store.Put(b)
}

// HasBlock checks if the blockchain has a block at the given height
func (bc *Blockchain) HasBlock(height int) bool {
	return height <= bc.Height()
}

// ValidateBlock checks if the block is valid to be added to the blockchain
func (bc *Blockchain) ValidateBlock(b *proto.Block) error {
	// Check if the blockchain already has the block
	if bc.HasBlock(int(b.Header.GetHeight())) {
		blockHash, _ := HashBlock(b)
		return fmt.Errorf("blockchain already has block at height (%d) with hash (%s)", b.Header.Height, hex.EncodeToString(blockHash))
	}

	// Check if the block height is the next height in the blockchain
	if b.Header.GetHeight() != uint64(bc.Height()+1) {
		return fmt.Errorf("block height %d is not the next height in the blockchain. Current Height: %d", b.Header.GetHeight(), bc.Height())
	}

	// Check if the block is valid
	if ok, err := VerifyBlock(b); err != nil || !ok {
		return fmt.Errorf("block verification failed: %v", err)
	}

	// Retrieve the last header in the blockchain, calculate the hash and compare it with the previous block hash
	lastHeader, err := bc.GetHeaderByHeight(bc.Height())
	if err != nil {
		return err
	}
	lastHeaderHash, err := HashHeader(lastHeader)
	if err != nil {
		return err
	}
	// Check if the previous block hash in the new block is the same as the hash of the last header in the blockchain
	if !bytes.Equal(lastHeaderHash, b.Header.PrevBlockHash) {
		return fmt.Errorf("invalid previous block hash")
	}

	return nil
}

// GetBlockByHash returns the block with the given hash
func (bc *Blockchain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashStr := hex.EncodeToString(hash)
	block, err := bc.store.Get(hashStr)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetBlockByHeight returns the block at the given height
func (bc *Blockchain) GetBlockByHeight(height int) (*proto.Block, error) {
	if !bc.HasBlock(height) {
		return nil, fmt.Errorf("blockchain does not have block at height (%d)", height)
	}
	header := bc.headers.Get(height)
	headerHash, err := HashHeader(header)
	if err != nil {
		return nil, err
	}
	return bc.GetBlockByHash(headerHash)
}

// GetHeaderByHeight returns the header at the given height
func (bc *Blockchain) GetHeaderByHeight(height int) (*proto.Header, error) {
	if !bc.HasBlock(height) {
		return nil, fmt.Errorf("blockchain does not have block at height (%d)", height)
	}
	return bc.headers.Get(height), nil
}

// Height returns the height of the blockchain
func (bc *Blockchain) Height() int {
	// [0, 1, 2 ,3] => 4 len
	// [0, 1, 2 ,3] => 3 height
	return bc.headers.Height()
}

// createGenesisBlock creates the genesis block of the blockchain
func createGenesisBlock() (*proto.Block, error) {
	private_key, err := crypto.NewPrivateKeyfromMnemonic(chainMnemonic)
	if err != nil {
		return nil, err
	}

	header := &proto.Header{
		PrevBlockHash: make([]byte, 32), // Genesis block has no previous block, so the hash is 32 bytes of zeros
		Version:       1,
		Height:        0, // Genesis block height is 0
		Timestamp:     time.Now().UnixNano(),
	}
	// Genesis block has no transactions and only needs the header
	block := &proto.Block{
		Header: header,
	}

	// Sign the block
	SignBlock(&private_key, block)

	return block, nil
}
