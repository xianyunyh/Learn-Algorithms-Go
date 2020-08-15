package tree

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 总是选择中间位置右边的数字作为根节点
	mid := (left + right + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = dfs(nums, left, mid-1)
	root.Right = dfs(nums, mid+1, right)
	return root
}
