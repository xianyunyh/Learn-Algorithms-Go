package lcoffer

// 计算左树和右树的高度
func isBalanced(root *TreeNode) bool {
	return !(maxDepth(root) == -1)
}
func maxDepth(root *TreeNode) int {
	// check
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// 为什么返回-1呢？（变量具有二义性）
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
