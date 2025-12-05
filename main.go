package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
 1. 每个文件必须要以 package 开头
 2. math/rand 是import path， rand 是 package
*/

func main() {
	fmt.Println("Hello 世界！")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Abs(-100))
}

// 每条语句后面无需 ;
