package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth, rightDepth := maxDepth(root.Left), maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
