package main

import "fmt"

// <-: channel operator. you can send and receive values with the channel operator, <-.
// ch <- v   	Send v to channel ch.
// v := <-ch  	Receive from ch, and assign value to v.
// sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
// channels must be created before use:  ch := make(chan int)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	// buffered channel: ch := make(chan int, 100)  创建一个100长度的channel
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	ch2 := make(chan int, 1)
	ch2 <- 1
	ch2 <- 2 // deadlock!，上面将channel的长度设置为1，这里会产生deadlock
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
