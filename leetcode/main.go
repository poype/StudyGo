package main

import "fmt"

func main() {
	one := ListNode{Val: 1, Next: nil}
	two := ListNode{Val: 2, Next: nil}
	three := ListNode{Val: 2, Next: nil}
	four := ListNode{Val: 5, Next: nil}
	one.Next = &two
	two.Next = &three
	three.Next = &four
	result := isPalindrome(&one)
	fmt.Println(result)
}
