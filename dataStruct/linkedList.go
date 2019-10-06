package dataStruct

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head *Node
	tail *Node
	lens int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		tail: nil,
		lens: 0,
	}
}

func (l *LinkedList) Append(data string) *LinkedList {
	//空链表
	node := &Node{
		Data: data,
		Next: nil,
	}
	if l.lens == 0 {
		l.Head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node;
	}
	l.lens ++
	return l;
}

func (l *LinkedList)Search(data string) bool  {
	if l.lens == 0 {
		return false
	}
	node := l.Head
	result := false
	for node != nil  {
		if node.Data == data {
			result = true
			break
		}
		node = node.Next
	}
	return result
}

//删除节点
func (l *LinkedList)Remove(data string) bool  {
	if l.lens == 0 {
		return false
	}
	node := l.Head
	//头节点
	if node.Data == data {
		l.Head = node.Next
		l.lens--
		return true
	}
	for node != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
			l.lens--
			return true
		} else {
			node = node.Next
		}
	}
	return false;
}

// 删除指定节点 O(1) 节点不是尾节点
func (l *LinkedList)RemoveNode(node *Node)  {
	nextNode := node.Next
	if nextNode != nil {
		node.Data = nextNode.Data
		node.Next = nextNode.Next.Next
	} else if l.Head == node && l.lens == 1 {
		l.Head = nil
		l.tail = nil
		l.lens = 0
	} else  {
		tail := l.Head
		for tail.Next != node {
			tail = tail.Next
		}
		tail.Next = nil
	}

}

func (l *LinkedList)Lens() int  {
	return l.lens
}