package task

import (
	"github.com/google/uuid"
	"simulator/core/random"
	"strings"
)

func GenerateTaskPackage(count int) ([]Task, int, int) {
	// randomizers
	boolgen := random.BoolRandom()
	intgen := random.IntRandom()
	// tasks generation
	var tasks []Task
	total_required_cpu_time, total_required_io_time := 0, 0
	for i := 0; i < count; i++ {
		var task Task
		// 1. generate pid
		task.PID = get_uuid()
		// 2. generate live
		required_cpu_time := intgen.Int(10, 100)
		required_io_time := intgen.Int(100, 500)
		total_required_cpu_time += required_cpu_time
		total_required_io_time += required_io_time
		// disperse task time
		for required_cpu_time != 0 && required_io_time != 0 {
			if required_io_time == 0 {
				// then write remaining cpu time to task
				for i := 0; i < required_cpu_time; i++ {
					task.RequiredSteps = append(task.RequiredSteps, CPU_STEP)
				}
				required_cpu_time = 0
			} else if required_cpu_time == 0 {
				for i := 0; i < required_io_time; i++ {
					task.RequiredSteps = append(task.RequiredSteps, IO_STEP)
				}
				required_io_time = 0
			} else if boolgen.Bool() != boolgen.Bool() {
				task.RequiredSteps = append(task.RequiredSteps, CPU_STEP)
				required_cpu_time--
			} else {
				task.RequiredSteps = append(task.RequiredSteps, IO_STEP)
				required_io_time--
			}
		}
		//
		task.NexLiveTimeStep, task.NextRequiredStep = 0, 0
		tasks = append(tasks, task)
	}
	return tasks, total_required_cpu_time, total_required_io_time
}

func get_uuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
