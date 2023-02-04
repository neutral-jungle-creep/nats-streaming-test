package storage

import (
	"github.com/jackc/pgx/v4"
	"nats-listener/internal/domain"
)

type Order interface {
	GetDBLines() (*[]domain.Order, error)
	GetLineFromId(id int) (*domain.Order, error)
	AddLineToDB(order *domain.Order) error
}

type Cache interface {
	SetItem(key int, value interface{})
	GetItem(key int) (interface{}, bool)
	DeleteItem(key int)
}

type Storage struct {
	Order
	Cache
}

func NewStorage(conn *pgx.Conn, cacheCapacity int) *Storage {
	return &Storage{
		Order: NewPgOrderStorage(conn),
		Cache: NewCacheStorage(cacheCapacity),
	}
}
