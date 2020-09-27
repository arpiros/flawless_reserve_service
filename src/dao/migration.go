package dao

import "reserve_service/models"

func Migration() {
	err := commonDb.AutoMigrate(
		&models.Member{},
	)
	if err != nil {
		panic("Failure Auto Migrate")
	}
}
