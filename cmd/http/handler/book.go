package handler

import (
	"bootcamp-homework/model"
	"bootcamp-homework/repository"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BookRepository interface {
	All() []*model.Book
	Read(id int) (*model.Book, error)
	ReadByAuthor(authorID int) []*model.Book
	Delete(id int) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id int, book *model.Book) error

	Filter(filter *model.BookFilter) []*model.Book
}

type BookHandler struct {
	repository BookRepository
}

func NewBookHandler(repository BookRepository) *BookHandler {
	return &BookHandler{repository: repository}
}

// GetAllBooks returns a list of books
// if authors query parameter with format '1,2,3' is specified, it will return books filtered by authors
// if price query parameter with format '<operator> <price>' is specified, it will return books filtered by price
func (handler BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	bookFilter := model.NewBookFilter()

	if query.Has("authors") {
		authors := strings.Split(query.Get("authors"), ",")
		for _, author := range authors {
			authorID, err := strconv.Atoi(author)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid author id", http.StatusBadRequest)
				return
			}
			bookFilter.AuthorIDs = append(bookFilter.AuthorIDs, authorID)
		}
	}

	if query.Has("price") {
		priceFilter := model.NewBookPriceFilterFromQuery(query.Get("price"))
		bookFilter.PriceFilter = priceFilter
	}

	books := handler.repository.Filter(bookFilter)

	response := &JSONResponse{Data: books}

	JSONWriter(w, response)
}

func (handler *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &model.Book{}

	JSONReader(w, r.Body, book)

	book, err := handler.repository.Create(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := &JSONResponse{Data: book}
	w.WriteHeader(http.StatusCreated)

	JSONWriter(w, response)
}

func (handler *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	book, err := handler.repository.Read(id)

	if errors.Is(err, repository.ErrBookNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = book

	JSONWriter(w, response)
}

func (handler *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	book := &model.Book{}

	JSONReader(w, r.Body, book)

	response := &JSONResponse{}

	err = handler.repository.Update(id, book)

	if errors.Is(err, repository.ErrBookNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = book

	JSONWriter(w, response)
}

func (handler *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	book, err := handler.repository.Delete(id)

	if errors.Is(err, repository.ErrBookNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = book

	JSONWriter(w, response)
}
