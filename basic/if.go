package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Int()

	fmt.Println(num)
	// if的表达式不需要()，{}是必须的
	if num > 5 {
		fmt.Println("The number is greater than 5")
	} else if num > 0 {
		fmt.Println("The number is greater than 0")
	} else {
		fmt.Println("The number is equal or less than 0")
	}

	// 在if表达式中包含条件变量的初始化， 下面的num变量只在 if - else 代码块中可见
	if num := rand.Int(); num > 4232820145393171355 {
		fmt.Println("The number is greater than 4232820145393171355")
	} else if num > 0 {
		fmt.Println("The number is greater than 0")
	} else {
		fmt.Println("The number is equal or less than 0")
	}
}
