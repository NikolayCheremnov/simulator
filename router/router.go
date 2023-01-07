package router

import (
	"github.com/gorilla/mux"
	"simulator/middleware/base"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// common
	router.HandleFunc("/api/ping", base.Ping)
	return router
}
