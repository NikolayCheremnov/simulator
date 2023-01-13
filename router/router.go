package router

import (
	"github.com/gorilla/mux"
	"simulator/middleware/base"
	"simulator/middleware/batch"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(base.SetApiContentType) // set content type header

	// common
	router.HandleFunc("/api/ping", base.Ping)

	// batch
	router.HandleFunc("/api/generate-task-package", batch.GenerateRandomTaskPackage).Methods("POST")
	router.HandleFunc("/api/process-task-package", batch.ProcessTaskPackage).Methods("POST")
	router.HandleFunc("/api/task-package-time-report", batch.GenerateTaskPackageTimeReport).Methods("POST")
	router.HandleFunc("/api/cpu-activity-report", batch.GenerateCpuActivityReport).Methods("GET")
	router.HandleFunc("/api/average-task-metrics-report", batch.GenerateAverageTaskMetricsReport).Methods("GET")

	return router
}
