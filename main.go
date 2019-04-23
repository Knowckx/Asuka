package main

import (
	"fmt"
)

func main() {
	lis := []int{1}
	fmt.Println(len(lis))
	fmt.Println(len(lis[1:2]))

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
