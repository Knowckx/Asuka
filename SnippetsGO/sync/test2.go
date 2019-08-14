package main

import (
	"fmt"
	"time"

	"github.com/Knowckx/Asuka/SnippetsGO/onlyone"
)

//--------------------------------------------模型1 使用ch来完成，保证唯一执行的lock功能
func TestOnlyOne2() {
	one := onlyone.GetOnlyOneCh()

	for i := 0; i < 3; i++ {
		for j := 0; j <= 1000; j++ {
			go OnlyOneDo(one)
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Println("main wait over!")
	// <-done
	time.Sleep(7 * time.Second)
	fmt.Println("main over!")

}

func OnlyOneDo(one onlyone.OnlyOneCh) {
	err := one.TryLock()
	if err != nil {
		// fmt.Println(err)
		return
	}
	defer one.UnLock()
	// fmt.Println("In and do!")
	HeavyWork()
	fmt.Println("HeavyWork over!")
	// done <- true
}

func HeavyWork() {
	time.Sleep(1)
}
