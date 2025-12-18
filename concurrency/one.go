package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

// A goroutine is a lightweight thread managed by the Go runtime.

func main() {
	// execution in the new goroutine, but the evaluation in the current goroutine
	go say("world")
	say("hello")
}
