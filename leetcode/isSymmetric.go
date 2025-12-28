package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	rightTree := invertTree(root.Right)

	return isSameTree(root.Left, rightTree)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
