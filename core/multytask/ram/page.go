package ram

const (
	R_FLAG uint8 = 0b00000001
)

type PageInfo struct {
	// InDiskAddress string  // may be added
	ImMemoryPage *Page
}

type Page struct {
	VirtualAddress  int
	PhysicalAddress *Block
	Flags           uint8
}

// R-bit methods

func (page *Page) CheckR() bool {
	return page.Flags&R_FLAG == 1
}

func (page *Page) SetR() *Page {
	page.Flags |= R_FLAG
	return page
}

func (page *Page) ResetR() *Page {
	page.Flags &= ^R_FLAG
	return page
}
