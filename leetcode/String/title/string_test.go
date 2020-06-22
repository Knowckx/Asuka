package title

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {

	s1 := "2"
	s2 := "3"
	ss := Multiply3(s1, s2)
	fmt.Println(ss)
}

func TestForCase1(t *testing.T) {
	bys := []byte{}
	bys = append(bys, '1')
	bys = append(bys, 2)
	bys = append(bys, 3)
	fmt.Println(string(bys))
}
