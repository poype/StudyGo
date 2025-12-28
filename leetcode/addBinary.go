package main

import (
	"strings"
)

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	stack := Stack{data: make([]rune, 0)}

	carry := '0'
	for i >= 0 && j >= 0 {
		if a[i] == '1' && b[j] == '1' {
			stack.Push(carry)
			carry = '1'
		} else if a[i] == '0' && b[j] == '0' {
			stack.Push(carry)
			carry = '0'
		} else {
			if carry == '1' {
				stack.Push('0')
				carry = '1'
			} else {
				stack.Push('1')
				carry = '0'
			}
		}
		i--
		j--
	}

	if i < 0 {
		i = j
		a = b
	}

	for i >= 0 {
		if a[i] == '1' && carry == '1' {
			stack.Push('0')
			carry = '1'
		} else if a[i] == '0' && carry == '0' {
			stack.Push('0')
			carry = '0'
		} else {
			stack.Push('1')
			carry = '0'
		}
		i--
	}

	if carry == '1' {
		stack.Push('1')
	}

	var builder strings.Builder
	for !stack.isEmpty() {
		builder.WriteRune(stack.Pop())
	}
	return builder.String()
}
