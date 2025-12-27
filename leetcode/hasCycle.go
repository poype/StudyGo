package main

func hasCycle(head *ListNode) bool {
	p, q := head, head
	for p != nil && q != nil {
		p = p.Next
		q = q.Next
		if q == nil {
			return false
		}
		q = q.Next
		if p == q {
			return true
		}
	}
	return false
}
