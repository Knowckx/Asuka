package handle

import (
	"context"
	"fmt"

	"github.com/Knowckx/Asuka/curdinfm/do"
	"github.com/Knowckx/Asuka/curdinfm/mod"
	"github.com/Knowckx/Asuka/curdinfm/pb"
)

// ----------------- 排行榜配置外语翻译文本 start -----------------
// OA 增加 榜单配置对应的外语文本
func (*DatastatisticSrv) AddNewObjMod(ctx context.Context, in *pb.NewObjMod, out *pb.EmptyMessage) error {
	err := CheckArgsMT4Account(in.Acc)
	if err != nil {
		return err
	}
	args := mod.NewNewObjModPB(in)
	err = do.AddNewObjMod(args)
	return err
}

func (*DatastatisticSrv) DelNewObjMod(ctx context.Context, in *pb.RankIndexWithLang, out *pb.EmptyMessage) error {
	err := checkArgsNewObjMod(in.RankIndex, int32(in.Lang))
	if err != nil {
		return err
	}
	cfg := mod.NewRankIndexWithLang(in)
	err = do.DelNewObjMod(cfg)
	return err
}

func (*DatastatisticSrv) GetNewObjMods(ctx context.Context, in *pb.NewObjModSearch, reply *pb.NewObjMods) error {
	ars := mod.NewNewObjModSearchPB(in)
	out, err := do.GetNewObjMods(ars)
	if err != nil {
		return err
	}
	reply.List = out.ToProto()
	return nil
}

func (*DatastatisticSrv) UpdateNewObjMod(ctx context.Context, in *pb.NewObjMod, reply *pb.EmptyMessage) error {
	err := checkArgsNewObjMod(in.ID, int32(in.Lang))
	if err != nil {
		return err
	}
	cfg := mod.NewNewObjModPB(in)
	err = do.UpdateNewObjMod(cfg)
	if err != nil {
		return err
	}
	return nil
}

func checkArgsNewObjMod(RankIndex int32, Lang int32) error {
	if RankIndex == 0 {
		return fmt.Errorf("Args check:rank index should not be 0")
	}
	_, ok := pb.LanguageType_name[Lang]
	if !ok {
		return fmt.Errorf("Args check:unknown LanguageType")
	}
	return nil
}

// ----------------- 排行榜配置外语翻译文本 end -----------------

// ----------------- 参数检查 Start -----------------
func CheckArgsMT4Account(in *pb.MT4Account) error {
	err := fmt.Errorf("Args Check failed:MT4Account")
	if in == nil || in.BrokerID == 0 || in.Account == "" {
		return err
	}
	return nil
}

// ----------------- 参数检查 end -----------------
type DatastatisticSrv struct {
}
