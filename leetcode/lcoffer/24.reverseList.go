package lcoffer

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    prev := head 
    var current  *ListNode
    for prev != nil {
        t  := prev.Next
        prev.Next = current
        current = prev
        prev = t
    }
    return current
}