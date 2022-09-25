package models

import "gorm.io/gorm"

type TokenModel struct {
	gorm.Model
	HashToken      string `gorm:"type:VARCHAR(32)"`
	TokenBelongsTo uint
	User           UserModel `gorm:"foreignKey:TokenBelongsTo"`
	IsTokenActive  bool      `gorm:"column:token_active"`
}

func (TokenModel) TableName() string {
	return "user_tokens"
}
