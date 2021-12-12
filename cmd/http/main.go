package main

import (
	"bootcamp-homework/cmd/http/handler"
	"bootcamp-homework/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	authorRepo := repository.NewInMemoryAuthorRepository()
	authorHandler := handler.NewAuthorHandler(authorRepo)

	bookRepo := repository.NewInMemoryBookRepository()
	bookHandler := handler.NewBookHandler(bookRepo)

	authorBookHandler := handler.NewAuthorBookHandler(authorRepo, bookRepo)

	r := mux.NewRouter()

	r.HandleFunc("/authors", authorHandler.GetAllAuthors).Methods("GET")
	r.HandleFunc("/authors", authorHandler.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id:[0-9]+}", authorHandler.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id:[0-9]+}", authorHandler.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id:[0-9]+}", authorHandler.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors/{id:[0-9]+}/books", authorBookHandler.GetBooksOfAuthor).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}/author", authorBookHandler.GetAuthorOfBook).Methods("GET")

	log.Println("Server is listening on port 8089")
	log.Fatal(http.ListenAndServe(":8089", r))
}
