package main

import (
	"log"

	"github.com/d0kur0/clientDB/router"

	"github.com/d0kur0/clientDB/utils/configMgr"
)

func main() {
	if err := configMgr.Parse(); err != nil {
		log.Println(err)
		return
	}

	if err := router.Init(); err != nil {
		log.Println(err)
		return
	}
}
