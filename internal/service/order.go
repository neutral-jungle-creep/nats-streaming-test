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

func (s *OrderService) AddNewOrder(order *domain.Order) error {

	o, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := s.storage.AddOrderToDB(string(o)); err != nil {
		return err
	}

	s.storage.AddOrderToCache(order)

	return nil
}
