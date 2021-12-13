package handler

import (
	"bootcamp-homework/model"
	"bootcamp-homework/service"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// GetAllBooks returns a list of books
// if authors query parameter with format '1,2,3' is specified, it will return books filtered by authors
// if price query parameter with format '<operator> <price>' is specified, it will return books filtered by price
func (handler *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
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

	books := handler.service.GetBooksByFilter(bookFilter)

	response := &JSONResponse{Data: books}

	JSONWriter(w, response)
}

func (handler *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &model.Book{}

	JSONReader(w, r.Body, book)

	book, err := handler.service.CreateBook(book)
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

	book, err := handler.service.GetBookByID(id)

	if errors.Is(err, model.ErrBookNotFound) {
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

	err = handler.service.UpdateBook(id, book)

	if errors.Is(err, model.ErrBookNotFound) {
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

	book, err := handler.service.DeleteBook(id)

	if errors.Is(err, model.ErrBookNotFound) {
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

func (handler *BookHandler) GetBooksOfAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	books, err := handler.service.GetBooksByAuthorID(id)

	if errors.Is(err, model.ErrAuthorNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = books

	JSONWriter(w, response)
}

func (handler *BookHandler) GetAuthorOfBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	author, err := handler.service.GetAuthorOfBook(id)

	if errors.Is(err, model.ErrBookNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = author

	JSONWriter(w, response)
}
