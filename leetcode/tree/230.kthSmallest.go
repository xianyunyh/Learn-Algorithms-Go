package tree

//中序遍历 就是递增
func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	inOrder(root, &arr)
	return arr[k-1]
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
