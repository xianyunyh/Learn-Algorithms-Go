package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(arr []int) *ListNode {
	head := &ListNode{}
	node := head

	for _, v := range arr {
		newNode := &ListNode{
			Val:  v,
			Next: nil,
		}
		node.Next = newNode
		node = newNode
	}
	return head.Next
}
