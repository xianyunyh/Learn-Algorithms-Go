package linklist

// 快慢指针，快指针先走K步，然后快慢指针一起走，等到快指针走到结束，慢指针就是倒数第k个
func getKthFromEnd(head *ListNode, k int) *ListNode {
    fast := head
    slow := head
    for ;k >0 ;k-- {
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}