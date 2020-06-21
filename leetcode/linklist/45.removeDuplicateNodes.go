package linklist

// 题解： 利用hash表 + 双指针，哈希表存储每个节点的值。
// 两个指针。一个指向前节点，一个执行当前节点，当哈希表值存在，当前节点前进两步。
func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]struct{})
	dummyHead := &ListNode{}
	pre := dummyHead
	for head != nil {
		if _, ok := m[head.Val]; !ok {
			m[head.Val] = struct{}{}
			pre.Next = head
			pre = head
		} else {
			pre.Next = head.Next
		}
		head = head.Next
	}
	return dummyHead.Next
}
