package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"nats-listener/pkg/logger"
)

const (
	addLineToDB = `INSERT INTO orders (order_data) VALUES ($1)`
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

func (s *PgOrderStorage) AddOrderToDB(order string) error {
	_, err := s.conn.Exec(context.Background(), addLineToDB, order)

	if err != nil {
		return err
	}
	s.log.Infof("add to data base order: %s", order)
	return nil
}
