package mod

import (
	"time"

	"github.com/Knowckx/Asuka/curdinfm/pb"
)

//----------------- Mod NewObjMod start -----------------
type NewObjMod struct {
	ID         int        `xorm:"ID"`
	UserID     int        `xorm:"UserID"`
	NickName   string     `xorm:"NickName"`
	Lang       string     `xorm:"Lang"`
	CreateTime *time.Time `xorm:"CreateTime"`
	MT4Account `xorm:"extends"`
}

func (*NewObjMod) TableName() string {
	return "t_new_obj"
}

func NewNewObjModPB(in *pb.NewObjMod) *NewObjMod {
	out := &NewObjMod{}
	out.UserID = int(in.UserID)
	out.NickName = in.NickName
	out.Lang = in.Lang.String()
	// out.CreateTime = time.Unix(in.CreateTime, 0) //理论上不需要前端的CreateTime
	return out
}

func (in *NewObjMod) ToProto() *pb.NewObjMod {
	out := &pb.NewObjMod{}
	out.ID = int32(in.ID)
	out.Acc = in.MT4Account.ToProto()
	out.Lang = pb.LanguageType(pb.LanguageType_value[in.Lang])
	out.CreateTime = in.CreateTime.Unix()
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

//----------------- Mod NewObjModSearch Start -----------------
type NewObjModSearch struct {
	UserID   int32
	NickName string
	Account  string
	Page     *PageOp
}

func NewNewObjModSearchPB(in *pb.NewObjModSearch) *NewObjModSearch {
	out := &NewObjModSearch{}
	out.UserID = in.UserID
	out.NickName = in.NickName
	out.Account = in.Account
	out.Page = NewPageOpPB(in.Page)
	return out
}

//----------------- Mod NewObjModSearch end -----------------

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
