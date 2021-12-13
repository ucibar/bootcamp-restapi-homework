package service

import "bootcamp-homework/model"

type OrderRepository interface {
	All() []*model.Order
	Read(orderID int) (*model.Order, error)
	Create(orderID *model.Order) (*model.Order, error)
}

type OrderService struct {
	orderRepo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) GetAllOrders() []*model.Order {
	return s.orderRepo.All()
}

func (s *OrderService) CreateOrder(order *model.Order) (*model.Order, error) {
	return s.orderRepo.Create(order)
}

func (s *OrderService) GetOrderByID(orderID int) (*model.Order, error) {
	return s.orderRepo.Read(orderID)
}
