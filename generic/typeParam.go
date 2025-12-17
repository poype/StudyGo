package main

import "fmt"

type Runner interface {
	Run()
}

type Dog struct{}

func (dog *Dog) Run() {
	fmt.Println("dog is running")
}

type Cat struct{}

func (cat *Cat) Run() {
	fmt.Println("cat is running")
}

// 该方法可以接受任何实现了Runner接口的对象作为参数
func sport[T Runner](pet T) {
	pet.Run()
}

func main() {
	cat := &Cat{}
	dog := &Dog{}
	// dog2 := Dog{}

	sport(cat)
	sport(dog)
	// sport(dog2) // dog2对象不是指针类型，没有实现Run方法，就不能作为sport方法的参数
}
