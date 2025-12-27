package main

func isPalindrome(head *ListNode) bool {
	length := getLength(head)
	halfLength := length / 2

	// 这里的len值如果大于0，会用nil指针占位
	stack := make([]*ListNode, 0, halfLength)
	for halfLength > 0 {
		stack = append(stack, head)
		head = head.Next
		halfLength--
	}

	if length%2 == 1 {
		head = head.Next
	}

	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		if topNode.Val != head.Val {
			return false
		}
		stack = stack[:len(stack)-1]
		head = head.Next
	}
	return true
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
