package main

import "github.com/Knowckx/Asuka/leetcode/String/title"

func main() {
	// title.Stack
	nds1 := title.GenListNode(5)
	nds1.Print()
	nds2 := title.GenListNode(7)
	nds2.Print()
	out := title.AddTwoNumbers(nds1, nds2)
	out.Print()
}

// E:\Develop\Go\Asuka\leetcode\linkedlist\main.go
// E:\Develop\Go
