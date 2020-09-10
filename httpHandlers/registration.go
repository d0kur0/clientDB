package httpHandlers

import (
	"log"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test"))

	if err != nil {
		log.Println(err)
	}
}
