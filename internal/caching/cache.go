package caching

import (
	"sync"
)

// TODO переименовать
type CacheItemsController interface {
	SetItem(int, interface{})
	GetItem(int) (interface{}, bool)
	DeleteItem(int)
}

type Cache struct {
	sync.RWMutex
	defaultCapacity int
	amountItems     int
	items           map[int]interface{}
}

func NewCache(capacity int) *Cache {
	var items = map[int]interface{}{}

	cache := Cache{
		defaultCapacity: capacity,
		amountItems:     0,
		items:           items,
	}

	return &cache
}

func (c *Cache) SetItem(key int, value interface{}) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}

func (c *Cache) GetItem(key int) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	return item, true
}

func (c *Cache) DeleteItem(key int) {
	c.Lock()
	defer c.Unlock()

	_, found := c.items[key]
	if found {
		delete(c.items, key)
	}
}
