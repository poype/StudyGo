package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lastNodeA, lengthA := getLastNodeAndLength(headA)
	lastNodeB, lengthB := getLastNodeAndLength(headB)
	if lastNodeA != lastNodeB {
		return nil
	}
	if lengthA < lengthB {
		lengthA, lengthB = lengthB, lengthA
		headA, headB = headB, headA
	}
	diffSteps := lengthA - lengthB
	for diffSteps > 0 {
		headA = headA.Next
		diffSteps--
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func getLastNodeAndLength(head *ListNode) (*ListNode, int) {
	if head == nil {
		return nil, 0
	}

	listLength := 1
	for head.Next != nil {
		listLength++
		head = head.Next
	}
	return head, listLength
}
