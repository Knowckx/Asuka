package SnippetsGO

import (
	"fmt"
)

// 映射proto
type NewMod struct {
	UserID int
	Text   string
}

// 映射proto
type pbdsMod struct {
	UserID int
	Text   string
}

func NewModsFromProto(ins []*pbdsMod) []*NewMod {
	outs := []*NewMod{}
	for _, in := range ins {
		out := NewModFromProto(in)
		outs = append(outs, out)
	}
	return outs
}

func NewModFromProto(in *pbdsMod) *NewMod {
	out := &NewMod{}
	out.UserID = in.UserID
	//...
	return out
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
