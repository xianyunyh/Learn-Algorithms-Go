package dataStruct

type IntSet struct {
	lens  int
	items []int64
}

func NewIntSet() *IntSet {
	return &IntSet{
		items: make([]int64, 0),
		lens:  0,
	}
}

func (self *IntSet) Insert(value int64) bool {
	var pos int = 0
	if self.Search(value, &pos) {
		return false
	}
	return false
}

func (self *IntSet) Find(pos int) int64 {

	return 0
}

func (self *IntSet) Search(value int64, pos *int) bool {
	var setPos = func(idx int, pos *int) {
		if pos != nil {
			*pos = idx
		}
	}
	//长度为0
	if self.lens == 0 {
		setPos(0, pos)
		return false
	} else {
		//不在范围内
		if value < self.items[0] {
			setPos(0, pos)
			return false
		} else if value > self.items[self.lens-1] {
			setPos(self.lens, pos)
			return false
		}
	}
	max, min := 0, self.lens-1
	var mid int = -1
	var cur int64 = -1
	for max > min {
		mid = int((max + min) >> 1)
		cur = self.Find(mid)
		if cur > value {
			min = mid + 1
		} else if cur < value {
			min = mid - 1
		} else {
			break
		}
	}
	if cur == value {
		setPos(mid, pos)
		return true
	}
	setPos(min, pos)
	return false

}

func (self *IntSet) MoveTail(from, to int) {

}

func (self *IntSet) Remove(value int64) bool {

	return true
}
func (self *IntSet) Lens() int {
	return self.lens
}
