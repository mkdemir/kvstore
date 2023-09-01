package kvstorage_test

import (
	"testing"

	"github.com/mkdemir/kvstore/src/internal/storage/memory/kvstorage"
)

func TestSet(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB(map[string]interface{}{})
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	_, err := storage.Set(key, "value")
	if err != nil {
		t.Errorf("Error setting value: %v", err)
	}

	if _, ok := memoryStorage[key]; !ok {
		t.Error("value not equal")
	}
}
