package main

import (
	"simulator/core/multytask/cpu"
	"simulator/core/multytask/ram"
	"simulator/core/multytask/scheduler"
	"simulator/logger"
)

func local_test() {
	memory := ram.CreateMemoryStructure(32, 16,
		8, 4096)
	logger.Info.Println("empty memory generated")
	// 3. generate process list
	processList := ram.GenerateProcessConditionListToExecution(16, 0, 31, 16, 8, 4096, cpu.Specifications{
		100,
	})
	logger.Info.Println("process list generated")
	// 4. execute
	scheduler.ExecuteProcessList(&memory, 32, processList, cpu.Specifications{100})
	logger.Info.Println("executed")
}
