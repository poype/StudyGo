// https://leetcode.cn/problems/add-two-numbers/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{Val: -1, Next: nil}
	l := head
	for l1 != nil && l2 != nil {
		sum := carry + l1.Val + l2.Val
		carry = sum / 10
		l.Next = &ListNode{Val: sum % 10, Next: nil}
		l = l.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		l1 = l2
	}

	for l1 != nil {
		sum := carry + l1.Val
		carry = sum / 10
		l.Next = &ListNode{Val: sum % 10, Next: nil}
		l = l.Next
		l1 = l1.Next
	}

	if carry > 0 {
		l.Next = &ListNode{Val: carry, Next: nil}
	}

	return head.Next
}

var node13 = &ListNode{Val: 3, Next: nil}
var node12 = &ListNode{Val: 4, Next: node13}
var node11 = &ListNode{Val: 2, Next: node12}

var node23 = &ListNode{Val: 5, Next: nil}
var node22 = &ListNode{Val: 6, Next: node23}
var node21 = &ListNode{Val: 4, Next: node22}
