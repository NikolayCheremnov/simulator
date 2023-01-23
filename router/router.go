package router

import (
	"github.com/gorilla/mux"
	"simulator/middleware/base"
	"simulator/middleware/batch"
	"simulator/middleware/multitask"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(base.SetApiHeaders) // set content type header

	// common
	router.HandleFunc("/api/ping", base.Ping)

	// prefixes
	const (
		BATCH_URL_PREFIX     = "/api/batch/"
		MULTITASK_URL_PREFIX = "/api/multitask/"
	)

	// batch
	router.HandleFunc(BATCH_URL_PREFIX+"generate-task-package", batch.GenerateRandomTaskPackage).Methods("POST")
	router.HandleFunc(BATCH_URL_PREFIX+"process-task-package", batch.ProcessTaskPackage).Methods("POST")
	router.HandleFunc(BATCH_URL_PREFIX+"task-package-time-report", batch.GenerateTaskPackageTimeReport).Methods("POST")
	router.HandleFunc(BATCH_URL_PREFIX+"cpu-activity-report", batch.GenerateCpuActivityReport).Methods("GET")
	router.HandleFunc(BATCH_URL_PREFIX+"average-task-metrics-report", batch.GenerateAverageTaskMetricsReport).Methods("GET")

	// multitask
	router.HandleFunc(MULTITASK_URL_PREFIX+"empty-memory-dump-report", multitask.GenerateEmptyMemoryDump).Methods("POST")
	router.HandleFunc(MULTITASK_URL_PREFIX+"simple-test", multitask.SimpleSystemTest).Methods("POST")
	router.HandleFunc(MULTITASK_URL_PREFIX+"generate-process-time-report", multitask.GenerateProcessTimeReport).Methods("POST")
	router.HandleFunc(MULTITASK_URL_PREFIX+"generate-cpu-metrics-report", multitask.GenerateCpuMetricsReport).Methods("POST")
	router.HandleFunc(MULTITASK_URL_PREFIX+"generate-memory-metrics-report", multitask.GenerateMemoryMetricsReport).Methods("POST")
	router.HandleFunc(MULTITASK_URL_PREFIX+"generate-memory-dump-report", multitask.GenerateMemoryDumpReport).Methods("POST")
	return router
}
