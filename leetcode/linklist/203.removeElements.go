package linklist

// 利用虚节点 + 双指针。
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	prev := dummyHead
	current := head
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}
	return dummyHead.Next
}
