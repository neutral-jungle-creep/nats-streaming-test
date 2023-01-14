package service

import (
	"sync"
)

type CacheService struct {
	sync.RWMutex
	defaultCapacity int
	amountItems     int
	items           map[int]interface{}
}

func NewCacheService(capacity int) *CacheService {
	var items = map[int]interface{}{}

	cache := CacheService{
		defaultCapacity: capacity,
		amountItems:     0,
		items:           items,
	}

	return &cache
}

func (c *CacheService) SetItem(key int, value interface{}) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}

func (c *CacheService) GetItem(key int) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	return item, true
}

func (c *CacheService) DeleteItem(key int) {
	c.Lock()
	defer c.Unlock()

	_, found := c.items[key]
	if found {
		delete(c.items, key)
	}
}
