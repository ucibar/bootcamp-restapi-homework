package main

import (
	"bootcamp-homework/cmd/http/handler"
	"bootcamp-homework/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// initialize the author repository
	authorRepo := repository.NewInMemoryAuthorRepository()
	// initialize the handler for author resource
	authorHandler := handler.NewAuthorHandler(authorRepo)

	// initialize the book repository
	bookRepo := repository.NewInMemoryBookRepository()
	// initialize the handler for book resource
	bookHandler := handler.NewBookHandler(bookRepo)

	// initialize a handler for author books resource
	authorBookHandler := handler.NewAuthorBookHandler(authorRepo, bookRepo)

	// initialize the order repository
	orderRepo := repository.NewInMemoryOrderRepository()
	// initialize the handler for order resource
	orderHandler := handler.NewOrderHandler(orderRepo)

	// initialize mux router
	r := mux.NewRouter()

	// Global middleware for SetContentType application/json
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	// endpoints for author resource
	r.HandleFunc("/authors", authorHandler.GetAllAuthors).Methods("GET")
	r.HandleFunc("/authors", authorHandler.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id:[0-9]+}", authorHandler.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id:[0-9]+}", authorHandler.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id:[0-9]+}", authorHandler.DeleteAuthor).Methods("DELETE")

	// endpoints for book resource
	r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.DeleteBook).Methods("DELETE")

	// endpoints for author books resource
	r.HandleFunc("/authors/{id:[0-9]+}/books", authorBookHandler.GetBooksOfAuthor).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}/author", authorBookHandler.GetAuthorOfBook).Methods("GET")

	// endpoints for order resource
	r.HandleFunc("/orders", orderHandler.GetAllOrders).Methods("GET")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id:[0-9]+}", orderHandler.GetOrder).Methods("GET")

	log.Println("Server is listening on port 8089")
	log.Fatal(http.ListenAndServe(":8089", r))
}
