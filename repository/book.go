package repository

import (
	"bootcamp-homework/model"
)

// InMemoryBookRepository TODO: make goroutine safe
type InMemoryBookRepository struct {
	books map[int]*model.Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{books: make(map[int]*model.Book)}
}

func (repository *InMemoryBookRepository) All() []*model.Book {
	books := make([]*model.Book, len(repository.books))

	index := 0
	for _, book := range repository.books {
		books[index] = book
		index++
	}

	return books
}

func (repository *InMemoryBookRepository) Read(id int) (*model.Book, error) {
	book, ok := repository.books[id]
	if !ok {
		return nil, ErrBookNotFound
	}
	return book, nil
}

func (repository *InMemoryBookRepository) ReadByAuthor(authorID int) []*model.Book {
	books := make([]*model.Book, 0)

	for _, book := range repository.books {
		if book.AuthorID == authorID {
			books = append(books, book)
		}
	}

	return books
}

func (repository *InMemoryBookRepository) Filter(filter *model.BookFilter) []*model.Book {
	books := repository.All()

	if filter.AuthorIDs != nil {
		books = repository.FilterByAuthor(books, filter.AuthorIDs)
	}

	if filter.PriceFilter != nil {
		books = repository.FilterByPrice(books, filter.PriceFilter.Price, filter.PriceFilter.Operator)
	}

	return books
}

func (repository *InMemoryBookRepository) ReadByPrice(price float64, operator string) []*model.Book {
	books := make([]*model.Book, 0)

	for _, book := range repository.books {
		switch operator[0] {
		case '>':
			if book.Price > price {
				books = append(books, book)
			}
		case '<':
			if book.Price < price {
				books = append(books, book)
			}
		case '=':
			if book.Price == price {
				books = append(books, book)
			}
		}
	}

	return books
}

func (repository *InMemoryBookRepository) Create(book *model.Book) (*model.Book, error) {
	book.ID = len(repository.books) + 1
	repository.books[book.ID] = book
	return book, nil
}

func (repository *InMemoryBookRepository) Update(id int, book *model.Book) error {
	_, ok := repository.books[id]
	if !ok {
		return ErrBookNotFound
	}
	book.ID = id
	repository.books[id] = book
	return nil
}

func (repository *InMemoryBookRepository) Delete(id int) (*model.Book, error) {
	book, ok := repository.books[id]
	if !ok {
		return nil, ErrBookNotFound
	}
	delete(repository.books, id)
	return book, nil
}

func (repository *InMemoryBookRepository) FilterByAuthor(collection []*model.Book, authorsIDs []int) []*model.Book {
	books := make([]*model.Book, 0)

	for _, book := range collection {
		for _, authorID := range authorsIDs {
			if book.AuthorID == authorID {
				books = append(books, book)
			}
		}
	}

	return books
}

func (repository *InMemoryBookRepository) FilterByPrice(collection []*model.Book, price float64, operator string) []*model.Book {
	books := make([]*model.Book, 0)

	for _, book := range collection {
		switch operator[0] {
		case '>':
			if book.Price > price {
				books = append(books, book)
			}
		case '<':
			if book.Price < price {
				books = append(books, book)
			}
		case '=':
			if book.Price == price {
				books = append(books, book)
			}
		}
	}

	return books
}
