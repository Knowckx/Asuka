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
