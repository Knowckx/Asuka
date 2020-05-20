package main

import (
	"fmt"

	"github.com/Knowckx/Asuka/leetcode/String/title"
)

func main() {
	fmt.Println("Test")
	strs := []string{"flower", "flow", "flight"}
	// strs := []string{"c", "c"}

	out := title.LongestCommonPrefix(strs)
	fmt.Printf("answer:%s\n", out)
}
