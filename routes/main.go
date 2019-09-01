package routes

import (
	"github.com/gorilla/mux"
)

// Router mux Router
var router *mux.Router

func init() {
	router = mux.NewRouter()
	setStuffRoutes()
}

// GetRouter function
func GetRouter() *mux.Router {
	return router
}
