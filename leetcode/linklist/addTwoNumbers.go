package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

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
