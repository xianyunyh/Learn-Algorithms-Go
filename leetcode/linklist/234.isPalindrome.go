package linklist

// 解题思路：快慢指针+翻转链表，
// 1. 快指针比慢指针多走一倍，那么快指针到终点，慢指针在中间，同时翻转前半个链表
// 2. 然后两个指针 前后一起跑，如果两个值不一样则不是回文
func isPalindrome(head *ListNode) bool {
	//快慢遍历
	slow := head
	fast := head
	var prev *ListNode
	var current *ListNode
	for fast != nil && fast.Next != nil {
		current = slow
		slow = slow.Next
		fast = fast.Next.Next

		current.Next = prev //翻转
		prev = current
	}
	if fast != nil {
		slow = slow.Next
	}
	for current != nil {
		if current.Val != slow.Val {
			return false
		}
		current = current.Next
		slow = slow.Next
	}
	return true
}
