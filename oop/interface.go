package main

import "fmt"

func main2() {
	// Interface values with nil underlying values
	var p Abser
	fmt.Printf("%T\n", p) // <nil>

	testNum := TestNum{-100}
	p = &testNum
	fmt.Println(p.Abs())
	fmt.Printf("%T\n", p)

	// 由于方法的receiver是 this *TestNum，所以不能直接将testNum变量赋值给p，只能将testNum变量的地址赋值给p
	// p = testNum  // Cannot use testNum (type TestNum) as the type Abse
}

// Abser An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type Abser interface {
	Abs() int
}

type TestNum struct {
	num int
}

func (this *TestNum) Abs() int {
	if this.num < 0 {
		return -this.num
	}
	return this.num
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// in Go it is common to write methods that gracefully handle being called with a nil receiver
	// 如果是java那只能抛NPE
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// interface{} 如果一个interface中没有任何方法，那就是一个empty interface.
// An empty interface may hold values of any type.
// Empty interfaces are used by code that handles values of unknown type.
