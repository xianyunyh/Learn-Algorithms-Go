package linklist

//解题思路
// 利用数学，设head到环的起点为M,环的一周为N，环的起点到两个指针的交点为Y,快指针比慢指针多跑2倍，则有以下公式
// M + x*N + Y = (M+Y) * 2 => M+Y = XN => M= N-Y + (X-1)*n
func DetectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
