package multitask

import "simulator/core/multytask/ram"

type MemoryBlockDump struct {
	IsFree bool   `json:"is_free"`
	Bytes  string `json:"bytes"`
}

type MemoryDump struct {
	MemoryBlockSizeB int               `json:"memory_block_size_b"`
	MemorySpaceSizeB int               `json:"memory_space_size_b"`
	MemoryBlocks     []MemoryBlockDump `json:"memory_blocks"`
}

func GenerateMemoryDump(memoryInfo *ram.MemoryInfo) MemoryDump {
	memoryBlocksCount := memoryInfo.MemorySpaceSizeB / memoryInfo.MemoryBlockSizeB
	memoryDump := MemoryDump{
		memoryInfo.MemoryBlockSizeB,
		memoryInfo.MemorySpaceSizeB,
		make([]MemoryBlockDump, memoryBlocksCount),
	}
	// write blocks info
	for i, memoryBlock := 0, memoryInfo.NextFreeMemoryBlock; i < memoryBlocksCount; i, memoryBlock = i+1, memoryBlock.Next {
		memoryDump.MemoryBlocks[i].IsFree = memoryBlock.IsFree
		memoryDump.MemoryBlocks[i].Bytes = string(memoryBlock.GetBytes())
	}

	return memoryDump
}
