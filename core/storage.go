package core

import (
	"fmt"

	"github.com/joaoh82/marvinblockchain/proto"
)

type Storage interface {
	Put(*proto.Block) error
	Get(string) (*proto.Block, error)
}

type MemoryStore struct {
}

func NewMemorystore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Put(b *proto.Block) error {
	return nil
}

func (s *MemoryStore) Get(hash string) (*proto.Block, error) {
	fmt.Println("Getting block from memory store")
	fmt.Println("Hash:", hash)
	return nil, nil
}
