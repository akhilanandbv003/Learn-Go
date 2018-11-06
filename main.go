package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
type Song struct {
	ID int `json:id`
	Title string `json:title`
	Artist string `json:author`
	Year string `json:year`
}

var songs [] Song

func main()  {
router := mux.NewRouter()

router.HandleFunc("/songs", getSongs).Methods("GET")
router.HandleFunc("/song/{id}", getSong).Methods("GET")
router.HandleFunc("/songs", addSong).Methods("POST")
router.HandleFunc("/songs", updateSong).Methods("PUT")
router.HandleFunc("/songs/{id}", removeSong).Methods("DELETE")

log.Fatal(http.ListenAndServe(":8000" , router))
}

func getSongs( w http.ResponseWriter , r *http.Request)  {
	log.Println(" Get all songs is called ")
}

func getSong( w http.ResponseWriter , r *http.Request)  {
	log.Println(" Get song is called ")
}

func addSong( w http.ResponseWriter , r *http.Request)  {
	log.Println(" addSong is called ")
}

func updateSong( w http.ResponseWriter , r *http.Request)  {
	log.Println(" updateSong is called ")
}

func removeSong( w http.ResponseWriter , r *http.Request)  {
	log.Println(" removeSong is called ")
}