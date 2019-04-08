package main

import (
	"fmt"
	"reflect"

	"github.com/Knowckx/Asuka/SnippetsGO/mod"
)

type Test struct {
	ID int
}

func (t *Test) PrintField() {
	v := reflect.ValueOf(t).Elem()
	fmt.Println(v.NumField())
}

func main() {
	t1 := mod.Test1{}

	t1.Account = "t1"
	t1.BrokerID = 11
	t1.Test2.Account = "t2"
	t1.Test2.BrokerID = 22

	t2 := t1
	t1.Test2.Account = "t3"

	pt1 := &(t1.Test2)
	pt2 := &(t2.Test2)
	fmt.Print(pt1, pt2)
	return

	brk := mod.MT4Account{}
	brk.Account = "33123"
	brk.BrokerID = 21
	// 泛型
	fmt.Println(brk)
	fmt.Printf("%#v %+v", brk, brk)
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
