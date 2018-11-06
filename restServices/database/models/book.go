package models

type Book struct {
	ID     int    `json:Id`
	Title  string `json:Title`
	Author string `json:Author`
	Year   string `json:Year`
}