package main

import "fmt"

// 方法也可以被当做参数和返回值
// Functions are values too. They can be passed around just like other values.

func main() {
	testFn := func(a, b float64) float64 {
		return a + b
	}
	fmt.Println(compute(testFn)) // 7

	returnFn()() // Return Fn

	adder1 := createAdder()
	adder2 := createAdder()

	for i := 0; i < 10; i++ {
		fmt.Println(adder1(1))
		fmt.Println(adder2(2))
		fmt.Println("---------------------")
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func returnFn() func() {
	return func() {
		fmt.Println("Return Fn")
	}
}

// closure
func createAdder() func(num int) int {
	// 每个adder都有自己独有的sum变量
	sum := 0

	return func(num int) int {
		sum += num
		return sum
	}
}
