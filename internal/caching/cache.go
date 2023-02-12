package caching

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"nats-listener/internal/domain"
	"sync"
)

const getDBLines = `SELECT order_uid, order_data FROM orders`

type Cache struct {
	rw         sync.RWMutex
	lastItemId int
	items      map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		lastItemId: 0,
		items:      map[string]interface{}{},
	}
}

func (c *Cache) FillCache(conn *pgx.Conn) error {
	var orders = map[string]interface{}{}

	lines, err := conn.Query(context.Background(), getDBLines)
	if err != nil {
		return err
	}

	for lines.Next() {
		var id, jsonOrder string
		var order domain.Order

		if err := lines.Scan(&id, &jsonOrder); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(jsonOrder), &order); err != nil {
			return err
		}
		orders[id] = order
	}
	c.items = orders
	return err
}

func (c *Cache) SetItem(key string, value interface{}) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.items[key] = value
}

func (c *Cache) GetItem(key string) (interface{}, bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	return item, true
}
