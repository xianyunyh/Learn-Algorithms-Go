package lcoffer

type CQueue struct {
	left  []int
	right []int
}

func Constructor() CQueue {
	return CQueue{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.left = append(this.left, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.left) == 0 && len(this.right) == 0 {
		return -1
	}
	res := -1
	// 右边有数据 从右边出
	if len(this.right) > 0 {
		res = this.right[len(this.right)-1]
		this.right = this.right[:len(this.right)-1]
		return res
	}
	if len(this.left) == 0 {
		return -1
	}
	for i := len(this.left) - 1; i >= 0; i-- {
		this.right = append(this.right, this.left[i])
	}
	this.left = []int{}
	res = this.right[len(this.right)-1]
	this.right = this.right[:len(this.right)-1]
	return res
}
