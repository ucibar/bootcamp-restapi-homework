package model

import "time"

type Order struct {
	ID int
	Items []*OrderItem `json:"items"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	BookID int `json:"book_id"`
	Quantity int `json:"quantity"`
}
