package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Book struct {
	ID int `json:Id`
	Title string `json:Title`
	Author string `json:Author`
	Year string `json:Year`
}

var books  []Book
func main() {

	books = append(books,
		Book{ID:1,Title:"Book1", Author:"author1", Year:"2018" },
		Book{ID:2,Title:"Book2", Author:"author2", Year:"2017" },
		Book{ID:3,Title:"Book3", Author:"author3", Year:"2016" },
		Book{ID:4,Title:"Book4", Author:"author4", Year:"2011" },
	)
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", addBooks).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",router))

}

func getBooks(w http.ResponseWriter , r *http.Request)  {
	log.Println("Inside getBooks: Getting all books ")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter , r *http.Request)  {
	log.Println("Inside getBook: Getting a book ")
	params := mux.Vars(r)
	log.Println(reflect.TypeOf(params["id"]))

	id ,_ := strconv.Atoi(params["id"])

	for _ , book := range books{
		if book.ID == id {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBooks(w http.ResponseWriter , r *http.Request)  {
	log.Println("Inside addBooks: Adding a book ")

	var addBookRequest Book
	_ = json.NewDecoder(r.Body).Decode(&addBookRequest)
	books = append(books,addBookRequest)
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter , r *http.Request)  {
	log.Println("Inside updateBook: Updating book ")

	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	for key , value := range books{
		if value.ID == updatedBook.ID{
			books[key] = updatedBook
		}
	}

	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter , r *http.Request)  {
	log.Println("Inside deleteBook: Deleting the book ")
	params := mux.Vars(r)
	id , _ := strconv.Atoi(params["id"])

	for i,item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}
