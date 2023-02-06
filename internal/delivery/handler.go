package delivery

import (
	"encoding/json"
	"nats-listener/internal/domain"
	"nats-listener/internal/service"
	"net/http"
	"strconv"
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
	h.service.AddNewOrder(order)
}

func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if orderId == "" {
		w.WriteHeader(404)
		w.Write([]byte("поле для ввода номера заказа не должно быть пустым"))
		return
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("заказ не найден"))
		return
	}

	order, err := h.service.GetOrderById(id)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("заказ не найден"))
		return
	}

	result, _ := json.Marshal(order)
	w.Write(result)
}
