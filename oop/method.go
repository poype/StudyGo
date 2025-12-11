package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

// Abs 定义Vertex类的方法，相比于普通函数，在func关键字和方法名之间增加了类Receiver参数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs a method is just a function with a receiver argument. 下面的方法与上面的方法等价
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale Pointer receivers， Methods with pointer receivers can modify the value to which the receiver points
// 这里必须用指针，因为这里直接修改了对象的值
// pointer receivers are more common than value receivers. 应该一直使用 Pointer receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale 这个函数与上面的方法等价
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// MyFloat non-struct types.
type MyFloat float64

// Abs You can declare a method on non-struct types.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
// 所以我们不能直接给string类型创建一个方法，因为string类型不在这个包内
//
//	func (f string) Abs() float64 {
//		 return f
//	}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	var f MyFloat = -4.0
	fmt.Println(f.Abs())

	// methods with pointer receivers take either a value or a pointer as the receiver when they are called
	v.Scale(10)
	fmt.Println(v.Abs())

	var p *Vertex = &Vertex{3, 4}
	p.Scale(10)
	fmt.Println(p.Abs())
}

/*
	There are two reasons to use a pointer receiver.
	The first is that the method can modify the value that its receiver points to.
	The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct.
*/
