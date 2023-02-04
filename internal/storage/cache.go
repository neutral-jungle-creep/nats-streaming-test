package storage

import (
	"sync"
)

type CacheStorage struct {
	rw              sync.RWMutex
	defaultCapacity int
	amountItems     int
	items           map[int]interface{}
}

func NewCacheStorage(capacity int) *CacheStorage {
	return &CacheStorage{
		defaultCapacity: capacity,
		amountItems:     0,
		items:           map[int]interface{}{},
	}
}

func (c *CacheStorage) SetItem(key int, value interface{}) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.items[key] = value
}

func (c *CacheStorage) GetItem(key int) (interface{}, bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	return item, true
}

func (c *CacheStorage) DeleteItem(key int) {
	c.rw.Lock()
	defer c.rw.Unlock()

	_, found := c.items[key]
	if found {
		delete(c.items, key)
	}
}
