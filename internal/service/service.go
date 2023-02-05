package service

import (
	"nats-listener/internal/domain"
	"nats-listener/internal/storage"
)

type Order interface {
	GetOrderById(id int) (interface{}, error)
	AddNewOrder(order *domain.Order) error
}

type Service struct {
	Order
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Order: NewOrderService(storage),
	}
}
