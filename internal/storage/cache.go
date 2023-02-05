package storage

import (
	"fmt"
	"nats-listener/internal/caching"
	"nats-listener/internal/domain"
	"nats-listener/pkg/logger"
)

type CacheStorage struct {
	cache *caching.Cache
	log   *logger.Logger
}

func NewCacheStorage(cache *caching.Cache, log *logger.Logger) *CacheStorage {
	return &CacheStorage{
		cache: cache,
		log:   log,
	}
}

func (c *CacheStorage) AddOrderToCache(order *domain.Order) {
	c.cache.SetItem(order)
	c.log.Infof("add to cache order №[%s]", order.OrderUID)
}

func (c *CacheStorage) GetOrderFromCache(id int) (interface{}, error) {
	order, found := c.cache.GetItem(id)
	if !found {
		c.log.Errorf("can't find in cache order [%d]", id)
		return nil, fmt.Errorf("order id=[%d] not found", id)
	}
	return order, nil
}
