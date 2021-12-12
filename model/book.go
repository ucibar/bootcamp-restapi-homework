package model

type Book struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	AuthorID int     `json:"authorID"`
}

func NewBook(name string, price float64, authorID int) *Book {
	return &Book{
		Name:     name,
		Price:    price,
		AuthorID: authorID,
	}
}
