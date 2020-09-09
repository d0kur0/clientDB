package httpHandlers

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	_, err :=w.Write([]byte("test"))

	if err != nil {
		log.Println(err)
	}
}