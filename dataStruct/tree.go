package dataStruct

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

//二叉搜索树 的插入
func Append(root *TreeNode, val int) {
	if root == nil {
		NewNode(val)
		return
	}
	var parent *TreeNode
	for root != nil {
		parent = root
		if root.Val > val {
			root = root.Right
		} else if root.Val > val {
			root = root.Left
		} else {
			return
		}
	}
	if parent.Val > val {
		parent.Left = NewNode(val)
	} else {
		parent.Right = NewNode(val)
	}
}

//二叉搜索树的树递归版本插入
func appendRecurrence(root *TreeNode, val int) {
	if root == nil {
		NewNode(val)
		return
	}
	if root.Val > val {
		appendRecurrence(root.Right, val)
		return
	}
	if root.Val < val {
		appendRecurrence(root.Left, val)
		return
	}
}

//二叉搜索树的查找
func Search(root *TreeNode, val int) bool {

	for root != nil {
		if root.Val == val {
			return true
		}
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return false
}

func SearchRecurrence(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val > val {
		return SearchRecurrence(root.Left, val)
	}
	if root.Val < val {
		return SearchRecurrence(root.Right, val)
	}
	return true
}

// 前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树 根->左->右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = node.Right
		}

	}
	return result
}

// 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树 左->根->右
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		} else {
			// 弹出
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			root = node.Right
		}
	}
	return result
}

//递归版本
func preorderTraversalRec(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点 左->右->根
func postorderTraversal(root *TreeNode) []int {
	// 通过lastVisit标识右子节点是否已经弹出
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

func levelOrder2(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 为什么要取length？
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

//层级遍历  维护一个队列 二叉树的层级遍历 又叫广度遍历
// 先在队列中加入根结点。之后对于任意一个结点来说，在其出队列的时候，访问之。同时如果左孩子和右孩子有不为空的，入队列
func levelOrder(root *TreeNode) []int {
	queue := []*TreeNode{}
	result := []int{}
	if root == nil {
		return result
	}
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return result
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

//对称树
// 左边的节点 == 右边的节点值
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return rec(root.Left, root.Right)
}
func rec(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return rec(left.Left, right.Right) && rec(left.Right, right.Left)

}
