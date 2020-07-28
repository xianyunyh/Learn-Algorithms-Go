package lcoffer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 解题思路，利用hash表记录每个链表的映射，然后再遍历链表
// 找到hash表中链表的映射，并修改
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	table := make(map[*Node]*Node)
	for temp := head; temp != nil; temp = temp.Next {
		table[temp] = &Node{
			Val:    temp.Val,
			Next:   nil,
			Random: nil,
		}
	}
	for temp := head; temp != nil; temp = temp.Next {
		node := table[temp]
		if temp.Next != nil {
			node.Next = table[temp.Next]
		}
		if temp.Random != nil {
			node.Random = table[temp.Random]
		}
		table[temp] = node
	}
	return table[head]
}
