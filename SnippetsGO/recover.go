// 关于Panic与recover的最佳实践
package SnippetsGO

import (
	"fmt"
	"reflect"
)

func StartRecover() {
	err := DoPanic()
	fmt.Println(err)
}

// 函数内使用了反射，会产生Panic
func DoPanic() (err error) {

	defer CheckPanic(&err)
	var a interface{} = 10
	va := reflect.ValueOf(a)
	// _ = va
	va.SetBool(true) // will panic
	fmt.Println("DoPanic End")
	return nil
}

func GenTestArgs() (int, error) {
	return 1, fmt.Errorf("123")
}
