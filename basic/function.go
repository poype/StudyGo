package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// 一个方法可以有多个返回值
// 参数 x 和 y 是同一个类型，只要在最后给一个类型声明即可
func calculate(x, y int) (int, int, int, int) {
	return x + y, x - y, x * y, x / y
}

// Named return values
// 相当于在函数最开始定义了 max、min 这两个变量
func takeMaxMin(nums [5]int) (max, min int) {
	max, min = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	// 虽然return语句没有明确指定要返回的变量，但这里会返回 max, min的值
	return
}

// 在 package 级别定义变量，变量的类型都是bool类型
var java, golang, rust bool

// 只有 main package 下的main函数才能被执行
func main1() {
	fmt.Println(add(1, 2))

	a, b, c, d := calculate(8, 4)
	fmt.Println(a, b, c, d)

	nums := [5]int{3, 1, 5, 4, 2}
	maxVal, minVal := takeMaxMin(nums)
	fmt.Println(maxVal, minVal)

	fmt.Println(java, golang, rust)

	// 定义两个变量，自动推断是 int 类型
	var x, y = 1, 2
	fmt.Println(x, y)

	// 在方法内部可以使用 := ， 可以省略 var, 只能在方法内这样用
	o := 100
	fmt.Println(o)
}
