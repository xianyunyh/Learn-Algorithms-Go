package lcoffer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 利用中序遍历，然后存入中间数组中，然后返回第k-1个元素的值
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	res := []int{}
	dfs(root, &res, k)
	return res[k-1]
}

func dfs(root *TreeNode, res *[]int, k int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		dfs(root.Right, res, k)
	}
	*res = append(*res, root.Val)
	if len(*res) == k {
		return
	}
	if root.Left != nil {
		dfs(root.Left, res, k)
	}

}
