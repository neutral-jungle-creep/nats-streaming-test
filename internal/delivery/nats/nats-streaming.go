package nats

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"nats-listener/configs"
	"nats-listener/internal/domain"
)

type NatsQueue struct {
	conn  stan.Conn
	subj  string
	valid *validator.Validate
}

func NewNatsQueue(config *configs.NatsConfig, valid *validator.Validate) (*NatsQueue, error) {
	conn, err := nats.Connect(config.URL)
	if err != nil {
		return nil, err
	}

	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsConn(conn))
	if err != nil {
		return nil, err
	}

	return &NatsQueue{
		conn:  sc,
		subj:  config.Subj,
		valid: valid,
	}, nil
}

type OrdersMessageQueue interface {
	OnNewMessage(listener func(order *domain.Order))
	Close()
}

func (nq *NatsQueue) OnNewMessage(handler func(order *domain.Order)) {
	nq.conn.Subscribe(nq.subj, func(msg *stan.Msg) {
		var order domain.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			return
		}
		//logrus.Infof("полученное сообщение, %v", order.OrderUID)

		if err := nq.valid.Struct(order); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				logrus.Infof("order validate error, %s", err.Error())
				return
			}
		}
		handler(&order)
	})
}

func (nq *NatsQueue) Close() {
	nq.conn.Close()
}
