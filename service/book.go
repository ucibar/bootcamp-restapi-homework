package service

import "bootcamp-homework/model"

type BookRepository interface {
	All() []*model.Book
	Read(bookID int) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(bookID int, book *model.Book) error
	Delete(bookID int) error

	Filter(filter *model.BookFilter) []*model.Book
}

type BookService struct {
	bookRepo BookRepository
	authorRepo AuthorRepository
}

func NewBookService(bookRepo BookRepository, authorRepo AuthorRepository) *BookService {
	return &BookService{bookRepo: bookRepo, authorRepo: authorRepo}
}

func (service *BookService) GetAllBooks() []*model.Book {
	return service.bookRepo.All()
}

func (service *BookService) GetBooksByFilter(filter *model.BookFilter) []*model.Book {
	return service.bookRepo.Filter(filter)
}

func (service *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	return service.bookRepo.Create(book)
}

func (service *BookService) GetBookByID(bookID int) (*model.Book, error) {
	return service.bookRepo.Read(bookID)
}

func (service *BookService) UpdateBook(bookID int, book *model.Book) error {
	return service.bookRepo.Update(bookID, book)
}

func (service *BookService) DeleteBook(bookID int) (*model.Book, error) {
	book, err := service.bookRepo.Read(bookID)
	if err != nil {
		return nil, err
	}

	err = service.bookRepo.Delete(bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (service *BookService) GetBooksByAuthorID(authorID int) ([]*model.Book, error) {
	author, err := service.authorRepo.Read(authorID)
	if err != nil {
		return nil, err
	}

	return service.bookRepo.Filter(&model.BookFilter{AuthorIDs: []int{author.ID}}), nil
}

func (service *BookService) GetAuthorOfBook(bookID int) (*model.Author, error) {
	book, err := service.bookRepo.Read(bookID)
	if err != nil {
		return nil, err
	}

	return service.authorRepo.Read(book.AuthorID)
}