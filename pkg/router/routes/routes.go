package routes

import (
	"go-friends/pkg/middlewares"
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

	routes = append(routes, authRoutes...)

	for _, route := range routes {
		if route.NeedAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Auth(route.Handler))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	return r
}
