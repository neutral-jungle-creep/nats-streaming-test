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

type Storage struct {
	Order
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Order: NewPgOrderStorage(conn),
	}
}
