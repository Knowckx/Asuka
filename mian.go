package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Knowckx/Asuka/goutil"
)

func main() {

	fmt.Println("Hello Asuka")
	v := []int64{2, 3, 1}
	Display(v)
}

func Display(x interface{}) {
	name := "Text"
	fmt.Printf("Display %s (%T):\n", name, x)
	goutil.Display(name, reflect.ValueOf(x))
}

func PrintS() {
	fmt.Printf("!")
	// out []int{}
	// for ii, vv := range out {
	// 	fmt.Printf("List num %d \n%+v\n", ii+1, vv)
	// }
}

func recordTime() {
	t1 := time.Now()
	//XXX
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

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
