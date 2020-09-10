package main

import (
	"log"

	"github.com/d0kur0/clientDB/redisWrapper"

	"github.com/d0kur0/clientDB/router"

	"github.com/d0kur0/clientDB/dbWrapper"

	"github.com/d0kur0/clientDB/utils/configMgr"
)

func main() {
	if err := configMgr.Parse(); err != nil {
		log.Println(err)
		return
	}

	if err := dbWrapper.Init(); err != nil {
		log.Println(err)
		return
	}

	if err := dbWrapper.Migrate(); err != nil {
		log.Println(err)
		return
	}

	if err := redisWrapper.Init(); err != nil {
		log.Println(err)
		return
	}

	if err := router.Init(); err != nil {
		log.Println(err)
		return
	}
}
