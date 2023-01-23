package ram

type MemoryInfo struct {
	MemoryBlockSizeB    int    // in bytes
	MemorySpaceSizeB    int    // in bytes
	NextFreeMemoryBlock *Block // always free
}

func (memory *MemoryInfo) WritePage(pageSizeB int, fillerData string) (*Block, error) {
	var pagePhysicalAddress *Block = nil
	if !memory.NextFreeMemoryBlock.IsFree {
		// because it should be free always
		panic("Not free memory.NextFreeMemoryBlocks")
	} else {
		pagePhysicalAddress = memory.NextFreeMemoryBlock
	}

	requiredBlocksCount := pageSizeB / memory.MemoryBlockSizeB
	if pageSizeB%memory.MemoryBlockSizeB != 0 {
		requiredBlocksCount++
	}

	var prevInGroup *Block = nil
	var err error
	for i := 0; i < requiredBlocksCount; i++ {
		// 1. write data to NextFreeMemoryBlock
		memory.NextFreeMemoryBlock, err = memory.NextFreeMemoryBlock.WriteStringToBlock(fillerData)
		if err != nil {
			return nil, err
		}
		memory.NextFreeMemoryBlock.IsFree = false
		// 2. update next in group
		if prevInGroup != nil {
			prevInGroup.NextInGroup = memory.NextFreeMemoryBlock
		}
		// 3. set prev block in group
		prevInGroup = memory.NextFreeMemoryBlock
		// 4. find next free memory block
		for memory.NextFreeMemoryBlock = memory.NextFreeMemoryBlock.Next; !memory.NextFreeMemoryBlock.IsFree; memory.NextFreeMemoryBlock = memory.NextFreeMemoryBlock.Next {
		}
	}
	return pagePhysicalAddress, nil
}

func (memory *MemoryInfo) RemovePage(pagePhysicalAddress *Block) {
	var curBlock *Block = pagePhysicalAddress
	var oldCurBlock *Block = nil
	for curBlock != nil {
		curBlock.IsFree = true
		if oldCurBlock != nil {
			oldCurBlock.NextInGroup = nil
		}
		oldCurBlock = curBlock
		curBlock = curBlock.NextInGroup
	}
}

func CreateMemoryStructure(memoryBlockSizeB int, maxProcessNumb int, processRealPageCount int, pageSizeB int) MemoryInfo {
	osSpace := 0 // temporary use 0
	memoryInfo := MemoryInfo{
		memoryBlockSizeB,
		maxProcessNumb*processRealPageCount*pageSizeB + osSpace,
		nil,
	}
	memoryBlockCount := memoryInfo.MemorySpaceSizeB / memoryInfo.MemoryBlockSizeB
	// next init ram blocks
	memoryInfo.NextFreeMemoryBlock = CreateEmptyFreeBlock(memoryBlockSizeB)
	var curBlock *Block
	var i int
	for i, curBlock = 1, memoryInfo.NextFreeMemoryBlock; i < memoryBlockCount; i, curBlock = i+1, curBlock.Next {
		curBlock.Next = CreateEmptyFreeBlock(memoryBlockSizeB)
		curBlock.Next.Prev = curBlock
	}
	// make list as cycle
	curBlock.Next = memoryInfo.NextFreeMemoryBlock
	memoryInfo.NextFreeMemoryBlock.Prev = curBlock.Next
	//
	return memoryInfo
}
