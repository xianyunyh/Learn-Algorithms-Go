package linklist

func middleNode(head *ListNode) *ListNode {
	lens := 0
	middle := 0
	current := head
	for current != nil {
		lens += 1
		current = current.Next
	}
	// 是不是偶数
	middle = (lens / 2) + 1
	current = head
	i := 1
	for current != nil {
		if i == middle {
			break
		}
		current = current.Next
		i = i + 1
	}
	return current
}
