package tasks

import "simulator/core/task"

type TaskPackageGenerationArgs struct {
	Count int `json:"count"`
}

type GeneratedTaskPackage struct {
	Tasks        []task.Task `json:"tasks"`
	TotalCpuTime int         `json:"total_cpu_time"`
	TotalIoTime  int         `json:"total_io_time"`
}
