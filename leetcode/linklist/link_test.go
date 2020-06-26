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
	l1 := CreateLinkList([]int{1, 2, 4})
	l2 := CreateLinkList([]int{1, 3, 4})
	a.Equal(mergeTwoLists(l1, l2).Next.Val, 1)
	l3 := CreateLinkList([]int{1, 0, 1})
	getDecimalValue(l3)
	hasCycle(l3)
	l4 := CreateLinkList([]int{1, 2, 3, 4, 5})
	removeNthFromEnd(l4, 2)
}

func TestLru(t *testing.T) {

}

func TestMiddleNode(t *testing.T) {
	head := CreateLinkList([]int{1, 2, 3, 4, 5})
	middleNode(head)
}

func TestRemoveDuplicateNodes(t *testing.T) {
	head := CreateLinkList([]int{1, 2, 3, 3, 2, 1})
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

func TestCreateLinkList(t *testing.T) {
	link := CreateLinkList([]int{1, 2, 3, 4})
	t.Log(link)
}
