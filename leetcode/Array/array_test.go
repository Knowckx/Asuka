package Array

import (
	"fmt"
	"testing"
)

func TestMerge2Slice(t *testing.T) {
	s1 := []int{1, 2, 3, 0, 0, 0}
	s2 := []int{2, 5, 6}
	fmt.Println("We Test")
	merge2Slice(s1, 3, s2, 3)
}
