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
