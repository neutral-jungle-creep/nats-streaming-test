package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"nats-listener/internal/domain"
)

const (
	getDBLines    = `SELECT * FROM orders ORDER BY id DESC LIMIT 5`
	getLineFromId = `SELECT * FROM orders WHERE id=$1`
)

type PgOrderStorage struct {
	conn *pgx.Conn
}

func NewPgOrderStorage(conn *pgx.Conn) *PgOrderStorage {
	return &PgOrderStorage{
		conn: conn,
	}
}

func (s *PgOrderStorage) GetDBLines() (*[]domain.Order, error) {
	var orders []domain.Order
	lines, err := s.conn.Query(context.Background(), getDBLines)
	if err != nil {
		return nil, err
	}

	for lines.Next() {
		var order domain.Order

		err := lines.Scan(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func (s *PgOrderStorage) GetLineFromId(id int) (*domain.Order, error) {
	var order domain.Order

	line := s.conn.QueryRow(context.Background(), getLineFromId, id)

	err := line.Scan(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
