package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Knowckx/Asuka/SnippetsGO"
)

// Go里面，函数中实现泛型切片的最佳实践是？

type MyInt int

func (MyInt) test() {}

func main() {
	SnippetsGO.StartRecover()
	return
}

func test(in interface{}) {
	// slice := make([]int, 0)
	slice2 := []int{}
	t0 := time.Now()
	// for i := 1; i < 10; i++ {
	// 	slice = append([]int{i}, slice[0:]...)
	// 	// slice = append(slice, i)

	// }
	// fmt.Println(slice)

	t1 := time.Now()
	for i := 1; i < 10000; i++ {
		slice2 = Insert(slice2, 0, i).([]int)
	}
	t2 := time.Now()
	fmt.Println(t1.Sub(t0), t2.Sub(t1))
	fmt.Println(t1.Sub(t0), t2.Sub(t1))

}

func Insert(slice interface{}, pos int, value interface{}) interface{} {

	v := reflect.ValueOf(slice)

	fmt.Println(v.String())

	var aa interface{} = 1
	var aa1 interface{} = []int{}

	tyaa := reflect.TypeOf(aa)
	tyaa1 := reflect.TypeOf(aa1)

	fmt.Println(tyaa.String())
	fmt.Println(tyaa1.String())

	// ne := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 1, 1)
	ne := reflect.MakeSlice(tyaa1, 1, 1)

	fmt.Println(ne.String())

	ne.Index(0).Set(reflect.ValueOf(value))
	v = reflect.AppendSlice(v.Slice(0, pos), reflect.AppendSlice(ne, v.Slice(pos, v.Len())))

	return v.Interface()
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
