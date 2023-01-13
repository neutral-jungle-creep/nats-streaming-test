package delivery

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"nats-listener/configs"
	"nats-listener/internal/domain"
)

type NatsQueue struct {
	conn stan.Conn
	subj string
}

func NewNatsQueue(config *configs.NatsConfig) (*NatsQueue, error) {
	conn, err := nats.Connect(config.URL)
	// TODO обработать исключение
	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsConn(conn))
	if err != nil {
		return nil, err
	}
	return &NatsQueue{
		conn: sc,
		subj: config.Subj,
	}, nil
}

// TODO переименовать
type OrdersMessageQueue interface {
	OnNewMessage(listener func(option domain.Order))
	Close()
}

func (nq *NatsQueue) OnNewMessage(listener func(o domain.Order)) {
	nq.conn.Subscribe(nq.subj, func(msg *stan.Msg) {
		var order = domain.Order{}

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			//TODO
		}
		//TODO validation
		listener(order)
	})
}

func (nq *NatsQueue) Close() {
	nq.conn.Close()
}
