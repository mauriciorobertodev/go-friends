package router

import (
	"go-friends/pkg/router/routes"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.SetRoutes(r)
}
