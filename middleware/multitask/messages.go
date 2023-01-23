package multitask

type InitialMemoryArgs struct {
	MemoryBlockSizeB     int `json:"memory_block_size_b"`
	MaxProcessNumb       int `json:"max_process_numb"`
	ProcessRealPageCount int `json:"process_real_page_count,omitempty"`
	PageSizeB            int `json:"page_size_b"`
}
