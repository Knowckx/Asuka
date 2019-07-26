package main

import (
	"fmt"
)

func main() {
	var a interface{} = 215
	// fmt.Printf("%T",a)
	fmt.Print(a)
	fmt.Println(a)
	rst := test()
	fmt.Printf("%T", rst)

	return
}

//  golang 缺乏总体断言   从 []int   []interface{}
// 问题就变成了，如何从一个给定的类型，自动创建该类型的slice
func test(in interface{}) interface{} {
	outs := make([]interface{}, 0)
	var a1 interface{} = 10
	var a3 interface{} = 13
	var a4 interface{} = 14

	outs = append(outs, a1)
	outs = append(outs, a3)
	outs = append(outs, a4)
	return outs

}

type AA struct {
	v1 int
	v2 string
	li []int
}

func (A AA) Copy() AA {
	return A
}

// 数组指针的改变问题
func testPoint() {
	lis1 := []int{1, 2, 3, 4, 5}
	lis2 := lis1
	pb1 := &lis1
	pb2 := &lis2
	fmt.Printf("%p %p %p %p\n", lis1, lis2, pb1, pb2)
	lis1 = []int{1}
	fmt.Printf("%v %v", lis1, lis2)
	fmt.Printf("%v %v", lis1, lis2)
}

func SliceTest(in *AA) {
	list := in.li
	list = list[0:1]
	in.li = list
	// fmt.Printf("%p\n", ins)
	// ins[2] = 10
	// ins = nil
	// fmt.Printf("%p\n", ins)
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
