package delivery

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"nats-listener/internal/domain"
	"nats-listener/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) HandleNewOrder(order *domain.Order) {
	if err := h.service.AddNewOrder(order); err != nil {
		logrus.Info(err.Error())
	}
}

func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if orderId == "" {
		w.WriteHeader(204)
		w.Write([]byte("поле для ввода номера заказа не должно быть пустым"))
		return
	}

	order, err := h.service.GetOrderById(orderId)
	if err != nil {
		w.WriteHeader(204)
		w.Write([]byte("заказ не найден"))
		return
	}

	result, _ := json.Marshal(order)
	w.Write(result)
}
