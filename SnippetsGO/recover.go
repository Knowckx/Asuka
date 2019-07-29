// 关于Panic与recover的最佳实践
package SnippetsGO

import (
	"fmt"
	"reflect"
)

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

// UT 把panic转换成error
func CheckPanic(err *error) {
	p := recover()
	if p != nil {
		*err = fmt.Errorf("CheckPanic Result [%v]", p)
	}
}
