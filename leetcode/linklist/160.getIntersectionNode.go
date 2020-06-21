package linklist

// A.B 同时走，如果A、
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	currentA := headA
	currentB := headB
	for currentA != currentB {
		if currentA == nil {
			currentA = headB
		} else {
			currentA = currentA.Next
		}
		if currentB == nil {
			currentB = headA
		} else {
			currentB = currentB.Next
		}

	}
	return currentA
}
