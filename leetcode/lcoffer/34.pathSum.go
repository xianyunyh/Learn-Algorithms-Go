package lcoffer

var path []int

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	recur2(root, sum, &result)
	return result
}

func recur2(root *TreeNode, target int, result *[][]int) {

	if root == nil {
		return
	}
	path = append(path, root.Val)
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	recur2(root.Left, target, result)
	recur2(root.Right, target, result)
	path = path[:len(path)-1]
}
