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

func TestAddstring(t *testing.T) {
	// s1 := "99"
	// s2 := "101"
	s1 := "1"
	s2 := "9"
	ss := addstring(s1, s2)
	fmt.Println(ss)

}

func TestSkill(t *testing.T) {
	bby := byte(8 + '0')
	fmt.Println(string(bby))
	return

	// tar := 8
	// i1 := byte(tar) + '0'
	// // i1 := int('8' - '0')
	// fmt.Println(string(i1))
	DoSkill()
}
