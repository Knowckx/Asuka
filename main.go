package main

import (
	"fmt"
	"reflect"

	"github.com/Knowckx/Asuka/asuka"
	"github.com/Knowckx/Asuka/pkg/excel"
)

type Test struct {
	ID int
}

func (t *Test) PrintField() {
	v := reflect.ValueOf(t).Elem()
	fmt.Println(v.NumField())
}

func main() {
	excel.StartExcel()
}

type AA struct {
	v1 int
	v2 string
	v3 float32
}

func test() {
	asuka.Test()

}

func Display(x interface{}) {
	name := "Text"
	fmt.Printf("Display %s (%T):\n", name, x)
	// goutil.Display(name, reflect.ValueOf(x))
}

// List去重 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
