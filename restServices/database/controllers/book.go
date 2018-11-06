package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restServices/database/driver"
	"restServices/database/models"
)

type Controller struct{}

var books []models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		rows, err := db.Query("select * from books")
		driver.LogFatal(err)

		defer rows.Close()
		for rows.Next() {
			rows.Scan(&book.ID, &book.Year, &book.Author, &book.Title)
			driver.LogFatal(err)
			books = append(books, book)
		}
		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		parse := mux.Vars(r)

		rows := db.QueryRow("select * from books where id = $1 ", parse["id"])

		err1 := rows.Scan(&book.ID, &book.Year, &book.Author, &book.Title)
		driver.LogFatal(err1)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookId int

		json.NewDecoder(r.Body).Decode(&book)
		err := db.QueryRow("INSERT into books (title,author,year) values ($1,$2,$3) returning ID;", book.Title, book.Author, book.Year).Scan(&bookId)
		driver.LogFatal(err)
		json.NewEncoder(w).Encode(bookId)

	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		result, err := db.Exec("UPDATE books set title = $1 , author = $2 , year=$3 where id = $4 returning id",
			&book.Title, &book.Author, &book.Year, &book.ID)
		rowsUpdated, err := result.RowsAffected()
		driver.LogFatal(err)
		json.NewEncoder(w).Encode(rowsUpdated)

	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parse := mux.Vars(r)
		id := parse["id"]

		result, err := db.Exec("delete from books where id = $1 returning id", id)
		rowsUpdated, err := result.RowsAffected()
		driver.LogFatal(err)
		json.NewEncoder(w).Encode(rowsUpdated)

	}
}
