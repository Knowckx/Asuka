package db

import "github.com/xormplus/xorm"

func Session() *xorm.Session {
	return &xorm.Session{}
}
