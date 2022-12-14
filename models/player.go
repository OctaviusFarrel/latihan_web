package models

import "gorm.io/gorm"

type PlayerModel struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int8   `json:"age" binding:"required"`
}

func (PlayerModel) TableName() string {
	return "players"
}
