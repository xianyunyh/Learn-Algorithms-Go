package linklist

import "math"

func getDecimalValue(head *ListNode) int {

	node := head
	i := 0
	for node != nil {
		node = node.Next
		i++
	}
	num := 0
	for head != nil {
		i--
		var n float64
		if head.Val == 1 {
			n = math.Pow(float64(2), float64(i))
		}

		num += int(n)
		head = head.Next
	}
	return num
}
