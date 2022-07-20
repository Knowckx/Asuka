package gomonkey

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
)

/*
写一下gomonkey的用法
1. go get github.com/agiledragon/gomonkey/v2
2. gomonkey需要加flag： go test -gcflags=all=-l
3. mac会报错 panic: permission denied  解决：https://juejin.cn/post/7069990247005683749
*/

// 使用1: mock函数
func Test_ApplyFunc(t *testing.T) {
	// 我们希望不返回错误
	replace := func(name string) ([]byte, error) {
		return []byte("11"), nil
	}

	// 打桩掉一个标准包的函数
	patches := gomonkey.ApplyFunc(os.ReadFile, replace)
	defer patches.Reset()

	res, err := os.ReadFile("aa/bb/cc")
	fmt.Println(string(res), err) // 输出 11 nil
}

// 方法打桩
func Test_ApplyMethod(t *testing.T) {
	var fakeIn time.Time // 看对方是值方法还是指针方法
	repl2 := func(_ time.Time, layout string) string {
		return "123"
	}
	pa := gomonkey.ApplyMethod(fakeIn, "Format", repl2)
	defer pa.Reset()

	res := time.Now().Format("332211")
	fmt.Println(res)
}
