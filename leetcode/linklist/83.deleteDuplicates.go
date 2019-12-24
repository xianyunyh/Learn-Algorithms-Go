package linklist

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	if current == nil || current.Next == nil {
		return head
	}
	for current != nil && current.Next != nil {
		next := current.Next
		if current.Val == next.Val {
			current.Next = next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
