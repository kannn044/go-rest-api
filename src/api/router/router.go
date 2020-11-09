package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}