package routes

import (
	"go-friends/pkg/controllers"
	"net/http"
)

var authRoutes = []Route{
	{
		Uri:                "/login",
		Method:             http.MethodPost,
		Handler:            controllers.Login,
		NeedAuthentication: false,
	},
}
