package ram

import (
	"errors"
)

// Block of memory
// Doubly linked list
type Block struct {
	Next        *Block
	Prev        *Block
	IsFree      bool
	NextInGroup *Block // next block in group of page blocks
	//
	bytes []byte // emulation of memory data in this block
}

func CreateEmptyFreeBlock(memoryBlockSizeB int) *Block {
	block := new(Block)
	block.IsFree = true
	block.Next, block.Prev, block.NextInGroup = nil, nil, nil
	block.bytes = make([]byte, memoryBlockSizeB)
	block, err := block.WriteStringToBlockAsASCII("empty")
	if err != nil {
		panic(err) // fatal error
	}
	return block
}

func (block *Block) WriteStringToBlock(data string) (*Block, error) {
	return block.WriteBytesToBlock([]byte(data))
}

func (block *Block) WriteStringToBlockAsASCII(data string) (*Block, error) {
	return block.WriteBytesToBlock([]byte(data))
}

func (block *Block) WriteBytesToBlock(bytes []byte) (*Block, error) {
	if len(bytes) > len(block.bytes) {
		return nil, errors.New("not enough space to write in memory block")
	}
	copy(block.bytes, bytes)
	return block, nil
}

func (block *Block) GetBytes() []byte {
	return block.bytes
}
