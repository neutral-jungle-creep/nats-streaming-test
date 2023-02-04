package service

import (
	"encoding/json"
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

func (s *OrderService) GetOrderById(id int) (interface{}, error) {
	return s.storage.GetOrderFromCache(id)
}

func (s *OrderService) AddNewOrder(o string) error {
	var order *domain.Order

	if err := s.storage.AddOrderToDB(o); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(o), &order); err != nil {
		return err
	}
	s.storage.AddOrderToCache(order)

	return nil
}
