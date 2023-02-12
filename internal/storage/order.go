package storage

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"nats-listener/internal/domain"
	"nats-listener/pkg/logger"
)

const (
	addLineToDB = `INSERT INTO orders (order_uid, order_data) VALUES ($1, $2)`
)

type PgOrderStorage struct {
	conn *pgx.Conn
	log  *logger.Logger
}

func NewPgOrderStorage(conn *pgx.Conn, log *logger.Logger) *PgOrderStorage {
	return &PgOrderStorage{
		conn: conn,
		log:  log,
	}
}

func (s *PgOrderStorage) AddOrderToDB(order *domain.Order) error {
	o, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if _, err := s.conn.Exec(context.Background(), addLineToDB, order.OrderUID, o); err != nil {
		return err
	}
	s.log.Infof("add to data base order: %s", o)
	return nil
}
