package SnippetsGO

import (
	"fmt"
	"sort"
)

type person struct {
	ID   int
	name string
}

type Users []*person

// 这个类型需要实现sort接口的Len，Less和Swap方法
func (s Users) Len() int {
	return len(s)
}
func (s Users) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Users) Less(i, j int) bool {
	return s[i].ID > s[j].ID
}

func TestSort() {
	P1 := &person{ID: 1, name: "A1"}
	P2 := &person{ID: 2, name: "B2"}

	people := []*person{P1, P2}
	sort.Sort(Users(people))
	fmt.Println(people[0], people[1])
}
