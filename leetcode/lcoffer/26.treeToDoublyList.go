package lcoffer

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	nodes := []*TreeNode{}
	dfsNode(root, &nodes)
	head := nodes[0]
	current := head
	var prev *TreeNode = head
	for i := 0; i < len(nodes); i++ {

		current = nodes[i]
		if i == len(nodes)-1 {
			current.Right = head
			current.Left = prev
			break
		}
		current.Left = prev
		current.Right = nodes[i+1]
		prev = current
	}
	head.Left = current
	return head
}

func dfsNode(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	dfsNode(root.Left, nodes)
	*nodes = append(*nodes, root)
	dfsNode(root.Right, nodes)
}
