package batch

import (
	"simulator/core/batch/cpu"
	"simulator/core/batch/task"
)

type TaskPackageGenerationArgs struct {
	Count int `json:"count"`
}

type GeneratedTaskPackage struct {
	TaskPackage []task.Task `json:"task_package"`
}

type TaskPackageProcessingArgs struct {
	TaskPackage []task.Task `json:"task_package"`
}

type TaskPackageProcessingResult struct {
	TaskPackage []task.Task     `json:"task_package"`
	CpuInfo     cpu.WorkingInfo `json:"cpu_info"`
}
