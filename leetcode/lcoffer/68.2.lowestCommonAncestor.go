package lcoffer

// 二叉搜索树的最近公共祖先，迭代，
// 如果根节点大于两个节点的值，则两个元素在二叉树的左边 root =root.left
// 如果根节点小于两个节点的值 两个元素在二叉树的右边 root = root.right
// 否则该节点就是最近祖先
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}
