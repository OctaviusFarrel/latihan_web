package pgsql

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/requests"
)

type IUserRepo interface {
	CreateUser(user requests.UserRequest) (userResult models.UserModel, err error)
	GetUserByUsername(user requests.UserRequest) (result models.UserModel, err error)
}

type UserRepo struct{}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) CreateUser(user requests.UserRequest) (userResult models.UserModel, err error) {
	data, errVal := dbPool.Query(context.Background(), "SELECT NEXTVAL('users_id_seq')")

	if errVal != nil {
		err = errVal
		return
	}

	var value int
	if data.Next() {
		data.Scan(&value)
	} else {
		err = fmt.Errorf("could not get last index of database")
		return
	}

	_, err = dbPool.Query(context.Background(), "INSERT INTO users (id,username,password,permission) VALUES ($1,$2,$3,$4)", value, user.Username, fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password))), "")

	if err != nil {
		return
	}

	userResult = models.UserModel{
		Id:         value,
		Username:   user.Username,
		Permission: "",
	}
	return
}

func (repo *UserRepo) GetUserByUsername(userRequest requests.UserRequest) (models.UserModel, error) {
	sqlScript := fmt.Sprintf("SELECT id,username,permission FROM users WHERE username = '%s' AND password = '%s'", userRequest.Username, fmt.Sprintf("%x", sha256.Sum256([]byte(userRequest.Password))))
	rows, err := dbPool.Query(context.Background(), sqlScript)

	if err != nil {
		return models.UserModel{}, err
	}

	user := models.UserModel{}

	if rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Permission)
		return user, nil
	} else {
		return models.UserModel{}, errors.New("user not found")
	}
}
