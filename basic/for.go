package main

import "fmt"

func main() {
	// 没有() ， {}是必须的
	for i := 0; i < 10; i++ {
		// i 变量只在for代码块内可见
		fmt.Println(i)
	}

	fmt.Println("*********************")

	// 类似 while 的循环用法
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("*********************")

	// 无限循环
	for {
		fmt.Println("infinite loop")
	}
}
