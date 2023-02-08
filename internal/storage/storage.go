package storage

import (
	"github.com/jackc/pgx/v4"
	"nats-listener/internal/caching"
	"nats-listener/internal/domain"
	"nats-listener/pkg/logger"
)

type Order interface {
	AddOrderToDB(order *domain.Order) error
}

type Cache interface {
	GetOrderFromCache(id string) (interface{}, error)
	AddOrderToCache(order *domain.Order)
}

type Storage struct {
	Order
	Cache
}

func NewStorage(conn *pgx.Conn, cache *caching.Cache, log *logger.Logger) *Storage {
	return &Storage{
		Order: NewPgOrderStorage(conn, log),
		Cache: NewCacheStorage(cache, log),
	}
}
