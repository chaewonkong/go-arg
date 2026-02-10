package tree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	// 현재 노드가 범위를 벗어나는지 확인
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	// 왼쪽 서브트리: 최대값을 현재 노드로 제한
	// 오른쪽 서브트리: 최소값을 현재 노드로 제한
	return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
}
