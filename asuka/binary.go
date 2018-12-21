package asuka

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// 结构体配置转二进制数
type Cfg struct {
	Score bool //1
	ROI   bool //2
	MR    bool //3
}

func Test() {
	// var int2 = 5

	fmt.Println("Start")
	cf := new(Cfg)
	cf.MR = false
	cf.ROI = true
	// fmt.Println(cf.ToBin())
	v := reflect.ValueOf(cf).Elem()
	fmt.Println(v.Kind())
	_ = v.NumField() - 5
	return

	// // 位数不足，左补0
	// func (c *RankFiledHideConfig) FromDec(dec int) error {
	// 	ss := DecBin(dec)
	// 	leng := len(ss)
	// 	v := reflect.ValueOf(c).Elem()

	err := cf.FromBin("110")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", cf)

}

func (c *Cfg) ToBin() string {
	v := reflect.ValueOf(*c)
	rst := ""
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i).Interface().(bool)
			if val == false {
				rst = rst + "0"
				continue
			}
			rst = rst + "1"
		}
	}
	return rst
}

func (c *Cfg) FromBin(ins string) error {
	leng := len(ins)

	v := reflect.ValueOf(c).Elem()
	switch v.Kind() {
	case reflect.Struct:
		if leng != v.NumField() {
			return fmt.Errorf("length of binary and struct should be equal")
		}
		for i := 0; i < leng; i++ {
			intv, err := strconv.Atoi(ins[i : i+1])
			if err != nil {
				return err
			}
			boV := false
			if intv != 0 {
				boV = true
			}
			v.Field(i).SetBool(boV)

		}
	}
	return nil
}

// Decimal to binary  10 -> 1010
func DecBin(n int) string {
	return fmt.Sprintf("%b", n)
}

// binary to Decimal  1010 -> 10
func BinDec(b string) (int, error) {
	ss := strings.Split(b, "")
	l := len(ss)
	d := float64(0)
	for i := 0; i < l; i++ {
		f, err := strconv.ParseFloat(ss[i], 10)
		if err != nil {
			return -1, fmt.Errorf("Binary to decimal error:", err.Error())
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int(d), nil
}
