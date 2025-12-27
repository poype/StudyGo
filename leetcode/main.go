package main

import "fmt"

func main() {
	one := TreeNode{Val: 1, Left: nil, Right: nil}
	two := TreeNode{Val: 2, Left: nil, Right: nil}
	three := TreeNode{Val: 3, Left: nil, Right: nil}
	four := TreeNode{Val: 4, Left: nil, Right: nil}
	five := TreeNode{Val: 5, Left: nil, Right: nil}
	six := TreeNode{Val: 6, Left: nil, Right: nil}
	seven := TreeNode{Val: 7, Left: nil, Right: nil}
	eight := TreeNode{Val: 8, Left: nil, Right: nil}
	nine := TreeNode{Val: 9, Left: nil, Right: nil}
	one.Left = &two
	one.Right = &three
	two.Left = &four
	two.Right = &five
	five.Left = &six
	five.Right = &seven
	three.Right = &eight
	eight.Left = &nine
	result := inorderTraversal(&one)
	fmt.Println(result)
}
