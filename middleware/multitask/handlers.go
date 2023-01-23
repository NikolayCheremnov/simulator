package multitask

import (
	"encoding/json"
	"net/http"
	"simulator/core/multytask/cpu"
	"simulator/core/multytask/ram"
	"simulator/core/multytask/scheduler"
	"simulator/logger"
	"simulator/middleware/base"
	"simulator/reports/multitask"
)

func GenerateEmptyMemoryDump(w http.ResponseWriter, r *http.Request) {
	// 1. get memory args
	var initMemoryArgs InitialMemoryArgs
	err := json.NewDecoder(r.Body).Decode(&initMemoryArgs)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateEmptyMemoryDump args", w) != nil {
		return
	}
	// 2. create memory
	memory := ram.CreateMemoryStructure(initMemoryArgs.MemoryBlockSizeB, initMemoryArgs.MaxProcessNumb,
		initMemoryArgs.ProcessRealPageCount, initMemoryArgs.PageSizeB)
	logger.Info.Println("empty memory generated")
	// 3. create dump
	emptyMemoryDump := multitask.GenerateMemoryDump(&memory)
	logger.Info.Println("empty memory dump generated")
	err = json.NewEncoder(w).Encode(emptyMemoryDump)
	if base.ErrorHandling(err, "Can`t encode result for GenerateCpuActivityReport", w) != nil {
		return
	}
}

func SimpleSystemTest(w http.ResponseWriter, r *http.Request) {
	// 1. get memory args
	var initMemoryArgs InitialMemoryArgs
	err := json.NewDecoder(r.Body).Decode(&initMemoryArgs)
	if base.ErrorHandling(err, "Can`t decode request body for SimpleSystemTest args", w) != nil {
		return
	}
	// 2. create memory
	memory := ram.CreateMemoryStructure(initMemoryArgs.MemoryBlockSizeB, initMemoryArgs.MaxProcessNumb,
		initMemoryArgs.ProcessRealPageCount, initMemoryArgs.PageSizeB)
	logger.Info.Println("empty memory generated")
	// 3. generate process list
	processList := ram.GenerateProcessConditionListToExecution(16, 0, 31, 16, 8, 32, cpu.Specifications{
		100,
	})
	logger.Info.Println("process list generated")
	// 4. execute
	scheduler.ExecuteProcessList(&memory, 32, processList, cpu.Specifications{100})
	//
	err = json.NewEncoder(w).Encode(base.BasicResponse{Message: "passed"})
	if base.ErrorHandling(err, "Can`t encode result for SimpleSystemTest", w) != nil {
		return
	}
}

func GenerateProcessTimeReport(w http.ResponseWriter, r *http.Request) {
	// 1. get report args
	var args multitask.ProcessTimeReportArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateProcessTimeReport args", w) != nil {
		return
	}
	// 2. generate report
	report := multitask.GenerateProcessTimeReport(args)
	err = json.NewEncoder(w).Encode(report)
	if base.ErrorHandling(err, "Can`t encode result for GenerateProcessTimeReport", w) != nil {
		return
	}
}

func GenerateCpuMetricsReport(w http.ResponseWriter, r *http.Request) {
	// 1. get report args
	var args multitask.ProcessTimeReportArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateCpuMetricsReport args", w) != nil {
		return
	}
	// 2. generate report
	report := multitask.GenerateCpuMetricsReport(args)
	err = json.NewEncoder(w).Encode(report)
	if base.ErrorHandling(err, "Can`t encode result for GenerateCpuMetricsReport", w) != nil {
		return
	}
}

func GenerateMemoryMetricsReport(w http.ResponseWriter, r *http.Request) {
	// 1. get report args
	var args multitask.ProcessTimeReportArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateMemoryMetricsReport args", w) != nil {
		return
	}
	// 2. generate report
	_, report := multitask.GenerateMemoryMetricsReport(args)
	err = json.NewEncoder(w).Encode(report)
	if base.ErrorHandling(err, "Can`t encode result for GenerateMemoryMetricsReport", w) != nil {
		return
	}
}

func GenerateMemoryDumpReport(w http.ResponseWriter, r *http.Request) {
	// 1. get report args
	var args multitask.ProcessTimeReportArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if base.ErrorHandling(err, "Can`t decode request body for GenerateMemoryDumpReport args", w) != nil {
		return
	}
	// 2. generate report
	memory, _ := multitask.GenerateMemoryMetricsReport(args)
	dump := multitask.GenerateMemoryDump(memory)
	err = json.NewEncoder(w).Encode(dump)
	if base.ErrorHandling(err, "Can`t encode result for GenerateMemoryDumpReport", w) != nil {
		return
	}
}
