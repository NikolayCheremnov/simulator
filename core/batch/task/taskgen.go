package task

import (
	"github.com/google/uuid"
	random2 "simulator/utils/random"
	"strings"
)

func GenerateTaskPackage(count int) ([]Task, int, int) {
	// randomizers
	boolgen := random2.BoolRandom()
	intgen := random2.IntRandom()
	// batch generation
	var tasks []Task
	totalRequiredCpuTime, totalRequiredIoTime := 0, 0
	for i := 0; i < count; i++ {
		var task Task
		// 1. generate pid
		task.PID = get_uuid()
		// 2. generate live
		requiredCpuTime := intgen.Int(100, 1000)
		requiredIoTime := intgen.Int(100, 10000)
		totalRequiredCpuTime += requiredCpuTime
		totalRequiredIoTime += requiredIoTime
		// disperse task time
		for requiredCpuTime != 0 && requiredIoTime != 0 {
			if requiredIoTime == 0 {
				// then write remaining cpu time to task
				for i := 0; i < requiredCpuTime; i++ {
					task.RequiredSteps = append(task.RequiredSteps, CPU_STEP)
				}
				requiredCpuTime = 0
			} else if requiredCpuTime == 0 {
				for i := 0; i < requiredIoTime; i++ {
					task.RequiredSteps = append(task.RequiredSteps, IO_STEP)
				}
				requiredIoTime = 0
			} else if boolgen.Bool() {
				task.RequiredSteps = append(task.RequiredSteps, CPU_STEP)
				requiredCpuTime--
			} else {
				task.RequiredSteps = append(task.RequiredSteps, IO_STEP)
				requiredIoTime--
			}
		}
		//
		task.NextLiveTimeStep, task.NextRequiredStep = 0, 0
		tasks = append(tasks, task)
	}
	return tasks, totalRequiredCpuTime, totalRequiredIoTime
}

func get_uuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)[:8]
}
