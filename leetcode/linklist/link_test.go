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

func TestMiddleNode(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	middleNode(head)
}

func TestRemoveDuplicateNodes(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	removeDuplicateNodes(head)
}

func TestMyLinkList(t *testing.T) {
	// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
	// [[],[1],[3],[1,2],[1],[1],[1]]
	obj := Constructor2()
	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	obj.Get(5)
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)
	param_1 := obj.Get(1)
	obj.DeleteAtIndex(1)
	t.Log(param_1)
}
