package reports

import (
	"simulator/core/batch/cpu"
	task2 "simulator/core/batch/task"
)

type TaskTimeInfo struct {
	PID         string `json:"pid"`
	CpuTime     int    `json:"cpu_time"`
	IoTime      int    `json:"io_time"`
	WaitingTime int    `json:"waiting_time"`
	WorkingTime int    `json:"working_time"`
}

func GenerateTaskPackageTimeReport(completedTaskPackage []task2.Task) []TaskTimeInfo {
	var report []TaskTimeInfo
	for _, completedTask := range completedTaskPackage {
		report = append(report, TaskTimeInfo{
			completedTask.PID,
			completedTask.GetCpuTime(),
			completedTask.GetIoTime(),
			completedTask.GetWaitingTime(),
			completedTask.GetTime(),
		})
	}
	return report
}

type CpuActivityInfo struct {
	TaskCount          int     `json:"task_count"`
	CpuActivityMetrics float64 `json:"cpu_activity_metrics"`
}

func GenerateCpuActivityReport() []CpuActivityInfo {
	var report []CpuActivityInfo
	for taskCount := 10; taskCount <= 100; taskCount += 5 {
		// 1. generate task package
		taskPackage, _, _ := task2.GenerateTaskPackage(taskCount)
		// 2. process task package
		workingInfo, taskPackage := cpu.ExecuteTaskPackage(taskPackage)
		// 3. add info
		report = append(report, CpuActivityInfo{
			taskCount,
			workingInfo.CalculateCpuActivityMetrics(),
		})
	}
	return report
}

type AverageTaskMetrics struct {
	TaskCount             int `json:"task_count"`
	AverageTaskLiveTime   int `json:"average_task_live_time"`
	AverageTaskComplexity int `json:"average_task_complexity"`
}

func GenerateAverageTaskMetricsReport() []AverageTaskMetrics {
	var report []AverageTaskMetrics
	for taskCount := 10; taskCount <= 100; taskCount += 5 {
		// 1. generate task package
		taskPackage, _, _ := task2.GenerateTaskPackage(taskCount)
		// 2. process task package
		_, taskPackage = cpu.ExecuteTaskPackage(taskPackage)
		// 3. calculate average live time and complexity
		averageTaskLiveTime, averageTaskComplexity := 0.0, 0.0
		for _, executableTask := range taskPackage {
			averageTaskComplexity += float64(executableTask.GetCpuTime()) + float64(executableTask.GetIoTime())/10.0
			averageTaskLiveTime += float64(executableTask.GetTime())
		}
		averageTaskComplexity /= float64(len(taskPackage))
		averageTaskLiveTime /= float64(len(taskPackage))

		report = append(report, AverageTaskMetrics{
			taskCount,
			int(averageTaskLiveTime),
			int(averageTaskComplexity),
		})
	}
	return report
}
