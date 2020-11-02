package routes

import (
	bookscontroller "bookstore/controllers/books"

	"github.com/gorilla/mux"
)

// Routes returns a new router instance with the API routes
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", bookscontroller.Create).Methods("POST")
	router.HandleFunc("/api/books/{skip}/{limit}", bookscontroller.FindAllPaginated).Methods("GET")
	router.HandleFunc("/api/books/{id}", bookscontroller.Update).Methods("PATCH")
	router.HandleFunc("/api/books/{id}", bookscontroller.Delete).Methods("DELETE")
	router.HandleFunc("/api/books/{id}", bookscontroller.FindByID).Methods("GET")
	return router
}
