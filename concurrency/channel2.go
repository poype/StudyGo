package main

import "fmt"

func genInt(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// A sender can close a channel to indicate that no more values will be sent.
	close(ch)
}

func main() {
	ch := make(chan int, 3)
	go genInt(ch)

	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for num := range ch {
		fmt.Println(num)
	}

	// Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression
	_, ok := <-ch
	fmt.Println(ok) // false  ok is false if there are no more values to receive and the channel is closed.
}

// 1. Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// 2. you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more
//    values coming, such as to terminate a range loop.
