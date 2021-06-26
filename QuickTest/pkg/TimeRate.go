package pkg

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/util/log"
	"golang.org/x/time/rate"
)

// time/rate 这个包的令桶牌更像滑动窗口
var GlbLit *rate.Limiter

// 模拟外部程序调用接口
func StartTest() {
	GlbLit = rate.NewLimiter(0.334, 1)
	for {
		Consumer()
		time.Sleep(1 * time.Second)
	}
}
func Consumer() {
	ok := GlbLit.Allow()
	if !ok {
		log.Info("No Token.returned.")
		return
	}
	fmt.Println("Consumer Start")
	time.Sleep(1 * time.Second)
	fmt.Println("Consumer End")
}
