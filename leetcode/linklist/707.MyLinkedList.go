package linklist

type MyLinkedList struct {
	Head   *ListNode
	Length int
}

/** Initialize your data structure here. */
func Constructor2() MyLinkedList {
	Head := &ListNode{Val: -1, Next: nil}
	return MyLinkedList{
		Head:   Head,
		Length: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Length {
		return -1
	}
	node := this.Head
	for i := 0; i < index+1; i++ {
		node = node.Next
	}
	return node.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Length, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}
	if index < 0 {
		index = 0
	}
	newNode := &ListNode{
		Val: val,
	}
	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	this.Length += 1
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}
	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	this.Length--
}
