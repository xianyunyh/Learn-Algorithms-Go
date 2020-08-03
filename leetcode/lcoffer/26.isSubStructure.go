package lcoffer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var ret bool

	//当在 A 中找到 B 的根节点时，进入helper递归校验
	if A.Val == B.Val {
		ret = recur(A, B)
	}
	if !ret {
		ret = isSubStructure(A.Left, B)
	}

	if !ret {
		ret = isSubStructure(A.Right, B)
	}
	return ret
}
func recur(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}

	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}
