package main

import (
	"fmt"

	"github.com/Knowckx/Asuka/asuka"
)

func main() {
	test()
	// fmt.Println("Hello Asuka")
	// v := []int64{2, 3, 1}
	// Display(v)
}

type AA struct {
	v1 int
	v2 string
	v3 float32
}

func test() {
	aa := make([]*AA, 0)

	a := &AA{
		v1: 1,
		v2: "st",
		v3: 1.21,
	}
	a1 := &AA{
		v1: 2,
		v2: "st22",
		v3: 1.21,
	}

	aa = append(aa, a)
	aa = append(aa, a1)

	asuka.Display(aa)
	// fmt.Printf("%v\n", a)

	// fmt.Printf("%+v\n", a)

	// fmt.Printf("%#v\n", a)

	// avg, err := strconv.ParseFloat(fmt.Sprintf("%.3f", 61/float64(12)), 3)
	// fmt.Print(avg, err)
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
