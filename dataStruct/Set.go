package dataStruct

import "sync"

// set 集合 集合中的元素不重复
type HashSet struct {
	items  map[string]string
	lens   int
	rwLock sync.RWMutex
}

func NewSet() *HashSet {
	return &HashSet{
		items:  make(map[string]string),
		lens:   0,
		rwLock: sync.RWMutex{},
	}
}

func (s *HashSet) Add(item string) bool {
	if s.Exist(item) {
		return false
	}
	s.rwLock.Lock()
	s.items[item] = item
	s.lens++
	s.rwLock.Unlock()
	return true
}
func (s *HashSet) Exist(item string) bool {
	_, ok := s.items[item]
	return ok
}

func (s *HashSet) Lens() int {
	return s.lens
}