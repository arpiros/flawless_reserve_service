package dao

import "reserve_service/models"

func Migration() {
	commonDb.AutoMigrate(
		&models.Member{},
	)
}
