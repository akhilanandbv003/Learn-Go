package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDb() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_URL"))
	LogFatal(err)
	log.Println(pgUrl)

	db, err = sql.Open("postgres", pgUrl)
	LogFatal(err)

	err = db.Ping()
	LogFatal(err)

	return db
}