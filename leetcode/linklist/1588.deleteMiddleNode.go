package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 这题的删除中间的节点，其实利用伪删除，将该节点的值换成 该节点Next的值，并将node.Next 指向node.Next.Next

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	// tihuan
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}
