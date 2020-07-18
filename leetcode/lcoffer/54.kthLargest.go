package lcoffer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	res := []int{}
	dfs(root, &res)
	return res[k-1]
}

func dfs(root *TreeNode, res *[]int) {
	if root.Right != nil {
		dfs(root.Right, res)
	}
	if root != nil {
		*res = append(*res, root.Val)
	}
	if root.Left != nil {
		dfs(root.Left, res)
	}

}
