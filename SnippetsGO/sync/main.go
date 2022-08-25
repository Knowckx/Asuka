package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// var a1 AA
	// pa := &a1
	// pa.D1()
	// pa.D2()

	TestOnlyOne2()
	// Test()
}

type AA int

// 指针方法
func (in *AA) D1() {
	fmt.Println("D1")
}

// 值方法
func (in AA) D2() {
	fmt.Println("D2")
}

type BB int

func Test() {
	var a AA = 1
	b := interface{}(a)
	c, ok := b.(BB)
	fmt.Println(c, ok)
}

func SyncOnceTest() {
	var once sync.Once
	InitFunc := func() {
		fmt.Println("Only once")
	}
	once.Do(InitFunc) // 可以保证此方法只被触发一次
	once.Do(InitFunc) // 已触发过之后，再调用就会直接跳过
}

func SyncWaitGroup() {
	urls := []string{}

	var wg sync.WaitGroup
	for _, url := range urls { //启动多个gro
		wg.Add(1) //每启动一个，计数器+1

		go func(url string) {
			defer wg.Done() //每完成一个，计数器-1
			http.Get(url)
		}(url)
	}
	wg.Wait() //等待全部完成
}
