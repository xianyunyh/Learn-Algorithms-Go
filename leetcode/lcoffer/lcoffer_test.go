package lcoffer

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	spiralOrder([][]int{
		{2, 5, 8},
		{4, 0, -1},
	})
}

func TestSlice(t *testing.T) {
	str := "hello world"
	t.Log(str[6:12])
}

func TestLengthOfLongestSubstring(t *testing.T) {
	a := lengthOfLongestSubstring("abcabcbb")
	if a != 3 {
		t.Fatalf("%s", "failed")
	}
}

func TestTreeToDoublyList(t *testing.T) {
	root := &TreeNode{4, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{5, nil, nil}
	node3 := &TreeNode{1, nil, nil}
	node4 := &TreeNode{3, nil, nil}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	head := treeToDoublyList(root)
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", head.Val)
		head = head.Right
	}
}
