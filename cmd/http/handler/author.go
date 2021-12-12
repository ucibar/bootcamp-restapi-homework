package handler

import (
	"bootcamp-homework/model"
	"bootcamp-homework/repository"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type AuthorRepository interface {
	All() []*model.Author
	Read(id int) (*model.Author, error)
	Delete(id int) (*model.Author, error)
	Create(author *model.Author) (*model.Author, error)
	Update(id int, author *model.Author) error
}

type AuthorHandler struct {
	repository AuthorRepository
}

func NewAuthorHandler(repository AuthorRepository) *AuthorHandler {
	return &AuthorHandler{repository: repository}
}

func (handler AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := handler.repository.All()

	response := &JSONResponse{Data: authors}

	JSONWriter(w, response)
}

func (handler *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := &model.Author{}

	JSONReader(w, r.Body, author)

	author, err := handler.repository.Create(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := &JSONResponse{Data: author}
	w.WriteHeader(http.StatusCreated)

	JSONWriter(w, response)
}

func (handler *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	author, err := handler.repository.Read(id)

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

	response.Data = author

	JSONWriter(w, response)
}

func (handler *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	author := &model.Author{}

	JSONReader(w, r.Body, author)

	response := &JSONResponse{}

	err = handler.repository.Update(id, author)

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

	response.Data = author

	JSONWriter(w, response)
}

func (handler *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	author, err := handler.repository.Delete(id)

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

	response.Data = author

	JSONWriter(w, response)
}