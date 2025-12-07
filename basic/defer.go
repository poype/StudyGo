package main

import "fmt"

func main() {
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	str := "one"
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println(str)

	str = "two"
	fmt.Println(str)
}

// 先打印 two，再打印one
