package main

import (
	"fmt"

	"github.com/Knowckx/Asuka/leetcode/String/title"
)

func main() {
	fmt.Println("Test")
	// strs := []string{"flower", "flow", "flight"}
	// strs := []string{"c", "c"}

	str := "cbbd"
	res := title.LongestPalindrome(str)
	fmt.Printf("answer:%s\n", res)
}
