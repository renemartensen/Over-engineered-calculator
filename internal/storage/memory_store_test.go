package storage

import (
	"sync"
	"testing"
)

func TestMemoryStore_AddAndGetAll(t *testing.T) {
	store := &MemoryStore{}

	// Add items
	store.Add("2+3", 5)
	store.Add("4*5", 20)

	history := store.GetAll()
	if len(history) != 2 {
		t.Fatalf("expected 2 items, got %d", len(history))
	}

	if history[0].Expression != "2+3" || history[0].Result != 5 {
		t.Errorf("first item mismatch, got %+v", history[0])
	}

	if history[1].Expression != "4*5" || history[1].Result != 20 {
		t.Errorf("second item mismatch, got %+v", history[1])
	}
}

func TestMemoryStore_GetAllReturnsCopy(t *testing.T) {
	store := &MemoryStore{}
	store.Add("1+1", 2)

	history := store.GetAll()
	history[0].Result = 999 // modify the copy

	history2 := store.GetAll()
	if history2[0].Result != 2 {
		t.Errorf("original store was modified by changing returned slice")
	}
}

func TestMemoryStore_ConcurrentAccess(t *testing.T) {
	store := &MemoryStore{}
	var wg sync.WaitGroup

	f := func(n int) {
		defer wg.Done()
		store.Add("calc", float64(n))
	}

	// Spawn multiple goroutines adding items
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go f(i)
	}

	wg.Wait()

	history := store.GetAll()
	if len(history) != 100 {
		t.Errorf("expected 100 items, got %d", len(history))
	}
}
