package ram

// Memory Management Unit procedures here

func PrepareRequiredPages(memory *MemoryInfo, pageSizeB int, process *Condition, requiredPages []int) {
	physicalSpacePageCount := len(process.PageTable)
	// for all required pages
	for _, pageIndex := range requiredPages {
		isResident := false
		for j := 0; j < physicalSpacePageCount && !isResident; j++ {
			if process.PageTable[j].VirtualAddress == pageIndex {
				// page in PageTable and in memory
				isResident = true
				process.PageTable[j].SetR() // set R-bit
			}
		}
		if !isResident {
			// page not found in page table
			process = PageFaultHandler(memory, pageSizeB, process, pageIndex)
		}
	}
}

func PageFaultHandler(memory *MemoryInfo, pageSizeB int, process *Condition, pageIndex int) *Condition {
	// 1. first try to find free page block
	freePageIndex := -1
	for i, page := range process.PageTable {
		if page.VirtualAddress == -1 {
			freePageIndex = i
			break
		}
	}

	// 2. add page if true else run substitution algorithm (WSClock)
	if freePageIndex != -1 {
		// add page to empty row in table page
		process.PageTable[freePageIndex].VirtualAddress = pageIndex
		process.PageTable[freePageIndex].SetR()
		process.PageTable[freePageIndex].PhysicalAddress, _ = memory.WritePage(pageSizeB, process.GetPPString(pageIndex))
	} else {
		// use WSClock
		for ; ; process.ClockPointer = (process.ClockPointer + 1) % len(process.PageTable) {
			if !process.PageTable[process.ClockPointer].CheckR() {
				// can remove
				memory.RemovePage(process.PageTable[process.ClockPointer].PhysicalAddress)
				process.PageTable[process.ClockPointer].VirtualAddress = pageIndex
				process.PageTable[process.ClockPointer].SetR()
				process.PageTable[process.ClockPointer].PhysicalAddress, _ = memory.WritePage(pageSizeB, process.GetPPString(pageIndex))
				break
			} else {
				process.PageTable[process.ClockPointer].ResetR()
			}
		}
	}
	return process
}

func PrepareRequiredPagesWithMemoryMetrics(memory *MemoryInfo, pageSizeB int,
	process *Condition, requiredPages []int) (isPageWrote bool, isPageRemoved bool) {
	isPageWrote = false
	isPageRemoved = false
	physicalSpacePageCount := len(process.PageTable)
	// for all required pages
	for _, pageIndex := range requiredPages {
		isResident := false
		for j := 0; j < physicalSpacePageCount && !isResident; j++ {
			if process.PageTable[j].VirtualAddress == pageIndex {
				// page in PageTable and in memory
				isResident = true
				process.PageTable[j].SetR() // set R-bit
			}
		}
		if !isResident {
			// page not found in page table
			process, isPageRemoved = PageFaultHandlerWithMemoryMetrics(memory, pageSizeB, process, pageIndex)
			isPageWrote = true
		}
	}

	return isPageWrote, isPageRemoved
}

func PageFaultHandlerWithMemoryMetrics(memory *MemoryInfo, pageSizeB int, process *Condition, pageIndex int) (*Condition, bool) {
	isPageRemoved := false

	// 1. first try to find free page block
	freePageIndex := -1
	for i, page := range process.PageTable {
		if page.VirtualAddress == -1 {
			freePageIndex = i
			break
		}
	}

	// 2. add page if true else run substitution algorithm (WSClock)
	if freePageIndex != -1 {
		// add page to empty row in table page
		process.PageTable[freePageIndex].VirtualAddress = pageIndex
		process.PageTable[freePageIndex].SetR()
		process.PageTable[freePageIndex].PhysicalAddress, _ = memory.WritePage(pageSizeB, process.GetPPString(pageIndex))
	} else {
		// use WSClock
		for ; ; process.ClockPointer = (process.ClockPointer + 1) % len(process.PageTable) {
			if !process.PageTable[process.ClockPointer].CheckR() {
				// can remove
				memory.RemovePage(process.PageTable[process.ClockPointer].PhysicalAddress)
				isPageRemoved = true
				process.PageTable[process.ClockPointer].VirtualAddress = pageIndex
				process.PageTable[process.ClockPointer].SetR()
				process.PageTable[process.ClockPointer].PhysicalAddress, _ = memory.WritePage(pageSizeB, process.GetPPString(pageIndex))
				break
			} else {
				process.PageTable[process.ClockPointer].ResetR()
			}
		}
	}
	return process, isPageRemoved
}
