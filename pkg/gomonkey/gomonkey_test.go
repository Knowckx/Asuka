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
2. gomonkey需要加flag： go test -gcflags=all=-l   | vscode 搜索配置：go.testFlags
3. mac会报错 panic: permission denied  解决：https://juejin.cn/post/7069990247005683749

作者说的用法 https://blog.csdn.net/kevin_tech/article/details/126204044
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
	appMe := func(_ time.Time, layout string) string {
		return "123"
	}
	pa := gomonkey.ApplyMethod(fakeIn, "Format", appMe)
	defer pa.Reset()

	res := time.Now().Format("332211")
	fmt.Println(res)
}

// 接口打桩
func TestApplyInterfaceReused(t *testing.T) {
	e := &Tom{}

	patches := gomonkey.ApplyFunc(NewPeopel, func(name string) Peopel {
		return e // 让返回的接口值指向我们的对象
	})
	defer patches.Reset()

	replace := func(_ *Tom) (string, error) { // 变成了打桩一个方法
		return "hello Tom 123", nil
	}
	patches.ApplyMethod(e, "GetName", replace)

	pe := NewPeopel("Tom2") // 调一下
	output, err := pe.GetName()
	fmt.Println(output, err)
}

// 桩序列
// https://blog.csdn.net/kevin_tech/article/details/126204044
