package core

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/joaoh82/marvinblockchain/proto"
	"github.com/joaoh82/marvinblockchain/types"
)

// Mempool is a pool of transactions that are not yet included in a block
type Mempool struct {
	lock         sync.RWMutex
	transactions map[string]*proto.Transaction
}

// NewMempool creates a new mempool
func NewMempool() *Mempool {
	return &Mempool{
		transactions: make(map[string]*proto.Transaction),
	}
}

// Flush removes all transactions from the mempool
func (m *Mempool) Flush() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.transactions = make(map[string]*proto.Transaction)
}

// Len returns the number of transactions in the mempool
func (m *Mempool) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.transactions)
}

// Has returns true if the mempool has the transaction
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

// Add adds a transaction to the mempool
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
