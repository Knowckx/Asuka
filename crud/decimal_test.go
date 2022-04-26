package crud

import (
	"fmt"
	"testing"
)

func Test_ToBin(t *testing.T) {
	cf := new(BoolConfig)
	cf.IsAddAddress = true

	fmt.Printf("Config: %+v\n", cf)

	resBin := cf.ToBin()
	fmt.Printf("Bin: %+v\n", resBin) // 01000

	dec, err := BinDec(resBin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dec: %+v\n", dec) // 8
	return
}
