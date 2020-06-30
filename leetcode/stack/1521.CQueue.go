package stack

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
	if len(this.right) != 0 {
		res = this.right[0]
		if len(this.right) == 1 {
			this.right = []int{}
		} else {
			this.right = this.right[1:]
		}
		return res
	}

	for i := len(this.left) - 1; i > 1; i-- {
		this.right = append(this.right, this.left[i])
	}
	res = this.left[0]
	this.left = []int{}
	return res
}
