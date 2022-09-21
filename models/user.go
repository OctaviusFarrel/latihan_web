package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Permission string
	Password   string `gorm:"" json:"-"`
}

func (UserModel) TableName() string {
	return "users"
}
