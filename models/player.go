package models

import "gorm.io/gorm"

type PlayerModel struct {
	gorm.Model
	Name string `form:"name" binding:"required"`
	Age  int8   `form:"age" binding:"required"`
}

func (PlayerModel) TableName() string {
	return "players"
}
