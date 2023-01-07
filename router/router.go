package router

import (
	"github.com/gorilla/mux"
	"simulator/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// common
	router.HandleFunc("/api/ping", middleware.Ping)
	return router
}
