package dao

import (
	"reserve_service/drivers"

	"gorm.io/gorm"
)

var commonDb *gorm.DB

func SetDriver() {
	commonDb = drivers.InitDb()
}

func GetCommonDb() *gorm.DB {
	return commonDb
}
