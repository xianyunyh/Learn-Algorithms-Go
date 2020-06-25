package linklist

// 解题思路： 使用hashSet 存储G，遍历链表，如果链表元素在G中，并且下个节点是空，则计数+1，或者下个节点的值不在G中，也+1/
func numComponents(head *ListNode, G []int) int {
	if head == nil {
		return 0
	}
	set := make(map[int]struct{})

	for _, v := range G {
		set[v] = struct{}{}
	}
	current := head
	num := 0
	for current != nil {

		if _, ok := set[current.Val]; ok {
			if current.Next == nil {
				num += 1
			} else {
				_, ok2 := set[current.Next.Val]
				if !ok2 {
					num += 1
				}
			}
		}
		current = current.Next
	}
	return num
}
