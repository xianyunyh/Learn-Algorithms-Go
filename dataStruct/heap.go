package dataStruct

type Heap struct {
	items []int
	top   int
	lens  int
}

func NewHeap() *Heap {
	return &Heap{
		top:   0,
		lens:  0,
		items: make([]int, 10),
	}
}