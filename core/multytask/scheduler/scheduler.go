package scheduler

import (
	"simulator/core/multytask/cpu"
	"simulator/core/multytask/ram"
)

func ExecuteProcessList(memory *ram.MemoryInfo, pageSizeB int, processList []ram.Condition, specifications cpu.Specifications) []ram.Condition {
	// 1. priority timesharing or fair guaranteed scheduling
	// cycle of cpu working
	executingCount, completedCount := len(processList), 0
	for i := 0; completedCount < executingCount; i = (i + 1) % len(processList) {
		if processList[i].ConditionString == ram.COMPLETED {
			continue // not touch completed
		}
		// 1. memory resources
		requiredPages := processList[i].GenerateRequiredPages() // generate required pages indexes
		// 2. uploading pages to memory
		ram.PrepareRequiredPages(memory, pageSizeB, &processList[i], requiredPages)
		// 3. calculate cpu resources
		cpuTime := int(float32(specifications.BaseCpuTimePart) / float32(processList[i].Process.Priority))
		// 4. process ready to work
		processList[i].ConditionString = ram.READY
		// 5. correct time for processes
		for j := 0; j < len(processList); j++ {
			if i == j {
				processList[j].ExecutingTime += cpuTime
			} else if processList[j].ConditionString != ram.COMPLETED {
				processList[j].WaitingTime += cpuTime
			}
		}
		// 6. process is active
		processList[i].ConditionString = ram.ACTIVE
		processList[i].CpuTime -= cpuTime
		// 7. process context changing
		if processList[i].CpuTime <= 0 {
			processList[i].ConditionString = ram.COMPLETED
			completedCount++
		} else {
			processList[i].ConditionString = ram.WAITING
		}
	}

	return processList
}

type CpuMetrics struct {
	Time          int `json:"time"`
	ProcessNumber int `json:"process_number"`
}

func ExecuteProcessListWithCpuMetrics(memory *ram.MemoryInfo, pageSizeB int,
	processList []ram.Condition, specifications cpu.Specifications) ([]ram.Condition, []CpuMetrics) {
	var cpuMetrics []CpuMetrics
	cpuMetrics = append(cpuMetrics, CpuMetrics{0, 0})
	timePassed := 0
	processPassed := 0

	// 1. priority timesharing or fair guaranteed scheduling
	// cycle of cpu working
	executingCount, completedCount := len(processList), 0
	for i := 0; completedCount < executingCount; i = (i + 1) % len(processList) {
		if processList[i].ConditionString == ram.COMPLETED {
			continue // not touch completed
		}
		// 1. memory resources
		requiredPages := processList[i].GenerateRequiredPages() // generate required pages indexes
		// 2. uploading pages to memory
		ram.PrepareRequiredPages(memory, pageSizeB, &processList[i], requiredPages)
		// 3. calculate cpu resources
		cpuTime := int(float32(specifications.BaseCpuTimePart) / float32(processList[i].Process.Priority))
		timePassed += cpuTime
		processPassed++

		if timePassed >= specifications.BaseCpuTimePart*10 {
			cpuMetrics = append(cpuMetrics, CpuMetrics{
				cpuMetrics[len(cpuMetrics)-1].Time + timePassed,
				processPassed,
			})
			timePassed = 0
			processPassed = 0
		}

		// 4. process ready to work
		processList[i].ConditionString = ram.READY
		// 5. correct time for processes
		for j := 0; j < len(processList); j++ {
			if i == j {
				processList[j].ExecutingTime += cpuTime
			} else if processList[j].ConditionString != ram.COMPLETED {
				processList[j].WaitingTime += cpuTime
			}
		}
		// 6. process is active
		processList[i].ConditionString = ram.ACTIVE
		processList[i].CpuTime -= cpuTime
		// 7. process context changing
		if processList[i].CpuTime <= 0 {
			processList[i].ConditionString = ram.COMPLETED
			completedCount++
		} else {
			processList[i].ConditionString = ram.WAITING
		}
	}

	return processList, cpuMetrics
}

type MemoryMetrics struct {
	Time         int `json:"time"`
	PagesWritten int `json:"pages_written"`
	PagesRemoved int `json:"pages_removed"`
}

func ExecuteProcessListWithMemoryMetrics(memory *ram.MemoryInfo, pageSizeB int,
	processList []ram.Condition, specifications cpu.Specifications) (*ram.MemoryInfo, []MemoryMetrics) {
	var memoryMetrics []MemoryMetrics
	memoryMetrics = append(memoryMetrics, MemoryMetrics{0, 0, 0})
	timePassed := 0
	pagesWrote := 0
	pagesRemoved := 0

	// 1. priority timesharing or fair guaranteed scheduling
	// cycle of cpu working
	executingCount, completedCount := len(processList), 0
	for i := 0; completedCount < executingCount; i = (i + 1) % len(processList) {
		if processList[i].ConditionString == ram.COMPLETED {
			continue // not touch completed
		}
		// 1. memory resources
		requiredPages := processList[i].GenerateRequiredPages() // generate required pages indexes
		// 2. uploading pages to memory
		isPageWrote, isPageRemoved := ram.PrepareRequiredPagesWithMemoryMetrics(memory, pageSizeB, &processList[i], requiredPages)
		if isPageWrote {
			pagesWrote++
		}
		if isPageRemoved {
			pagesRemoved++
		}

		// 3. calculate cpu resources
		cpuTime := int(float32(specifications.BaseCpuTimePart) / float32(processList[i].Process.Priority))
		timePassed += cpuTime

		if timePassed >= specifications.BaseCpuTimePart*10 {
			memoryMetrics = append(memoryMetrics, MemoryMetrics{
				memoryMetrics[len(memoryMetrics)-1].Time + timePassed,
				pagesWrote,
				pagesRemoved,
			})
			timePassed = 0
			pagesWrote = 0
			pagesRemoved = 0
		}

		// 4. process ready to work
		processList[i].ConditionString = ram.READY
		// 5. correct time for processes
		for j := 0; j < len(processList); j++ {
			if i == j {
				processList[j].ExecutingTime += cpuTime
			} else if processList[j].ConditionString != ram.COMPLETED {
				processList[j].WaitingTime += cpuTime
			}
		}
		// 6. process is active
		processList[i].ConditionString = ram.ACTIVE
		processList[i].CpuTime -= cpuTime
		// 7. process context changing
		if processList[i].CpuTime <= 0 {
			processList[i].ConditionString = ram.COMPLETED
			completedCount++
		} else {
			processList[i].ConditionString = ram.WAITING
		}
	}

	return memory, memoryMetrics
}
