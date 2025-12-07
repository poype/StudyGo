package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main5() {
	num := rand.Int() % 10
	fmt.Println(num)

	// 每个case不需要加break
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	switch num {
	case 9:
		fmt.Println("The number is greater than 5")
	case 8:
		fmt.Println("The number is greater than 5")
	case 7:
		fmt.Println("The number is greater than 5")
	case 6:
		fmt.Println("The number is greater than 5")
	case 5:
		fmt.Println("The number is equal to 5")
	default:
		fmt.Println("The number is less than 5")
	}

	// Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
