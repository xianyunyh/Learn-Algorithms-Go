package linklist

// 哑巴结点 + 双指针
// 定义prev指向前结点，node 指向当前， 如果node > prev 则继续往前
// 如果node< prev 则从头结点依次比较
// 插入结点 修改各结点指向
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dumyNode := &ListNode{Val: -1}
	dumyNode.Next = head
	prev := head
	node := head.Next
	for node != nil {
		if node.Val > prev.Val {
			prev = prev.Next
			node = node.Next
			continue
		}
		temp := dumyNode
		for temp.Next.Val < node.Val {
			temp = temp.Next
		}
		prev.Next = node.Next
		node.Next = temp.Next
		temp.Next = node
		node = prev.Next

	}
	return dumyNode.Next
}
