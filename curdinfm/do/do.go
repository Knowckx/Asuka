package do

import (
	"github.com/Knowckx/Asuka/curdinfm/db"
	"github.com/Knowckx/Asuka/curdinfm/mod"
	dbs "gitlab.followme.com/FollowmeGo/utils/db"
)

//----------------- NewObjMod CURD start -----------------

func AddNewObjMod(l *mod.NewObjMod) error {
	session := dbs.Session()
	err := db.AddNewObjMod(session, l)
	return err
}

func DelNewObjMod(l *mod.RankIndexWithLang) error {
	session := dbs.Session()
	err := db.DelNewObjMod(session, l.RankIndex, l.Lang)
	return err
}

func GetNewObjMods(se *mod.NewObjModSearch) (mod.NewObjMods, error) {
	session := dbs.Session()
	out, err := db.GetNewObjMods(session, se)
	return out, err
}

func UpdateNewObjMod(l *mod.NewObjMod) error {
	session := dbs.Session()
	err := db.UpdateNewObjMod(session, l)
	return err
}

//----------------- NewObjMod CURD start -----------------
