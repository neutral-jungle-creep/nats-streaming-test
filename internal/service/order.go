package service

import (
	"nats-listener/internal/domain"
	"nats-listener/internal/storage"
)

type OrderService struct {
	storage *storage.Storage
}

func NewOrderService(storage *storage.Storage) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) GetOrderById(id string) (interface{}, error) {
	return s.storage.GetOrderFromCache(id)
}

func (s *OrderService) AddNewOrder(order *domain.Order) error {
	if err := s.storage.AddOrderToDB(order); err != nil {
		return err
	}

	s.storage.AddOrderToCache(order)
	return nil
}
