package model

import (
	"strconv"
	"strings"
)

type Book struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	AuthorID int     `json:"author_id"`
}

func NewBook(name string, price float64, authorID int) *Book {
	return &Book{
		Name:     name,
		Price:    price,
		AuthorID: authorID,
	}
}

type BookFilter struct {
	AuthorIDs []int
	PriceFilter *BookPriceFilter
}

func NewBookFilter() *BookFilter {
	return &BookFilter{}
}

type BookPriceFilter struct {
	Price float64
	Operator string
}

// NewBookPriceFilterFromQuery creates a new BookPriceFilter from a query string, e.g. "<10"
// If the query string is empty or wrong, it returns *BookPriceFilter{0, "="}
func NewBookPriceFilterFromQuery(query string) *BookPriceFilter {
	fields := strings.Fields(query)

	var price float64
	var operator string = "="

	if len(fields) == 1 {
		price, _ = strconv.ParseFloat(fields[0], 64)
	} else if len(fields) == 2 {
		price, _ = strconv.ParseFloat(fields[1], 64)
		operator = fields[0]
	}

	return &BookPriceFilter{price, operator}
}