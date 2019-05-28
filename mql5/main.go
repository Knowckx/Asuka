package main

func main() {
	ins := makeTestDate() // []*NewObjMod

	// ins2 := []*NewObjMod(ins)
	ins2 := Change2(ins)

	rstMap := ins2.ToMT4AccMap()
	_ = rstMap
	// asuka.Display(rstMap)
	// asuka.Display(ins)

	// var testAcc fmt.Stringer
	// testAcc = MT4Account{BrokerID: 4, Account: "123"}
	// va, ok := testAcc.(fmt.Stringer)
	// fmt.Println(va, ok)

	// testData := rstMap[testAcc]
	// testData.
	// asuka.Display(testData)

	// ss := ins2[0]
	// st, ok := testAcc.(interface{})

	// var i interface{}
	// i = 2
	// b, ok := i.(interface{})
	// fmt.Println(b, ok)

}

// 这个步骤应该在new时就完成了。  对！ out的时候，就拿接口数组来接这个子类值，反正能放进来。
func Change2(ins NewObjMods) MT4AccDatas {
	outs := MT4AccDatas{}
	for _, in := range ins {
		in2 := IMT4AccData(in)
		outs = append(outs, in2)
	}
	return outs
}

// func ToMT4AccMap(li *IMT4AccDatas) MT4AccDataMap {

func ToMT4AccMap(lis []IMT4AccData) MT4AccDataMap {
	rstMap := MT4AccDataMap{}
	// for _, li := range lis {
	// 	li := *li

	// 	IMT4AccData(li).
	// 	mt4acc := li.GetMT4Account()
	// 	rstMap[*mt4acc] = &li
	// }
	return rstMap
}

func makeTestDate() NewObjMods {
	outs := NewObjMods{}
	for idx := 0; idx < 5; idx++ {
		out := NewObjMod{}
		out.ID = idx + 1
		out.MT4Account = MT4Account{BrokerID: int32(idx), Account: "123"}
		outs = append(outs, &out)
	}

	return outs
}
