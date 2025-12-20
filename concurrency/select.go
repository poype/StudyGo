package main

import (
	"fmt"
	"time"
)

func producer1(ch chan int) {
	i := 0
	for {
		ch <- i
		i += 2
		time.Sleep(1 * time.Second)

		if i > 100 {
			ch <- -1
		}
	}
}

func producer2(ch chan int) {
	i := 1
	for {
		ch <- i
		i += 2
		time.Sleep(2 * time.Second)

		if i > 100 {
			ch <- -1
		}
	}
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go producer1(ch1)
	go producer2(ch2)

	ch1Down, ch2Down := false, false

	for {
		// The select lets a goroutine wait on multiple channels.
		// A select blocks until one of its cases can run, then it executes that case.
		// It chooses one at random if multiple are ready.
		select {
		case x := <-ch1:
			if x == -1 {
				ch1Down = true
			}
			fmt.Println("ch1: ", x)
		case y := <-ch2:
			if y == -1 {
				ch2Down = true
			}
			fmt.Println("ch2: ", y)
		}
		if ch1Down && ch2Down {
			return
		}
	}
}
