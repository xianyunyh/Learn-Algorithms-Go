package linklist

//解题思路 ： 遍历两个链表，让两个链表的值相加。让商=carry 下一次的计算结果 加上carray，
// 如果已经结束，carray 还存在，说明进位了。补1
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	current := pre
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next
	}
	if carry == 1 {
		current.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return pre.Next
}
