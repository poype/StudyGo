package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	trace := make([]int, 0)
	if root == nil {
		return trace
	}
	inorderTraversalWithTrace(root, &trace)
	return trace
}

func inorderTraversalWithTrace(root *TreeNode, trace *[]int) {
	if root.Left != nil {
		inorderTraversalWithTrace(root.Left, trace)
	}
	*trace = append(*trace, root.Val)
	if root.Right != nil {
		inorderTraversalWithTrace(root.Right, trace)
	}
}
