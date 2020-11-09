package routes

import (
	"api/controllers"
	"net/http"
)

var usersRoutes = []Route {
	Route{
		Url: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Url: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		Url: "/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		Url: "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		Url: "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}