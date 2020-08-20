package lcoffer

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		head = head.Next
		return head
	}
	prev := head
	current := head.Next

	for current != nil && current.Val != val {
		prev = current
		current = current.Next
	}
	if current != nil {
		prev.Next = current.Next
	}
	return head
}
