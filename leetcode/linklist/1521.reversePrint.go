package linklist

// 利用一个栈
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := []int{}
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[len(arr)-i-1])
	}
	return res
}
