package do

import "github.com/Knowckx/Asuka/curdinfm/mod"

// 常用 把一组MT4账户 组装成索引map
func MakeAccMap(ins []*mod.MT4Account) map[mod.MT4Account]*mod.NewObjMod {
	rstMap := make(map[mod.MT4Account]*mod.NewObjMod) //组装索引map
	for _, in := range ins {
		objMod := new(mod.NewObjMod)
		objMod.MT4Account = in
		rstMap[*in] = objMod
	}
	return rstMap
}


type interface MT4AccountData {}



keys := make([]int, len(mmap))
for k := range mmap {
    keys = append(keys, k)
}

vs := make([]string, len(mmap))
for _,v := range mmap {
    vs = append(vs, v)
}