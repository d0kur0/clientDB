package router

import "github.com/d0kur0/clientDB/httpHandlers"

var Routes = []Route{
	{
		Methods:  []string{"POST"},
		Path:     "/registration",
		Handler:  httpHandlers.Registration,
		AuthNeed: false,
	},
}
