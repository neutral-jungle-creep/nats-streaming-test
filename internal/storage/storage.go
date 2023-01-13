package storage

import (
	"github.com/jackc/pgx/v4"
)

type Order interface {
}

type Storage struct {
	Order
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Order: NewPgOrderStorage(conn),
	}
}
