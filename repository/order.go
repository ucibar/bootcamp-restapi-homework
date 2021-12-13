package repository

import "bootcamp-homework/model"

// InMemoryOrderRepository TODO: make goroutine safe
type InMemoryOrderRepository struct {
	orders map[int]*model.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{orders: make(map[int]*model.Order)}
}

func (repository *InMemoryOrderRepository) All() []*model.Order {
	orders := make([]*model.Order, len(repository.orders))

	index := 0
	for _, order := range repository.orders {
		orders[index] = order
		index++
	}

	return orders
}

func (repository *InMemoryOrderRepository) Read(orderID int) (*model.Order, error) {
	order, ok := repository.orders[orderID]
	if !ok {
		return nil, model.ErrOrderNotFound
	}
	return order, nil
}

func (repository *InMemoryOrderRepository) Create(order *model.Order) (*model.Order, error) {
	order.ID = len(repository.orders) + 1
	repository.orders[order.ID] = order
	return order, nil
}
