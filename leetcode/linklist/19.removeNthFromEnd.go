package linklist

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	node := &ListNode{}
	node.Next = head
	fast := node
	slow := node
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow != nil {
		deleteNode := slow.Next
		slow.Next = deleteNode.Next
	}
	return node.Next
}
