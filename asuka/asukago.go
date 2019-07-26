package asuka

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Display(x interface{}) {
	name := "Text"
	fmt.Printf("Asuka Display..Target Type:[%T]:\n", name, x)
	displayPath(name, reflect.ValueOf(x))
}

func displayPath(path string, v reflect.Value) {
	va := v.Interface() //
	v.Type()
	switch va.(type) {
	case time.Time:
		fmt.Printf("%s = %s\n", path, va.(time.Time).String())
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		leng := v.Len()
		if leng > 5 {
			fmt.Printf("%s Slice.len = %d\n", path, leng)
			leng = 5
		}
		for i := 0; i < leng; i++ {
			displayPath(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		// fmt.Printf("%s %#v\n", path, v) //正面再嵌结构体的话…… %#v 对指针值无力
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			displayPath(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fieldPath := fmt.Sprintf("%s[%+v]", path, key)
			displayPath(fieldPath, v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fieldPath := fmt.Sprintf("*%s", path)
			displayPath(fieldPath, v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			displayPath(path+".value", v.Elem())
		}

	default: //简单类型喽
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// 基本类型，不会嵌套，直接打印
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + "value ????"
	}
}
