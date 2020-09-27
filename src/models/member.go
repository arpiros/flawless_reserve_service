package models

type Member struct {
	ID    int64  `json:"id" form:"id" query:"id" gorm:"primary_key" `
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}
