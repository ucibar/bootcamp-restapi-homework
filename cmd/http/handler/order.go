package handler

import (
	"bootcamp-homework/model"
	"bootcamp-homework/repository"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type OrderRepository interface {
	All() []*model.Order
	Read(id int) (*model.Order, error)
	Create(order *model.Order) (*model.Order, error)
}

type OrderHandler struct {
	repository OrderRepository
}

func NewOrderHandler(repository OrderRepository) *OrderHandler {
	return &OrderHandler{repository: repository}
}

func (handler OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders := handler.repository.All()

	response := &JSONResponse{Data: orders}

	JSONWriter(w, response)
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := &model.Order{}

	JSONReader(w, r.Body, order)

	order.CreatedAt = time.Now()

	order, err := handler.repository.Create(order)
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

	order, err := handler.repository.Read(id)

	if errors.Is(err, repository.ErrOrderNotFound) {
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
