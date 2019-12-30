package tree

// 解题思路
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/solution/tu-jie-104-er-cha-shu-de-zui-da-shen-du-di-gui-pyt/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
