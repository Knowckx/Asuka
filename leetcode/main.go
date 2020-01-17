package main

import "fmt"

import "time"

func main() {
	TestChan()
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
