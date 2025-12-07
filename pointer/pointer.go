package main

import "fmt"

func main1() {
	var i = 100
	// *int 类型，跟C语言比是反着的，这里的 *int 无需明确指定
	var p *int = &i
	*p++
	fmt.Println(*p)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

// Unlike C, Go has no pointer arithmetic.
