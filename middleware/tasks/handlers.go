package tasks

import (
	"encoding/json"
	"net/http"
	"simulator/core/cpu"
	"simulator/core/task"
	"simulator/logger"
	"simulator/middleware/base"
)

func GenerateRandomTaskPackage(w http.ResponseWriter, r *http.Request) {

	var args TaskPackageGenerationArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateRandomTaskPackage args", w) != nil {
		return
	}
	tasks, totalCpuTime, totalIoTime := task.GenerateTaskPackage(args.Count)
	result := GeneratedTaskPackage{
		tasks, totalCpuTime, totalIoTime,
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
	logger.Info.Println("task package generated")
	err = json.NewEncoder(w).Encode(result)
	if base.ErrorHandling(err, "Can`t encode result for ProcessTaskPackage", w) != nil {
		return
	}
}
