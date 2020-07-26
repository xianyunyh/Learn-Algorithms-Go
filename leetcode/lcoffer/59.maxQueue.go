package lcoffer

// 单调队列
type MaxQueue struct {
	queue    []int
	maxQueue []int
}

func Constructor2() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxQueue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.maxQueue) > 0 && this.maxQueue[len(this.maxQueue)-1] < value {
		this.maxQueue = this.maxQueue[:len(this.maxQueue)-1]
	}
	this.maxQueue = append(this.maxQueue, value)

}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	front := this.queue[0]
	if len(this.maxQueue) > 0 && this.maxQueue[0] == front {
		this.maxQueue = this.maxQueue[1:len(this.maxQueue)]
	}
	this.queue = this.queue[1:]
	return front
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
