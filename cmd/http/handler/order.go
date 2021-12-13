package handler

import (
	"bootcamp-homework/model"
	"bootcamp-homework/service"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (handler *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders := handler.service.GetAllOrders()

	response := &JSONResponse{Data: orders}

	JSONWriter(w, response)
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := &model.Order{}

	JSONReader(w, r.Body, order)

	order.CreatedAt = time.Now()

	order, err := handler.service.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := &JSONResponse{Data: order}
	w.WriteHeader(http.StatusCreated)

	JSONWriter(w, response)
}

func (handler *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	order, err := handler.service.GetOrderByID(id)

	if errors.Is(err, model.ErrOrderNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = order

	JSONWriter(w, response)
}
