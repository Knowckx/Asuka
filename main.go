package main

import (
	"fmt"
)

func main() {
	Atest()
}

type AA struct {
	v1 int
	v2 string
	v3 float32
}

func test() {
	// asuka.Test()

}

func Display(x interface{}) {
	name := "Text"
	fmt.Printf("Display %s (%T):\n", name, x)
	// goutil.Display(name, reflect.ValueOf(x))
}

func Atest() {
	slice_1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("main-->data:\t%#v\n", slice_1)
	fmt.Printf("main-->len:\t%#v\n", len(slice_1))
	fmt.Printf("main-->cap:\t%#v\n", cap(slice_1))
	test1(slice_1)
	fmt.Printf("main-->data:\t%#v\n", slice_1)

	test2(&slice_1)
	fmt.Printf("main-->data:\t%#v\n", slice_1)
}

func test1(slice_2 []int) {
	slice_2[1] = 6666               // 函数外的slice确实有被修改
	slice_2 = append(slice_2, 8888) // 函数外的不变
	fmt.Printf("test1-->data:\t%#v\n", slice_2)
	fmt.Printf("test1-->len:\t%#v\n", len(slice_2))
	fmt.Printf("test1-->cap:\t%#v\n", cap(slice_2))
}

func test2(slice_2 *[]int) { // 这样才能修改函数外的slice
	*slice_2 = append(*slice_2, 6666)
}
