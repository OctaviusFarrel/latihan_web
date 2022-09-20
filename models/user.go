package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Id         int
	Username   string
	Permission string
}
