package pgsql

import (
	"context"
	"fmt"
)

type ITokenRepo interface {
	InsertTokenByUser(userId int, token string) (result string, err error)
	ValidateToken(token string) (result string, err error)
}

type TokenRepo struct{}

func NewTokenRepo() ITokenRepo {
	return &TokenRepo{}
}

func (tokenRepo *TokenRepo) InsertTokenByUser(userId int, token string) (result string, err error) {
	sequence, err := dbPool.Query(context.Background(), "SELECT NEXTVAL('user_tokens_id_seq')")

	if err != nil {
		return
	}

	var value int
	sequence.Next()
	sequence.Scan(&value)

	_, err = dbPool.Query(context.Background(), "INSERT INTO user_tokens (id,token,user_id) VALUES ($1,$2,$3)", value, fmt.Sprintf("%d|%s", value, token), userId)

	if err != nil {
		return
	}

	data, err := dbPool.Query(context.Background(), "SELECT token FROM user_tokens WHERE id = $1", value)
	if err != nil {
		return
	}
	data.Next()
	data.Scan(&result)
	return
}

func (tokenRepo *TokenRepo) ValidateToken(token string) (result string, err error) {
	row, err := dbPool.Query(context.Background(), "SELECT user_id from user_tokens WHERE token = $1", token)
	if err != nil {
		return
	}

	if !row.Next() {
		err = fmt.Errorf("token is invalid")
		return
	}

	var i int
	row.Scan(&i)

	row, err = dbPool.Query(context.Background(), "SELECT permission from users WHERE id = $1", i)
	if err != nil {
		return
	}

	if !row.Next() {
		err = fmt.Errorf("token is invalid")
		return
	}

	row.Scan(&result)
	return
}
