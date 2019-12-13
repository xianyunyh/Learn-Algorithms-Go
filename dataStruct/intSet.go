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

func (is *IntSet) Insert(value int64) bool {
	var pos int = -1
	if is.Search(value, &pos) {
		return false
	}
	if pos < 0 {
		return false
	}
	is.MoveTail(pos, is.lens+1)
	is.items[pos] = value
	is.lens++
	return true
}

func (is *IntSet) Find(pos int) int64 {
	if pos < 0 {
		return -1
	}
	return  is.items[pos]
}

func (is *IntSet) Search(value int64, pos *int) bool {
	var setPos = func(idx int, pos *int) {
		if pos != nil {
			*pos = idx
		}
	}
	//长度为0
	if is.lens == 0 {
		setPos(0, pos)
		return false
	} else {
		//不在范围内
		if value < is.items[0] {
			setPos(0, pos)
			return false
		} else if value > is.items[is.lens-1] {
			setPos(is.lens, pos)
			return false
		}
	}
	max, min := 0, is.lens-1
	var mid int = -1
	var cur int64 = -1
	for max > min {
		mid = (max + min) >> 1
		cur = is.Find(mid)
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

func (is *IntSet) MoveTail(from, to int) {
	is.items = append(is.items,0)
	for to > from {
		is.items[to] = is.items[to-1]
		to--
	}
}

func (is *IntSet) Remove(value int64) bool {
	var pos int  = -1
	if !is.Search(value,&pos) {
		return false
	}
	is.items = append(is.items[:pos], is.items[pos:]...)
	return true
}
func (is *IntSet) Lens() int {
	return is.lens
}
