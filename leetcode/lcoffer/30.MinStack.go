package lcoffer

type MinStack struct {
	elements []int
	min      []int
	lens     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		min:      []int{},
	}
}

func (this *MinStack) Push(x int) {

	this.elements = append(this.elements, x)
	this.lens++
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.elements) == 0 {
		return
	}
	top := this.elements[this.lens-1]
	if top == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.elements = this.elements[:this.lens-1]
	this.lens--

}

func (this *MinStack) Top() int {
	if this.lens == 0 {
		return 0
	}
	return this.elements[this.lens-1]
}

func (this *MinStack) Min() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
