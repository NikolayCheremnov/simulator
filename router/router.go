package router

import (
	"github.com/gorilla/mux"
	"simulator/middleware/base"
	"simulator/middleware/tasks"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(base.SetApiContentType) // set content type header

	// common
	router.HandleFunc("/api/ping", base.Ping)

	// tasks
	router.HandleFunc("/api/generate-task-package", tasks.GenerateRandomTaskPackage).Methods("POST")
	router.HandleFunc("/api/process-task-package", tasks.ProcessTaskPackage).Methods("POST")

	return router
}
