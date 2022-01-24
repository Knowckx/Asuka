package db

import (
	"fmt"

	"github.com/xormplus/xorm"
)

// 只插一个

// 只插入一条，返回其ID
func DBInsertOne(tx *xorm.Session, tname string, data interface{}) (int, error) {
	maxIDBefore, err := GetMaxIDFromTable(tx, tname)
	if err != nil {
		return 0, err
	}
	if _, err := tx.InsertOne(data); err != nil {
		return 0, err
	}
	maxIDNow, err := GetMaxIDFromTable(tx, tname)
	if err != nil {
		return 0, err
	}

	offset := maxIDNow - maxIDBefore
	if offset != 1 {
		return 0, fmt.Errorf("unexpect offset %d", offset)
	}
	return maxIDNow, nil
}
