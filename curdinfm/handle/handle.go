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
	err := checkArgsNewObjMod(in.RankIndex, int32(in.Lang))
	if err != nil {
		return err
	}
	cfg := mod.NewNewObjMod(in)

	err = do.AddNewObjMod(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (*DatastatisticSrv) DelNewObjMod(ctx context.Context, in *pb.RankIndexWithLang, out *pb.EmptyMessage) error {
	err := checkArgsNewObjMod(in.RankIndex, int32(in.Lang))
	if err != nil {
		return err
	}
	cfg := mod.NewRankIndexWithLang(in)
	err = do.DelNewObjMod(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (*DatastatisticSrv) GetNewObjMods(ctx context.Context, in *pb.RankIndexWithLang, reply *pb.NewObjMods) error {
	if in.RankIndex == 0 {
		return fmt.Errorf("rank index should not be 0")
	}
	out, err := do.GetNewObjMods(int(in.RankIndex))
	if err != nil {
		return err
	}
	reply.Mods = out.ToProto()
	return nil
}

func (*DatastatisticSrv) UpdateNewObjMod(ctx context.Context, in *pb.NewObjMod, reply *pb.EmptyMessage) error {
	err := checkArgsNewObjMod(in.RankIndex, int32(in.Lang))
	if err != nil {
		return err
	}
	cfg := mod.NewNewObjMod(in)
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

type DatastatisticSrv struct {
}
