package main

import (
	"bufio"
	"fmt"
	"os"
)

// 顺利打印 1-100 使用chan来通信

var chSin = make(chan bool, 1)
var chDou = make(chan bool, 1)

func PrintSingle() {

	for i := 0; i < 100; i += 2 {
		<-chSin
		fmt.Println(i)
		chDou <- true
	}
}

func PrintDouble() {
	for i := 1; i < 100; i += 2 {
		<-chDou
		fmt.Println(i)
		chSin <- true
	}
}

func main() {
	chSin <- true
	go PrintSingle()
	go PrintDouble()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}

}
