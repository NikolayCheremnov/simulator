package ram

import (
	"github.com/google/uuid"
	"simulator/core/multytask/cpu"
	"simulator/core/multytask/process"
	"simulator/utils/random"
	"strconv"
	"strings"
)

type ConditionString string

const (
	READY     = "ready"
	ACTIVE    = "active"
	WAITING   = "waiting"
	COMPLETED = "completed"
)

// Condition of the process
type Condition struct {
	Process         process.Process
	ConditionString ConditionString
	// required resources for process at now time
	CpuTime      int
	RamSize      int
	ProcessPages []PageInfo
	PageTable    []Page
	ClockPointer int // pointer for WSClock algorithm
	// for reports
	ExecutingTime int
	WaitingTime   int
}

// GetPPString mean that get process-page info string
func (condition *Condition) GetPPString(pageIndex int) string {
	return "process " + condition.Process.PID + ", page#" + strconv.Itoa(pageIndex)
}

func (condition *Condition) GenerateRequiredPages() []int {
	intGen := random.IntRandom()
	pageCount := len(condition.ProcessPages)
	maxRequiredPagesCount := pageCount / 2 // half of all page count
	requiredPagesCount := intGen.Int(1, maxRequiredPagesCount)
	var requiredPages []int // slice of indexes
	pageIndex := intGen.IntN(pageCount - 1)
	for i := 0; i < requiredPagesCount; i++ {
		requiredPages = append(requiredPages, pageIndex)
		pageIndex = (pageIndex + 1) % pageCount
	}
	return requiredPages
}

func GenerateProcessConditionListToExecution(processNumb int, minPriority uint8, maxPriority uint8,
	virtualPageCount int, realPageCount int, pageSizeB int, specifications cpu.Specifications) []Condition {
	// pageCount - maxPageCount for process
	intGen := random.IntRandom()
	vitualAddrSpaceSize := virtualPageCount * pageSizeB
	processConditionList := make([]Condition, processNumb)
	for i := 0; i < processNumb; i++ {
		// 1. init the process
		processConditionList[i].Process.PID = get_uuid(8)
		processConditionList[i].Process.Priority = uint8(intGen.Int(int(minPriority), int(maxPriority)))
		processConditionList[i].Process.Specifications = "autogenerated #" + strconv.Itoa(i)
		// 2. init condition
		processConditionList[i].ConditionString = WAITING // waiting since resources have not been allocated yet
		processConditionList[i].CpuTime = intGen.Int(specifications.BaseCpuTimePart*10, specifications.BaseCpuTimePart*10)
		processConditionList[i].RamSize = intGen.Int(vitualAddrSpaceSize/2, vitualAddrSpaceSize)
		processConditionList[i].ClockPointer = 0 // start
		// 3. init page table
		processPageCount := processConditionList[i].RamSize / pageSizeB
		if processConditionList[i].RamSize%pageSizeB != 0 {
			processPageCount++
		}
		// add virtual pages
		for j := 0; j < processPageCount; j++ {
			processConditionList[i].ProcessPages = append(processConditionList[i].ProcessPages, PageInfo{nil})
		}
		// create page table with real pages
		for j := 0; j < realPageCount; j++ {
			processConditionList[i].PageTable = append(processConditionList[i].PageTable, Page{
				-1,
				nil,
				0,
			})
		}
		// 4. reports
		processConditionList[i].ExecutingTime = 0
		processConditionList[i].WaitingTime = 0
	}

	return processConditionList
}

func get_uuid(length int) string {
	return strings.Replace(uuid.New().String(), "-", "", -1)[:length]
}