package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	big   = 100
	small = 1
)

func main() {
	// %T 类型， %v对应 值
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 默认值 0
	var x int
	// 默认值 空字符串
	var y string
	fmt.Println("x:", x, "y:", y)

	var num1, num2 int = 3, 2
	fmt.Println(num1 / num2) // 1

	// 类型转换
	fmt.Println(float32(num1) / float32(num2)) // 1.5

	var o, p = 100, 3.14
	fmt.Printf("%T \n", o) // 证书的默认类型是 int
	fmt.Printf("%T \n", p) // 小数的默认类型是 float64

	var ch = 'a'
	fmt.Printf("%T \n", ch) // 字符类型是rune，rune是int32的别名,

	// const
	const pi = 3.1415926
	fmt.Printf("%T \n", pi)

	fmt.Println(big, small)
}
