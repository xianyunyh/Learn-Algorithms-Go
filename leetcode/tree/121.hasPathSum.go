package tree

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	nodeStack := []*TreeNode{}
	sumStack := []int{}
	nodeStack = append(nodeStack, root)
	sumStack = append(sumStack, root.Val)
	for len(nodeStack) > 0 {
		top := nodeStack[len(nodeStack)-1]
		temp := sumStack[len(sumStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		sumStack = sumStack[:len(sumStack)-1]
		if top.Left == nil && top.Right == nil && sum == temp {
			return true
		}
		if top.Left != nil {
			nodeStack = append(nodeStack, top.Left)
			sumStack = append(sumStack, temp+top.Left.Val)
		}
		if top.Right != nil {
			nodeStack = append(nodeStack, top.Right)
			sumStack = append(sumStack, temp+top.Right.Val)
		}
	}
	return false
}
