package multitask

import (
	"simulator/core/multytask/cpu"
	"simulator/core/multytask/ram"
	"simulator/core/multytask/scheduler"
	"simulator/logger"
)

type ProcessTimeInfo struct {
	Priority    uint8 `json:"priority"`
	WaitingTime int   `json:"waiting_time"`
}

type ProcessTimeReportArgs struct {
	// memory args
	MemoryBlockSizeB int `json:"memory_block_size_b"`

	// process args
	MaxProcessNumb int   `json:"max_process_numb"`
	ProcessNumb    int   `json:"process_numb"`
	MinPriority    uint8 `json:"min_priority"`
	MaxPriority    uint8 `json:"max_priority"`
	// pages
	ProcessRealPageCount    int `json:"process_real_page_count"`
	ProcessVirtualPageCount int `json:"process_virtual_page_count"`
	PageSizeB               int `json:"page_size_b"`
}

func GenerateProcessTimeReport(args ProcessTimeReportArgs) []ProcessTimeInfo {
	// 1. create memory
	memory := ram.CreateMemoryStructure(args.MemoryBlockSizeB, args.MaxProcessNumb, args.ProcessRealPageCount,
		args.PageSizeB)
	logger.Info.Println("empty memory generated")
	// 3. generate process list
	CPUS := cpu.Specifications{BaseCpuTimePart: 100}
	processList := ram.GenerateProcessConditionListToExecution(args.ProcessNumb, args.MinPriority, args.MaxPriority,
		args.ProcessVirtualPageCount, args.ProcessRealPageCount, args.PageSizeB, CPUS)
	logger.Info.Println("process list generated")
	// 4. execute
	resultProcessList := scheduler.ExecuteProcessList(&memory, args.PageSizeB, processList, CPUS)
	// 5. make report
	var result []ProcessTimeInfo
	for _, process := range resultProcessList {
		result = append(result, ProcessTimeInfo{
			process.Process.Priority,
			process.WaitingTime,
		})
	}
	return result
}

func GenerateCpuMetricsReport(args ProcessTimeReportArgs) []scheduler.CpuMetrics {
	// 1. create memory
	memory := ram.CreateMemoryStructure(args.MemoryBlockSizeB, args.MaxProcessNumb, args.ProcessRealPageCount,
		args.PageSizeB)
	logger.Info.Println("empty memory generated")
	// 3. generate process list
	CPUS := cpu.Specifications{BaseCpuTimePart: 100}
	processList := ram.GenerateProcessConditionListToExecution(args.ProcessNumb, args.MinPriority, args.MaxPriority,
		args.ProcessVirtualPageCount, args.ProcessRealPageCount, args.PageSizeB, CPUS)
	logger.Info.Println("process list generated")
	// 4. execute
	_, cpuMetrics := scheduler.ExecuteProcessListWithCpuMetrics(&memory, args.PageSizeB, processList, CPUS)
	return cpuMetrics
}

func GenerateMemoryMetricsReport(args ProcessTimeReportArgs) (*ram.MemoryInfo, []scheduler.MemoryMetrics) {
	// 1. create memory
	memory := ram.CreateMemoryStructure(args.MemoryBlockSizeB, args.MaxProcessNumb, args.ProcessRealPageCount,
		args.PageSizeB)
	logger.Info.Println("empty memory generated")
	// 3. generate process list
	CPUS := cpu.Specifications{BaseCpuTimePart: 100}
	processList := ram.GenerateProcessConditionListToExecution(args.ProcessNumb, args.MinPriority, args.MaxPriority,
		args.ProcessVirtualPageCount, args.ProcessRealPageCount, args.PageSizeB, CPUS)
	logger.Info.Println("process list generated")
	// 4. execute
	memoryInfo, memoryMetrics := scheduler.ExecuteProcessListWithMemoryMetrics(&memory, args.PageSizeB, processList, CPUS)
	return memoryInfo, memoryMetrics
}
