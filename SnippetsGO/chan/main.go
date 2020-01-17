package main

import (
	"fmt"
	"time"
)

// 模式：通过chan来通知停止工作
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
