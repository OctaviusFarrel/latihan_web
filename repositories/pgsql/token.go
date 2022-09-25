package pgsql

import (
	"context"
	"fmt"
	"regexp"

	"octaviusfarrel.dev/latihan_web/models"
)

type ITokenRepo interface {
	InsertTokenByUser(tokenModel models.TokenModel) (err error)
	ValidateToken(token string, permission string, ctx context.Context) (err error)
}

type TokenRepo struct{}

func NewTokenRepo() *TokenRepo {
	return &TokenRepo{}
}

func (tokenRepo *TokenRepo) InsertTokenByUser(tokenModel models.TokenModel) (err error) {

	err = dbPool.Create(&tokenModel).Error
	return
}

func (tokenRepo *TokenRepo) ValidateToken(token string, permission string, ctx context.Context) (err error) {
	tokenModel := models.TokenModel{}

	err = dbPool.WithContext(ctx).First(&tokenModel, "hash_token = ?", token).Error

	if err != nil {
		return
	}

	err = dbPool.WithContext(ctx).First(&tokenModel.User, "id = ?", tokenModel.TokenBelongsTo).Error

	if err != nil {
		return
	}

	if len(regexp.MustCompile(tokenModel.User.Permission).FindString(permission)) == 0 {
		err = fmt.Errorf("token for id: %d", tokenModel.User.ID)
		return
	}

	return
}
