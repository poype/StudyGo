package test_export

import "fmt"

func DogRun() {
	fmt.Println("Dog run")
	fmt.Println(pi)
}

func catRun() {
	fmt.Println("Cat run")
}

const PI = 3.1415926

const pi = 12345

// 必须以大写字母开头的变量和方法才能被export，小写字母开头的名字只能在package内使用
