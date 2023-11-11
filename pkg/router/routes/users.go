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
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodGet,
		Handler:            controllers.GetUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodPut,
		Handler:            controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodDelete,
		Handler:            controllers.DeleteUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}/follow",
		Method:             http.MethodPost,
		Handler:            controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}/unfollow",
		Method:             http.MethodPost,
		Handler:            controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}/followers",
		Method:             http.MethodGet,
		Handler:            controllers.GetFollowers,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{id}/following",
		Method:             http.MethodGet,
		Handler:            controllers.GetFollowing,
		NeedAuthentication: true,
	},
}
