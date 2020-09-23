package models

type Member struct {
	ID   int64  `json:"-" gorm:"primary_key"`
	Name string `json:"-"`
}
