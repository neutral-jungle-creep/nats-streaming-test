package service

type Order interface {
}

type Cache interface {
	SetItem(int, interface{})
	GetItem(int) (interface{}, bool)
	DeleteItem(int)
}

type Service struct {
	Order
	Cache
}

func NewService() *Service {
	return &Service{}
}
