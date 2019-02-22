package main

import (
	"fmt"
	"reflect"

	"github.com/Knowckx/Asuka/asuka"
)

type Test struct {
	ID int
}

func (t *Test) PrintField() {
	v := reflect.ValueOf(t).Elem()
	fmt.Println(v.NumField())
}

func main() {
	// ss := "userId_account_brokerId"
	// rst := strings.Split(ss, "_")
	// asuka.Display(rst)
	ints := []int32{1, 5, 45, 4, 2}

	ss := asuka.SliceToString(ints)
	fmt.Println(ss)
	a := asuka.StringToSlice(ss)
	fmt.Println(a)

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
