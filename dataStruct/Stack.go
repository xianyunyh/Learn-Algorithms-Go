package dataStruct

type Stack struct {
	lens     int
	capacity int
	items    []string
}

func NewStack(capacity int) *Stack {
	return &Stack{
		lens:     0,
		capacity: capacity,
		items:    make([]string, capacity),
	}
}

// 入栈
func (s *Stack) Push(item string) {
	if s.IsFull() {
		return
	}
	s.items[s.lens] = item
	s.lens++
}

func (s *Stack) IsFull() bool {
	return s.lens >= s.capacity
}

// 出栈
func (s *Stack) Pop() string {
	item := s.items[s.lens-1]
	s.items = s.items[:s.lens-1]
	s.lens--
	return item
}
