package pgSQL

import (
	"context"
	"github.com/jackc/pgx/v4"
	"nats-listener/configs"
)

func NewPgConnect(config *configs.DataBaseConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.ConnLink)
	if err != nil {
		return conn, err
	}

	return conn, nil
}
