package batch

import (
	"encoding/json"
	"net/http"
	"simulator/core/batch/cpu"
	"simulator/core/batch/task"
	"simulator/logger"
	"simulator/middleware/base"
	"simulator/reports/batch"
)

func GenerateRandomTaskPackage(w http.ResponseWriter, r *http.Request) {

	var args TaskPackageGenerationArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateRandomTaskPackage args", w) != nil {
		return
	}
	tasks, _, _ := task.GenerateTaskPackage(args.Count)
	result := GeneratedTaskPackage{
		tasks,
	}
	logger.Info.Println("task package generated")
	err = json.NewEncoder(w).Encode(result)
	if base.ErrorHandling(err, "Can`t encode result for GenerateRandomTaskPackage", w) != nil {
		return
	}
}

func ProcessTaskPackage(w http.ResponseWriter, r *http.Request) {
	var args TaskPackageProcessingArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for ProcessTaskPackage args", w) != nil {
		return
	}
	cpuInfo, taskPackage := cpu.ExecuteTaskPackage(args.TaskPackage)
	result := TaskPackageProcessingResult{
		taskPackage,
		cpuInfo,
	}
	logger.Info.Println("task package processed")
	err = json.NewEncoder(w).Encode(result)
	if base.ErrorHandling(err, "Can`t encode result for ProcessTaskPackage", w) != nil {
		return
	}
}

func GenerateTaskPackageTimeReport(w http.ResponseWriter, r *http.Request) {
	var args TaskPackageProcessingResult
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateTaskPackageTimeReport args", w) != nil {
		return
	}
	result := batch.GenerateTaskPackageTimeReport(args.TaskPackage)
	logger.Info.Println("task package time report generated")
	err = json.NewEncoder(w).Encode(result)
	if base.ErrorHandling(err, "Can`t encode result for GenerateTaskPackageTimeReport", w) != nil {
		return
	}
}

func GenerateCpuActivityReport(w http.ResponseWriter, r *http.Request) {
	result := batch.GenerateCpuActivityReport()
	logger.Info.Println("cpu activity report generated")
	err := json.NewEncoder(w).Encode(result)
	if base.ErrorHandling(err, "Can`t encode result for GenerateCpuActivityReport", w) != nil {
		return
	}
}

func GenerateAverageTaskMetricsReport(w http.ResponseWriter, r *http.Request) {
	result := batch.GenerateAverageTaskMetricsReport()
	logger.Info.Println("average task metrics report generated")
	err := json.NewEncoder(w).Encode(result)
	if base.ErrorHandling(err, "Can`t encode result for GenerateAverageTaskMetricsReport", w) != nil {
		return
	}
}
