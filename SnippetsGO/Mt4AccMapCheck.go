package SnippetsGO

import "github.com/Knowckx/Asuka/SnippetsGO/mod"

//双数组使用map进行合并查重
func MapCheck(ins []*mod.AccInfo, accs []*mod.AccScore) []*mod.AccInfo {
	if len(accs) == 0 {
		return ins
	}

	scoreMap := MakeAccMap(accs)

	for _, in := range ins {
		mt4acc := mod.MT4Account{
			BrokerID: in.BrokerID,
			Account:  in.Account,
		}
		s, ok := scoreMap[mt4acc] //检索
		if ok {
			in.Score = s
		}
	}
	return ins
}

//组装索引map
func MakeAccMap(ins []*mod.AccScore) map[mod.MT4Account]int {
	scoreMap := make(map[mod.MT4Account]int) //组装索引map
	for _, in := range ins {
		mt4acc := mod.MT4Account{
			BrokerID: in.BrokerID,
			Account:  in.Account,
		}
		scoreMap[mt4acc] = in.Score
	}
	return scoreMap
}
