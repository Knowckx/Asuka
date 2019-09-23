package funcmodel

import "testing"

func TestIntsToString(t *testing.T) {
	args := []int{0, 1, 2}
	// args := []int{10}
	// args := []int{}
	got := IntsToString(args)
	t.Log(got)
}
