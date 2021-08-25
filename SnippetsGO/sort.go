package SnippetsGO

import (
	"sort"
)

// 这个类型需要实现sort接口的Len，Less和Swap方法 | 这是官方包的例子
type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestSort() {
	ins := StringSlice([]string{"3", "2", "1"})
	sort.Sort(ins)
}
