package cpu

import "simulator/core/task"

func ExecuteTaskPackage(taskPackage []task.Task) (WorkingInfo, []task.Task) {
	completedTaskCount := 0
	taskCount := len(taskPackage)
	var cpuInfo WorkingInfo
	// working cycle
	currentTaskIndex := -1
	for completedTaskCount != taskCount {
		// choose task if it need
		if currentTaskIndex == -1 {
			currentTaskIndex = GetNextExecutableTask(taskPackage)
		}
		// update all tasks
		for i := 0; i < taskCount; i++ {
			if !taskPackage[i].IsCompleted() {
				taskPackage[i].Step(currentTaskIndex == i)
				if taskPackage[i].IsCompleted() {
					completedTaskCount++
				}
			}
		}
		// write cpu info and make tick
		if currentTaskIndex == -1 {
			// no task
			cpuInfo.StatusCache = append(cpuInfo.StatusCache, IDLE)
		} else {
			// with task
			cpuInfo.StatusCache = append(cpuInfo.StatusCache, WORKING)
			if taskPackage[currentTaskIndex].IsCompleted() || taskPackage[currentTaskIndex].NextStep() != task.CPU_STEP {
				currentTaskIndex = -1
			}
		}
	}
	return cpuInfo, taskPackage
}

func GetNextExecutableTask(taskPackage []task.Task) int {
	// 1. filter by executable task
	minValue, minIndex := -1, -1
	for i, nextTask := range taskPackage {
		if !nextTask.IsCompleted() && nextTask.ReadyWorkOnCpu() {
			if minIndex == -1 || nextTask.WorkLeft() < minValue {
				minValue = nextTask.WorkLeft()
				minIndex = i
			}
		}
	}
	return minIndex
}
