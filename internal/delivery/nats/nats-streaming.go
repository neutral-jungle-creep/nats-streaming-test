package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"nats-listener/configs"
)

type NatsQueue struct {
	conn stan.Conn
	subj string
}

func NewNatsQueue(config *configs.NatsConfig) (*NatsQueue, error) {
	conn, err := nats.Connect(config.URL)
	if err != nil {
		return nil, err
	}

	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsConn(conn))
	if err != nil {
		return nil, err
	}

	return &NatsQueue{
		conn: sc,
		subj: config.Subj,
	}, nil
}

type OrdersMessageQueue interface {
	OnNewMessage(listener func(order string)) (e error)
	Close()
}

func (nq *NatsQueue) OnNewMessage(handler func(order string)) (e error) {
	nq.conn.Subscribe(nq.subj, func(msg *stan.Msg) {
		handler(string(msg.Data))
	})
	return nil
}

func (nq *NatsQueue) Close() {
	nq.conn.Close()
}
