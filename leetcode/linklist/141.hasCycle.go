package linklist

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != slow {
		if fast == nil || slow == nil {
			return false
		}
		head = head.Next
		fast = fast.Next
	}
	return true
}

//利用hash
func hasCycle2(head *ListNode) bool {
	set := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := set[head]; ok {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return false
}
