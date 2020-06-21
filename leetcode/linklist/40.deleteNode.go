package linklist

// 需要考虑三种情况
// 1. 是头节点，直接让头节点指向下个节点
// 2. 尾节点，  让尾节点的上节点指向nil
// 3. 中间节点 直接 修改节点的值，并修改节点next指向

func DeleteNode(head *ListNode, val int) *ListNode {
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
