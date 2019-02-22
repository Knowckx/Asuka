package db

import (
	"fmt"

	"github.com/Knowckx/Asuka/curdinfm/mod"
	"github.com/xormplus/xorm"
)

//----------------- NewObjMod CURD start -----------------

// 增：一个配置
func AddNewObjMod(tx *xorm.Session, l *mod.NewObjMod) error {
	_, err := tx.InsertOne(l)
	return err
}

// 删：通过特定标识
func DelNewObjMod(tx *xorm.Session, RankIndex int, Lang string) error {
	count, err := tx.Where("RankIndex = ? and Lang = ?", RankIndex, Lang).Delete(new(mod.NewObjMod))
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("delete failed,affected row is 0")
	}
	return nil
}

// 查：一个 通过特定标识
func GetNewObjMod(tx *xorm.Session, RankIndex int, Lang string) (*mod.NewObjMod, error) {
	out := &mod.NewObjMod{}
	_, err := tx.Where("RankIndex = ? and Lang = ?", RankIndex, Lang).Get(out)
	return out, err
}

// 查：一组
func GetNewObjMods(tx *xorm.Session, RankIndex int) (mod.NewObjMods, error) {
	out := []*mod.NewObjMod{}
	err := tx.Where("RankIndex = ?", RankIndex).Find(&out)
	return out, err
}

// 查：更新
func UpdateNewObjMod(tx *xorm.Session, l *mod.NewObjMod) error {
	count, err := tx.Where("RankIndex = ? and Lang = ?", l.RankIndex, l.Lang).Update(l)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("Update failed,affected row is 0")
	}
	return err
}

//----------------- NewObjMod CURD end -----------------
