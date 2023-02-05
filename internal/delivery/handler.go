package delivery

import (
	"encoding/json"
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
func (h *Handler) HandleNewOrder(order string) {
	h.service.AddNewOrder(order)
}

func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	if orderId == "" {
		w.WriteHeader(404)
		return
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	order, err := h.service.GetOrderById(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	result, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(result)
}
