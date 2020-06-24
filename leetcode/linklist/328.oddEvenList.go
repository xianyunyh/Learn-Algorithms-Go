package linklist

// 将链表分成两个链表，然后头尾相连接
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddHead, odd, even, evenHead := head, head, head.Next, head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return oddHead
}
