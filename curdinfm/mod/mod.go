package mod

import (
	"time"

	"github.com/Knowckx/Asuka/curdinfm/pb"
)

//----------------- Mod NewObjMod start -----------------
type NewObjMod struct {
	ID         int `xorm:"ID"`
	MT4Account *MT4Account
	RankIndex  int    `xorm:"RankIndex"`
	Lang       string `xorm:"Lang"`
	BindTime   time.Time
}

func (*NewObjMod) TableName() string {
	return "t_new_obj"
}

func NewNewObjMod(in *pb.NewObjMod) *NewObjMod {
	out := &NewObjMod{}
	out.RankIndex = int(in.RankIndex)
	out.Lang = in.Lang.String()
	out.BindTime = time.Unix(in.BindTime, 0)
	return out
}

func (in *NewObjMod) ToProto() *pb.NewObjMod {
	out := &pb.NewObjMod{}
	out.MT4Account = in.MT4Account.ToProto()
	out.RankIndex = int32(in.RankIndex)
	out.Lang = pb.LanguageType(pb.LanguageType_value[in.Lang])
	out.BindTime = in.BindTime.Unix()
	return out
}

type NewObjMods []*NewObjMod

func (ins *NewObjMods) ToProto() []*pb.NewObjMod {
	outs := []*pb.NewObjMod{}
	for _, in := range *ins {
		out := in.ToProto()
		outs = append(outs, out)
	}
	return outs
}

//----------------- Mod NewObjMod end -----------------

// --- 为了测试
type RankIndexWithLang struct {
	RankIndex int    `xorm:"RankIndex"`
	Lang      string `xorm:"Lang"`
}

func NewRankIndexWithLang(in *pb.RankIndexWithLang) *RankIndexWithLang {
	out := &RankIndexWithLang{}
	out.RankIndex = int(in.RankIndex)
	out.Lang = in.Lang.String()
	return out
}

// MaxID 一张表的最大ID
type MaxID struct {
	ID int `xorm:"ID"`
}
