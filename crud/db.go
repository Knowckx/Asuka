package crud

import "gorm.io/gorm"

var defaultDB *gorm.DB
var CanWriteDB = true

func GetDefault() *gorm.DB {
	return defaultDB
}
