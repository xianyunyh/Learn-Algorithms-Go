package hash

type MyHashSet struct {
	items []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {

	return MyHashSet{
		items: make([]int, 0),
	}
}

func (this *MyHashSet) Add(key int) {
	if len(this.items) == 0 {
		this.items = append(this.items, key)
	} else {
		if this.Contains(key) {
			return
		}
		this.items = append(this.items, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	if len(this.items) == 0 {
		return
	}
	idx := -1
	for i, v := range this.items {
		if v == key {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	this.items = append(this.items[:idx], this.items[idx+1:]...)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if len(this.items) == 0 {
		return false
	}
	for _, v := range this.items {
		if v == key {
			return true
		}
	}
	return false
}
