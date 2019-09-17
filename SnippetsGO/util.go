package SnippetsGO

import (
	"fmt"
	"reflect"
	"strings"
)

// Printf auto \n
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}

// use '_'
func SliceToString(x interface{}) string {
	ss := ""
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Len() == 0 {
				return ss
			}
			if ss == "" {
				ss = fmt.Sprintf("%s", v.Index(i).Interface())
				continue
			}
			ss = fmt.Sprintf("%s_%s", ss, v.Index(i).Interface())
		}
	}
	return ss
}

func StringToSlice(ss string) []string {
	outs := strings.Split(ss, "_")
	return outs
}

// List去重 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

type IntSlice []int32

// 是否在列表里
func (ints *IntSlice) InSlice(v int32) bool {
	ins := []int32(*ints)
	sign := false
	for _, in := range ins {
		if in == v {
			sign = true
			break
		}
	}
	return sign
}
