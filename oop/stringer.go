package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main5() {
	p := &Person{Name: "Jacky", Age: 20}
	fmt.Println(p) // Jacky is 20 years old
}
