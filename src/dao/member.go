package dao

import (
	"reserve_service/models"
)

func GetMember(mId int64) (error, *models.Member) {
	var member models.Member
	err := commonDb.Where("id = ?", mId).First(&member).Error
	return err, &member
}

func InsertMember(mId int64, name string) error {
	return commonDb.Create(&models.Member{
		ID:   mId,
		Name: name,
	}).Error
}
