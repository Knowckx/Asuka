package SnippetsGO

import (
	"fmt"
)

// 新的结构体
type NewMod struct {
	UserID int
	Text   string
}

// func NewModsFromProto(ins []*pbmod.TestReq) []*NewMod {
// 	outs := []*NewMod{}
// 	for _, in := range ins {
// 		out := &NewMod{}
// 		out.New(in)
// 		outs = append(outs, out)
// 	}
// 	return outs
// }

// func (n *NewMod) New(in *TestReq) {
// 	n.UserID = in.UserID
// }

// func (n *NewMod) ToProto() *TestReq {
// 	out := &TestReq{}
// 	out.UserID = n.UserID
// 	return out
// }

// func NewModsToProto(ins []*NewMod) []*TestReq {
// 	outs := []*TestReq{}
// 	for _, in := range ins {
// 		out := in.ToProto()
// 		outs = append(outs, out)
// 	}
// 	return outs
// }

//-----------
func checkArgsNewMod(in *NewMod) error {
	if in.Text == "" {
		return fmt.Errorf("Args check:RankName is Empty")
	}
	if in.UserID == 0 {
		return fmt.Errorf("Args check:RankIndex is 0")
	}
	return nil
}
