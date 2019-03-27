package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	ID int
}

func (t *Test) PrintField() {
	v := reflect.ValueOf(t).Elem()
	fmt.Println(v.NumField())
}

func main() {
	// 泛型
	fmt.Printf("%d", 100)
	fmt.Printf("%d", int32(100))
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
