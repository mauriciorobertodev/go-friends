package routes

import (
	"go-friends/pkg/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:                "/users",
		Method:             http.MethodPost,
		Handler:            controllers.StoreUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users",
		Method:             http.MethodGet,
		Handler:            controllers.ListUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodGet,
		Handler:            controllers.GetUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodPut,
		Handler:            controllers.UpdateUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodDelete,
		Handler:            controllers.DeleteUser,
		NeedAuthentication: false,
	},
}
