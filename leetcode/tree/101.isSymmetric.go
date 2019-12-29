package tree

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}
func isMirror(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}
