package storage

import "sync"

type HistoryItem struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}

type MemoryStore struct {
	mu      sync.RWMutex
	history []HistoryItem
}

func (memoryStore *MemoryStore) Add(expr string, result float64) {
	memoryStore.mu.Lock()
	defer memoryStore.mu.Unlock() // unlocks when functions exits
	memoryStore.history = append(memoryStore.history, HistoryItem{Expression: expr, Result: result})
}

func (memoryStore *MemoryStore) GetAll() []HistoryItem {
	memoryStore.mu.RLock()
	defer memoryStore.mu.RUnlock()
	return append([]HistoryItem(nil), memoryStore.history...)
}

var MemoryStoreInstance = &MemoryStore{}
