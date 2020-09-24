package dao

import (
	"reserve_service/drivers"

	"gorm.io/gorm"
)

var commonDb *gorm.DB

func SetDriver() {
	commonDb = drivers.InitDb()
	Migration()
}

func GetCommonDb() *gorm.DB {
	return commonDb
}
