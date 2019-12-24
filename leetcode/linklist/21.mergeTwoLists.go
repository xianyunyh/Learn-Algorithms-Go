package linklist

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	prev := head
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}
	prev.Next = l1
	if l1 == nil {
		prev.Next = l2
	}
	return head.Next
}
