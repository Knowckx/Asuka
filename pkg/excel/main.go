package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("tst")
	// StartExcel()
	CreateCsv()
}

func CreateCsv() {
	f, err := os.Create("./test.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM(win支持)

	w := csv.NewWriter(f) //创建一个新的写入文件流

	ins := NewData()

	rrs := ins[10].ErrorLog.ToStrSlice()
	// data := [][]string{
	// 	{"1", "中国", "23"},
	// 	{"2", "美国", "23"},
	// 	{"3", "bb", "23"},
	// 	{"4", "bb", "23"},
	// 	{"5", "bb", "23"},
	// }
	// w.WriteAll(data) //写入数据
	w.Write(rrs)

	w.Flush()
}

func TestNewMod() {
	ins := NewData()
	ins[10].ErrorLog.ToStrSlice()
	// for _,in := range ins {
	// }

}
