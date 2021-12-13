package main

import (
	"bootcamp-homework/cmd/http/handler"
	"bootcamp-homework/repository"
	"bootcamp-homework/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// initialize the author repository
	authorRepo := repository.NewInMemoryAuthorRepository()
	// initialize the author service
	authorService := service.NewAuthorService(authorRepo)
	// initialize the handler for author resource
	authorHandler := handler.NewAuthorHandler(authorService)

	// initialize the book repository
	bookRepo := repository.NewInMemoryBookRepository()
	// initialize the book service
	bookService := service.NewBookService(bookRepo, authorRepo)
	// initialize the handler for book resource
	bookHandler := handler.NewBookHandler(bookService)

	// initialize the order repository
	orderRepo := repository.NewInMemoryOrderRepository()
	// initialize the order service
	orderService := service.NewOrderService(orderRepo)
	// initialize the handler for order resource
	orderHandler := handler.NewOrderHandler(orderService)

	// initialize mux router
	r := mux.NewRouter()

	// Global middleware for Set Content-Type to application/json
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
	r.HandleFunc("/authors/{id:[0-9]+}/books", bookHandler.GetBooksOfAuthor).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}/author", bookHandler.GetAuthorOfBook).Methods("GET")

	// endpoints for order resource
	r.HandleFunc("/orders", orderHandler.GetAllOrders).Methods("GET")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id:[0-9]+}", orderHandler.GetOrder).Methods("GET")

	log.Println("Server is listening on port 8089")
	log.Fatal(http.ListenAndServe(":8089", r))
}
