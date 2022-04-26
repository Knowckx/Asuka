package crud

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// a lot of bool fields
type BoolConfig struct {
	IsAddName    bool //1
	IsAddAddress bool //2
	IsAddEamil   bool //3
	IsAddAge     bool //4
	IsAddPwd     bool //5
}

// ToBin return binary like "01000"
func (c *BoolConfig) ToBin() string {
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

func (c *BoolConfig) FromBin(ins string) error {
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
			return -1, err
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int(d), nil
}
