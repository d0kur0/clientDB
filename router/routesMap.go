package router

import "github.com/d0kur0/clientDB/httpHandlers"

var Routes = []Route{
	{
		Methods: []string{"GET"},
		Path: "/",
		Handler: httpHandlers.Test,
	},
}