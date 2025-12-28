package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		root.Left = invertTree(root.Left)
	}
	if root.Right != nil {
		root.Right = invertTree(root.Right)
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}
