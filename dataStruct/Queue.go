package dataStruct

type node struct {
	next *node
	data string
}
type Queue struct {
	capacity int
	lens int
	head *node
	tail *node
}

func NewQueue(capacity int) *Queue  {
	return &Queue {
		capacity:capacity,
		lens:0,
		head:nil,
		tail:nil,
	}
}

func (q *Queue) EnQueue(data string) bool  {
	if	q.IsFull() {
		return false
	}
	//空的队列
	node := &node{
		next: nil,
		data: data,
	}
	if q.lens == 0 {
		q.head = node
		q.tail = node
	}else {
		q.tail.next = node
		q.tail = node
	}
	q.lens++
	return true
}

func (q *Queue)DeQueue()  string {
	if q.lens == 0 {
		return ""
	}
	q.lens--
	if q.head.next == nil {
		return q.head.data
	}

	data := q.head.data
	q.head = q.head.next

	return data
}

func (q *Queue)IsFull() bool  {
	if q.capacity == q.lens {
		return true
	}
	return false
}

func (q *Queue)Lens()int  {
	return q.lens
}