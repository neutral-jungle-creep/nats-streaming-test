package service

import (
	"nats-listener/internal/storage"
)

type Order interface {
	GetOrderById(id int) (interface{}, error)
	AddNewOrder(order string) error
}

type Service struct {
	Order
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Order: NewOrderService(storage),
	}
}
