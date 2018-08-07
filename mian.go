package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Hello Asuka")
}

func Prt(out []*interface{}) {
	for ii, vv := range out {
		fmt.Printf("List num %d \n%+v\n", ii+1, vv)
	}

}

func recordTime() {
	t1 := time.Now()
	//XXX
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

}
