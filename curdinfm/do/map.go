package do

import "github.com/Knowckx/Asuka/curdinfm/mod"

// 常用 把一组MT4账户 组装成索引map
func MakeAccMap(ins []*mod.MT4Account) *mod.MT4AccData {
	rstMap := mod.MT4AccData{} //组装索引map
	for _, in := range ins {
		objMod := new(mod.NewObjMod)
		objMod.MT4Account = in
		rstMap[*in] = objMod
	}
	return &rstMap
}

// 已经有一个数组，带mt4账户的  组装一下
func MakeAccMap2(ins []*mod.NewObjMod) *mod.MT4AccData {
	rstMap := mod.MT4AccData{} //组装索引map
	for _, in := range ins {
		mt4acc := in.MT4Account
		rstMap[*mt4acc] = in
	}
	return &rstMap
}

// type interface MT4AccountData {}

// keys := make([]int, len(mmap))
// for k := range mmap {
//     keys = append(keys, k)
// }

// vs := make([]string, len(mmap))
// for _,v := range mmap {
//     vs = append(vs, v)
// }

// 返回的是数组，拼一下数据
//	for _, user := range users {
// 	s, ok := infoMap[*user.MT4Account] //检索
// 	if ok {
// 		s.Equity = user.Equity
// 	}
// }
