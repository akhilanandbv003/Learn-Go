package songs_api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type Song struct {
	ID int `json:id`
	Title string `json:title`
	Artist string `json:author`
	Year string `json:year`
}

var songs [] Song

func main()  {
router :=mux.NewRouter()

router.HandleFunc


}
