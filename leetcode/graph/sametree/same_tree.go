package sametree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pq := []*TreeNode{p}
	qq := []*TreeNode{q}

	for len(pq) > 0 && len(qq) > 0 {
		pNode := pq[0]
		pq = pq[1:] // pop first
		qNode := qq[0]
		qq = qq[1:] // pop first

		if pNode == nil {
			if qNode == nil {
				continue
			}
			return false
		}

		if pNode.Val != qNode.Val {
			return false
		}

		pq = append(pq, pNode.Left)
		pq = append(pq, pNode.Right)
		qq = append(qq, qNode.Left)
		qq = append(qq, qNode.Right)
	}

	if len(pq) > 0 || len(qq) > 0 {
		return false
	}

	return true
}
