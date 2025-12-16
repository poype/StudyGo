package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	// This statement asserts that the interface value i holds the concrete type string
	// and assigns the underlying string value to the variable s.
	var s string = i.(string)
	fmt.Println(s) // Hello

	s, ok := i.(string)
	fmt.Println(s, ok) // Hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	//	n := i.(float64)
	//	fmt.Println(n)
	//	panic: interface conversion: interface {} is string, not float64
}
