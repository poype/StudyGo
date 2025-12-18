package main

import "fmt"

type ListNode[T any] struct {
	value T
	next  *ListNode[T]
}

func (listNode ListNode[T]) String() string {
	return fmt.Sprintf("ListNode[value: %v]", listNode.value)
}

func main() {
	strNode := ListNode[string]{value: "Hello", next: nil}
	intNode := ListNode[int]{value: 1, next: nil}
	floatNode := ListNode[float64]{value: 1.0, next: nil}
	fmt.Println(strNode, intNode, floatNode)
}
