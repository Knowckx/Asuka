package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas map[int64]string
var i = 0

func main() {

	datas = make(map[int64]string)
	go http.ListenAndServe("0.0.0.0:6060", nil)
	loop()
}

func loop() {
	for {
		if i >= 20000 {
			time.Sleep(time.Second * 2)
		}
		i++
		ss := fmt.Sprintf("%s%d", "https://github.com/EDDYCJY", i)
		Add(ss)
	}

}

func Add(str string) map[int64]string {
	datas[int64(i)] = str
	fmt.Println(i, datas[int64(i)])

	return datas
}
