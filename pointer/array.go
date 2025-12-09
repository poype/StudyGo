package main

import "fmt"

func main() {
	// An array's length is part of its type, so arrays cannot be resized.
	// 不用显示构建一个数组对象就能对其进行赋值
	var words [2]string
	words[0] = "Hello"
	words[1] = "World"
	fmt.Println(words)

	var p *[2]string = &words
	fmt.Println(*p)

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	q := &nums
	// 与C语言不同，go不能直接打印一个变量的地址
	fmt.Println(q) // &[1 2 3 4 5]
	fmt.Printf("%T\n", q)

	// 虽然是数组，但由于没有使用指针类型，这里传递的数组的值，并不能真正修改数组中的值
	// 这点跟C语言不同，C语言中的数组名就是指向数组的隐式指针。go中的数组名不是指针。
	changeArr(nums)
	fmt.Println(nums)

	// 这次传的是数组的指针，可以真正修改数组中的值
	changeArr2(q)
	fmt.Println(nums)

	//testSlice1()

	testSlice3()

	//testSlice4()
	//
	//testSlice5()
	//
	//testSlice6()
	//
	//testSlice7()
}

func changeArr(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++
	}
}

func changeArr2(pArr *[5]int) {
	for i := 0; i < len(pArr); i++ {
		pArr[i]++
	}
}

func testSlice1() {
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// A slice does not store any data, it just describes a section of an underlying array.
	nums := [6]int{1, 2, 3, 4, 5, 6}

	s := nums[1:5]
	fmt.Println(s)
	fmt.Printf("%T\n", s) // []int   slice类型

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	s[1] = 100
	fmt.Println(s)
	fmt.Println(nums)

	// slice literal
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)

	s3 := append(s2, 100)
	fmt.Println(s3)
}

func testSlice2() {
	var a = [5]int{1, 2, 3, 4, 5}

	// below slice expressions are equivalent
	fmt.Println(a[0:5])
	fmt.Println(a[:5])
	fmt.Println(a[0:])
	fmt.Println(a[:])
}

// 多个slice可以共用同一个底层数组，这可以看成存储和计算分离的一种方式
func testSlice3() {
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array.
	var s = []int{1, 2, 3, 4, 5}
	printSlice(s) // len=5 cap=5 [1 2 3 4 5]

	// append 函数不修改原slice, 若底层数组容量不足，会自动扩容，一般会扩容2倍
	// 当 append 后未触发扩容时，新旧 slice 仍共用底层数组；一旦触发扩容，新 slice 会指向新的底层数组，与老 slice 彻底分离。
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
	var s2 = append(s, 6)
	printSlice(s2) // len=6 cap=10 [1 2 3 4 5 6]
	printSlice(s)  // 老的slice还是用老的底层数组，所以没有变

	// s3没有初始化，是nil
	var s3 []int
	printSlice(s3)
	fmt.Println(s3 == nil)

	p := &s3
	fmt.Printf("%T\n", p) // *[]int
	fmt.Println(p)        // &[]
	fmt.Println(*p)       // []
	// slice 类型的变量可以看成是一个结构体，在这个结构体中包含指向底层数组的指针ptr、len、cap
	// 当slice 类型的变量是nil时，表示指向底层数组的指针ptr是"null"，len = 0, cap = 0
	// 这里的指针p指向的是这个结构体，对于数组，指针p是引用的引用。
}

func testSlice4() {
	var s1 = make([]int, 5)
	printSlice(s1)

	var s2 = make([]int, 5, 10)
	printSlice(s2)
}

func testSlice5() {
	// Slices of slices
	var board = [][]string{
		[]string{"O", "X", "O"},
		[]string{"O", "X", "O", "O", "X", "O"},
		[]string{"O", "X", "O"},
	}
	fmt.Println(board)
}

func testSlice6() {
	s := []int{1, 2, 3, 4, 5}
	// We can add more than one element at a time.
	s = append(s, 6, 7, 8, 9, 10)
	printSlice(s)
}

func testSlice7() {
	fmt.Println("---------------testSlice7 range----------------")

	s := []string{"one", "two", "three", "four", "five", "six", "seven"}
	// i是index，v是对一个的值
	// The first is the index, and the second is a copy of the element at that index.
	for i, v := range s {
		fmt.Println(i, v)
	}
	// range不是函数，属于表达式

	// skip the index or value by assigning to _
	for i, _ := range s {
		fmt.Println(i, s[i])
	}

	for _, v := range s {
		fmt.Println(v)
	}
	// If you only want the index, you can omit the second variable.
	for i := range s {
		fmt.Println(i)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
