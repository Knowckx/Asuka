package db

import (
	"fmt"

	"github.com/Knowckx/Asuka/curdinfm/mod"
	"github.com/xormplus/xorm"
)

//----------------- NewObjMod CURD start -----------------

// 增：一条数据
func AddNewObjMod(tx *xorm.Session, in *mod.NewObjMod) error {
	_, err := tx.InsertOne(in)
	if err != nil {
		return fmt.Errorf("InsertOne To DB failed,%s", err)
	}
	return nil
}

// 增：批量
func AddNewObjMods(tx *xorm.Session, l []*mod.NewObjMod) error {
	if len(l) == 0 {
		return nil
	}
	_, err := tx.Insert(&l)
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

// 删：很多账户
func DelAccs(tx *xorm.Session, accs mod.MT4Accounts) error {
	if len(accs) == 0 {
		return nil
	}
	for _, acc := range accs {
		tx = tx.Or("BrokerID = ? and Account = ?", acc.BrokerID, acc.Account)
	}
	count, err := tx.Delete(new(mod.NewObjMod))
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("delete result:affected row is 0")
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
func GetNewObjMods(tx *xorm.Session, se *mod.NewObjModSearch) (mod.NewObjMods, error) {
	out := mod.NewObjMods{}
	if se != nil {
		if se.UserID != 0 {
			tx.And("UserID = ?", se.UserID)
		}
		if se.NickName != "" {
			tx.And("NickName = ?", se.NickName)
		}
	}
	err := tx.Find(&out)
	return out, err
}

// 查：一组  来很多MT4账户
func GetUsers(tx *xorm.Session, ins mod.MT4Accounts) (mod.NewObjMods, error) {
	out := []*mod.NewObjMod{}
	if ins == nil || len(ins) == 0 { //防止查全表
		return out, nil
	}
	for _, in := range ins {
		tx = tx.Or("BrokerID = ? and Account = ?", in.BrokerID, in.Account)
	}
	err := tx.Find(&out)
	return out, err
}

// 查：更新
func UpdateNewObjMod(tx *xorm.Session, l *mod.NewObjMod) error {
	count, err := tx.Where("BrokerID = ? and Lang = ?", l.BrokerID, l.Lang).Update(l)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("Update failed,affected row is 0")
	}
	return err
}

// 联表查
func GetJoin(tx *xorm.Session) (mod.NewObjMods, error) {
	out := []*mod.NewObjMod{}
	tx = tx.Where("broker_id > ?", 3)
	tx = tx.And("user_type = ?", 1)
	tx.Join("INNER", "users", "users.id = user_accounts.user_id")
	err := tx.Find(&out)
	return out, err
}

//----------------- NewObjMod CURD end -----------------
