package kvstorage

import (
	"fmt"

	"github.com/mkdemir/kvstore/src/internal/kverror"
)

func (ms *memoryStorage) Set(key string, value any) (any, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	if _, err := ms.Get(key); err == nil {
		return nil, fmt.Errorf("%w", kverror.ErrKeyExists.AddData("'"+key+"' already exist"))
	}

	ms.db[key] = value
	return value, nil
}
