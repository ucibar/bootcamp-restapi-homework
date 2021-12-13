package handler

import (
	"bootcamp-homework/model"
	"bootcamp-homework/service"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	service *service.AuthorService
}

func NewAuthorHandler(service *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (handler *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := handler.service.GetAllAuthors()

	response := &JSONResponse{Data: authors}

	JSONWriter(w, response)
}

func (handler *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := &model.Author{}

	JSONReader(w, r.Body, author)

	author, err := handler.service.CreateAuthor(author)
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

	author, err := handler.service.GetAuthorByID(id)

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

	response.Data = author

	JSONWriter(w, response)
}

func (handler *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	author := &model.Author{}

	JSONReader(w, r.Body, author)

	response := &JSONResponse{}

	err = handler.service.UpdateAuthor(id, author)

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

	response.Data = author

	JSONWriter(w, response)
}

func (handler *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	response := &JSONResponse{}

	author, err := handler.service.DeleteAuthor(id)

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

	response.Data = author

	JSONWriter(w, response)
}
