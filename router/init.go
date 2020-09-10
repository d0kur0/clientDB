package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/d0kur0/clientDB/utils/configMgr"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() error {
	config := configMgr.Get()

	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

	for _, route := range Routes {
		router.HandleFunc(route.Path, authenticationMiddleware(route)).Methods(route.Methods...)
	}

	log.Println("Server started at port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handlers.CompressHandler(router)); err != nil {
		return errors.Wrap(err, "Starting server failed")
	}

	return nil
}
