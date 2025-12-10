package main

import "fmt"

func main() {
	//testMap1()

	testMap2()

	//testMap3()
}

func testMap1() {
	var m map[string]int
	fmt.Println(m) // map[]
	fmt.Println(m == nil)

	// map是nil，不能向nil map中加值
	m["one"] = 100 // panic: assignment to entry in nil map
	fmt.Println(m)
}

func testMap2() {
	var m = make(map[string]int)
	fmt.Println(m)        // map[]
	fmt.Println(m == nil) // false, 虽然map是空的，但这个map已经不是nil

	m["one"] = 100
	fmt.Println(m)        // map[one:100]
	fmt.Println(m["one"]) // 100

	m["two"] = 200
	fmt.Println(m)

	// delete
	delete(m, "one")
	fmt.Println(m)

	elem, ok := m["two"]
	fmt.Println(elem, ok) // 200 true

	elem, ok = m["one"]
	// If key is not in the map, then elem is the zero value for the map's element type, the ok is false.
	fmt.Println(elem, ok) // 0 false
}

type Person struct {
	firstName string
	lastName  string
}

func testMap3() {
	// map literal
	var m = map[string]Person{
		"one":   {firstName: "San", lastName: "Zhang"},
		"two":   {firstName: "Si", lastName: "Li"},
		"three": {firstName: "Wu", lastName: "Wang"},
	}
	fmt.Println(m)
	fmt.Println(m["one"])
}
