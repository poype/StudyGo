package main

import "fmt"

func main2() {
	p := Point{1, 2}
	fmt.Println(p)
	// 注意这里p不是指针
	fmt.Printf("%T\n", p)
	fmt.Println(p.x)
	fmt.Println(p.y)

	// q 是指针类型 *main.Point
	q := &p
	fmt.Printf("%T\n", q)
	fmt.Println(q)
	fmt.Println(*q)

	// 传统用法
	fmt.Println((*q).x)
	// 虽然q是指针类型，但这里也不用加* 就可以引用到结构体的成员变量
	// To access the field x of a struct when we have the struct pointer q we could write (*q).x.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	fmt.Println(q.x)

	// 1. 直返返回的就是指针类型， 2. 指定构造参数名字
	p2 := &Point{y: 200, x: 100}
	fmt.Printf("%T\n", p2)
	fmt.Println(*p2)

}

type Point struct {
	x int
	y int
}
