package core

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/joaoh82/marvinblockchain/types"
)

type Hash []byte

type Mempool struct {
	lock         sync.RWMutex
	transactions map[string]*proto.Transaction
}

func NewMempool() *Mempool {
	return &Mempool{
		transactions: make(map[string]*proto.Transaction),
	}
}

func (m *Mempool) Flush() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.transactions = make(map[string]*proto.Transaction)
}

func (m *Mempool) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.transactions)
}

func (m *Mempool) Has(tx *proto.Transaction) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	hash, err := types.HashTransaction(tx)
	if err != nil {
		return false
	}
	hashStr := hex.EncodeToString(hash)

	_, ok := m.transactions[hashStr]
	return ok
}

func (m *Mempool) Add(tx *proto.Transaction) error {
	if m.Has(tx) {
		return fmt.Errorf("transaction already exists in the mempool")
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	hash, err := types.HashTransaction(tx)
	if err != nil {
		return nil
	}
	hashStr := hex.EncodeToString(hash)

	m.transactions[hashStr] = tx
	return nil
}
