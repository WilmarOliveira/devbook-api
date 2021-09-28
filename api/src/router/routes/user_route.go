package routes

import (
	controller "api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Function:           controller.FindAllUsers,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:           controller.FindUserById,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Function:           controller.InsertUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodPut,
		Function:           controller.UpdateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodDelete,
		Function:           controller.DeleteUser,
		NeedAuthentication: false,
	},
}
