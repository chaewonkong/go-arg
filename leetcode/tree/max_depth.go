package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return walk(root, 0)
}

func walk(node *TreeNode, cnt int) int {
	if node == nil {
		return cnt
	}
	cnt++

	return max(walk(node.Left, cnt), walk(node.Right, cnt))
}
