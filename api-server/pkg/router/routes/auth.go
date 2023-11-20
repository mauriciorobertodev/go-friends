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
	{
		Uri:                "/update-password",
		Method:             http.MethodPost,
		Handler:            controllers.UpdatePassword,
		NeedAuthentication: true,
	},
}
