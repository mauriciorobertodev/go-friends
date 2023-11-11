package routes

import (
	"go-friends/pkg/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		Uri:                "/posts",
		Method:             http.MethodPost,
		Handler:            controllers.StorePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts",
		Method:             http.MethodGet,
		Handler:            controllers.ListPosts,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodGet,
		Handler:            controllers.GetPost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodPut,
		Handler:            controllers.UpdatePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodDelete,
		Handler:            controllers.DeletePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{userId}/posts",
		Method:             http.MethodGet,
		Handler:            controllers.GetPostsOfUser,
		NeedAuthentication: true,
	},
}
