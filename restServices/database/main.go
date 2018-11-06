package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID     int    `json:Id`
	Title  string `json:Title`
	Author string `json:Author`
	Year   string `json:Year`
}

var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

//"postgres://bob:secret@1.2.3.4:5432/mydb?sslmode=verify-full"

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_URL"))
	logFatal(err)
	log.Println(pgUrl)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)
	err = db.Ping()
	logFatal(err)
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", addBooks).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}
	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&book.ID, &book.Year, &book.Author, &book.Title)
		logFatal(err)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	parse := mux.Vars(r)

	rows := db.QueryRow("select * from books where id = $1 ", parse["id"])

	err1 := rows.Scan(&book.ID, &book.Year, &book.Author, &book.Title)
	logFatal(err1)

	json.NewEncoder(w).Encode(book)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
var book Book
var bookId int

 json.NewDecoder(r.Body).Decode(&book)
 err := db.QueryRow("INSERT into books (title,author,year) values ($1,$2,$3) returning ID;", book.Title,book.Author,book.Year).Scan(&bookId)
 logFatal(err)
 json.NewEncoder(w).Encode(bookId)


}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

result ,err := db.Exec("UPDATE books set title = $1 , author = $2 , year=$3 where id = $4 returning id" ,
	&book.Title,&book.Author , &book.Year, &book.ID)
rowsUpdated , err := result.RowsAffected()
logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdated)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	parse := mux.Vars(r)
	id := parse["id"]

	result ,err := db.Exec("delete from books where id = $1 returning id" ,id)
	rowsUpdated , err := result.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdated)


}
