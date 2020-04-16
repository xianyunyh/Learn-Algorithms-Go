package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 8}

	l2 := &ListNode{Val: 0}
	addTwoNumbers(l1, l2)
}
func TestDeleteDuplicates(t *testing.T) {
	// a := assert.New(t)
}
func TestMergeTwoLists(t *testing.T) {
	a := assert.New(t)
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}
	a.Equal(mergeTwoLists(l1, l2).Next.Val, 1)
	l3 := &ListNode{Val: 1}
	l3.Next = &ListNode{Val: 0}
	l3.Next.Next = &ListNode{Val: 1}
	getDecimalValue(l3)
	hasCycle(l3)
	l4 := &ListNode{Val: 1}
	l4.Next = &ListNode{Val: 2}
	l4.Next.Next = &ListNode{Val: 3}
	l4.Next.Next.Next = &ListNode{Val: 4}
	l4.Next.Next.Next.Next = &ListNode{Val: 5}
	removeNthFromEnd(l4, 2)
}

func TestLru(t *testing.T) {

}
