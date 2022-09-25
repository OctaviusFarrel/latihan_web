package pgsql

import (
	"crypto/sha256"
	"fmt"

	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/requests"
)

type IUserRepo interface {
	CreateUser(userRequest requests.UserRequest) (user models.UserModel, err error)
	GetUserByUsername(userRequest requests.UserRequest) (user models.UserModel, err error)
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) CreateUser(userRequest requests.UserRequest) (user models.UserModel, err error) {
	user = models.UserModel{
		Username:   userRequest.Username,
		Permission: "",
		Password:   fmt.Sprintf("%x", sha256.Sum256([]byte(userRequest.Password))),
	}

	if err = dbPool.Model(user).Create(&user).Error; err != nil {
		return
	}

	return
}

func (repo *UserRepo) GetUserByUsername(userRequest requests.UserRequest) (user models.UserModel, err error) {
	err = dbPool.Select("id", "username", "permission", "created_at", "updated_at").Where(map[string]interface{}{
		"username": userRequest.Username,
		"password": fmt.Sprintf("%x", sha256.Sum256([]byte(userRequest.Password))),
	}).First(&user).Scan(&user).Error
	return
}
