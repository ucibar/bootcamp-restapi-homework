package repository

import "bootcamp-homework/model"

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
