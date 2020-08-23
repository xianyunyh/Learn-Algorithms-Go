package lcoffer

func getKthFromEnd(head *ListNode, k int) *ListNode {

	fast := head
	slow := head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
