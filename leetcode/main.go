package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("123")
	// TestChan()

	out := testCache("in1")
	fmt.Println(string(out))
	out = testCache("in1")
	fmt.Println(string(out))

	out = testCache("in2")
	fmt.Println(string(out))

}

func TestChan() {
	stopCh := make(chan int) // 结束信号
	go Worker(stopCh)

	time.Sleep(5 * time.Second)
	fmt.Println("ask To Stop")
	stopCh <- 1
}

func Worker(stop chan int) {
	for {
		select {
		case <-stop:
			fmt.Println("Worker Stop!")
			return
		default:
			fmt.Println("Working")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteChan(ch chan<- int) {
	ch <- 1
	ch <- 2
	close(ch)
	fmt.Println("Do Close")
}

var cachemu sync.RWMutex
var cache = map[string][]byte{}

func testCache(in string) []byte {
	cachemu.RLock()
	va, ok := cache[in]
	cachemu.RUnlock()
	if ok {
		return va
	}
	// refresh cache
	newVa := []byte(in + "123")

	cachemu.Lock()
	cache[in] = newVa
	cachemu.Unlock()

	return newVa
}
