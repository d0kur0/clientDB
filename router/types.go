package router

import "net/http"

type Route struct {
	Methods  []string
	Path     string
	Handler  func(w http.ResponseWriter, r *http.Request)
	AuthNeed bool
}
