package main

import (
	"fmt"
	"sync"
	"time"
)

// make sure only one goroutine can access a variable at a time to avoid conflicts
// mutual /ˈmjuːtʃuəl/ 相互的，彼此的
// Go's standard library provides mutual exclusion with sync.Mutex and its two methods: Lock 、 Unlock

type SafeCounter struct {
	val int
	mux sync.Mutex
}

func (this *SafeCounter) Inc() {
	this.mux.Lock()
	defer this.mux.Unlock() // 利用defer确保锁一定会释放，避免遗漏

	this.val++
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), this.val)
	time.Sleep(1 * time.Second)
}

func main() {
	counter := SafeCounter{}
	for i := 0; i < 100; i++ {
		// 由于有mutex加锁的保护，同一个时刻只有一个goroutine可以执行Inc内的逻辑
		go counter.Inc()
	}
	time.Sleep(200 * time.Second)
}
