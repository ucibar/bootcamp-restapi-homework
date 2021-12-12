package handler

import (
	"bootcamp-homework/repository"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type AuthorBookHandler struct {
	authorRepository AuthorRepository
	bookRepository   BookRepository
}

func NewAuthorBookHandler(authorRepository AuthorRepository, bookRepository BookRepository) *AuthorBookHandler {
	return &AuthorBookHandler{authorRepository: authorRepository, bookRepository: bookRepository}
}

func (handler AuthorBookHandler) GetBooksOfAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	author, err := handler.authorRepository.Read(id)

	if errors.Is(err, repository.ErrAuthorNotFound) {
		w.WriteHeader(http.StatusNotFound)
		response.Error = err.Error()
		JSONWriter(w, response)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	books := handler.bookRepository.ReadByAuthor(author.ID)

	response.Data = books

	JSONWriter(w, response)
}

func (handler AuthorBookHandler) GetAuthorOfBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	book, err := handler.bookRepository.Read(id)

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

	author, err := handler.authorRepository.Read(book.AuthorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response.Data = author

	JSONWriter(w, response)
}
