package main

import "fmt"

func main4() {
	do(21)
	do("Test String")
	do(3.1415926)
	do('3') // I don't know about type switch
}

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("string: ", v)
	case int:
		fmt.Println("int: ", v)
	case float64:
		fmt.Println("float64: ", v)
	default:
		fmt.Println("I don't know about type switch")
	}
}
