package gosrc

import (
	"context"
	"fmt"
	"time"
)

func ContextStart() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go Executor(ctx)

	// cancel after 1s
	time.Sleep(time.Second)
	cancel() // 主函数主动取消
}

// 实际的执行者
func Executor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Doing")
		}
	}
}

// 模式2：通过chan来通知停止工作
func TestChan() {
	stopCh := make(chan int) // 结束信号
	go Worker(stopCh)

	time.Sleep(5 * time.Second)
	fmt.Println("Stop")
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
