package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"restServices/database/controllers"
	"restServices/database/driver"
	"restServices/database/models"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDb()
	controller := controllers.Controller{}
	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/book", controller.AddBooks(db)).Methods("POST")
	router.HandleFunc("/book", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
