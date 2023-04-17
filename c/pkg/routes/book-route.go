package routes

import (
	"c/pkg/controllers"
	"github.com/gorilla/mux"
)

var Init_book_route = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{ID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{ID}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{ID}", controllers.GetBook).Methods("GET")

}
