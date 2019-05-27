package utlity

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//WaitForServer 重试一个
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute                        //重试时间
	deadline := time.Now().Add(timeout)                    //定义退出时间
	for tries := 0; time.Now().Before(deadline); tries++ { //只要没到退出时间 一直循环
		_, err := http.Head(url) //可能出错
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) //这里的休眠时间是一个位运算
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout) //最后也没成功。

}

//功能，记录执行时间1
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // --事实上这一条语句可以控制函数所有的入口和出口  牛！
	// ...睡5秒，模拟大操作。
	time.Sleep(5 * time.Second)
}

//2
func trace(msg string) func() { //返回一个无参无返回值的内函数
	start := time.Now()
	log.Printf("enter %s", msg) //执行到了，记录时间
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start)) //最后退出了，记录时间
	}
}
