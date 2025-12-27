package main

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{Val: -1, Next: nil}
	for head != nil {
		p := head.Next
		head.Next = newHead.Next
		newHead.Next = head

		head = p
	}
	return newHead.Next
}
