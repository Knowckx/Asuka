package SnippetsGO

import (
	"fmt"
)

// 新的结构体
type NewMod struct {
	UserID int
	Text   string
}

// 映射proto
type pbMod struct {
	UserID int
	Text   string
}

func NewModsFromProto(ins []*pbMod) []*NewMod {
	outs := []*NewMod{}
	for _, in := range ins {
		out := &NewMod{}
		out.New(in)
		outs = append(outs, out)
	}
	return outs
}

func (n *NewMod) New(in *pbMod) {
	n.UserID = in.UserID
}

func (n *NewMod) ToProto() *pbMod {
	out := &pbMod{}
	out.UserID = n.UserID
	return out
}

func NewModsToProto(ins []*NewMod) []*pbMod {
	outs := []*pbMod{}
	for _, in := range ins {
		out := in.ToProto()
		outs = append(outs, out)
	}
	return outs
}

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

// func NewDoubleRange(min, max float64) *DoubleRange {
// 	out := &DoubleRange{}
// 	out.Max = max
// 	out.Min = min
// 	return out
// }
