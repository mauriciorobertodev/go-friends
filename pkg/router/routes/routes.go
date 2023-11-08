package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                string
	Method             string
	Handler            func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

func SetRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	return r
}
