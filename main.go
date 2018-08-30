package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	test()
	// fmt.Println("Hello Asuka")
	// v := []int64{2, 3, 1}
	// Display(v)
}
func test() {

	avg, err := strconv.ParseFloat(fmt.Sprintf("%.3f", 61/float64(12)), 3)
	fmt.Print(avg, err)
}

func Display(x interface{}) {
	name := "Text"
	fmt.Printf("Display %s (%T):\n", name, x)
	// goutil.Display(name, reflect.ValueOf(x))
}

func recordTime() {
	t1 := time.Now()
	//XXX
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

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

// go AsukaDebug(h)

// func AsukaDebug(h *handler.GatewaySrv) {
// 	fmlog.SetConsole()
// 	time.Sleep(time.Duration(2) * time.Second)
// 	h.GetTraderInfo(nil, nil, nil)
// }

//结构体转换
// func BrokerSpreadRecordFromProto(in *pbmo.BrokerSpreadRecord,
// 	) *BrokerSpreadRecord {
// 		out := &BrokerSpreadRecord{
// 			BrokerID:   in.BrokerID,
// 			Symbol:     in.Symbol,
// 			MinSpread:  in.MinSpread,
// 			MaxSpread:  in.MaxSpread,
// 			CreateTime: mdtime.TimeFromProto(in.CreateTime),
// 		}
// 		return out
// 	}

// 	func BrokerSpreadRecordFromProtoL(in []*pbmo.BrokerSpreadRecord,
// 	) []*BrokerSpreadRecord {
// 		outL := []*BrokerSpreadRecord{}
// 		for _, out := range in {
// 			r := BrokerSpreadRecordFromProto(out)
// 			outL = append(outL, r)
// 		}
// 		return outL
// 	}
