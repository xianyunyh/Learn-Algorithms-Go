package linklist

//解题思路： K % lens 寻找断开的节点的位置，然后拼成环
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	left := head
	lens := 0
	for left != nil {
		left = left.Next
		lens += 1
	}
	// 当K等于N的倍数，则链表复原
	n := k % lens
	if n == 0 {
		return head
	}
	n = lens - n - 1

	left = head
	// 得到左断开的链表节点
	for i := 0; i < n; i++ {
		left = left.Next
	}
	rightHead := left.Next
	rightTail := head
	// 链表尾部
	for rightTail.Next != nil {
		rightTail = rightTail.Next
	}

	left.Next = nil
	rightTail.Next = head
	return rightHead
}
