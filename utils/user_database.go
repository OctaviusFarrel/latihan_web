package utils

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"octaviusfarrel.dev/latihan_web/models"
	// . "octaviusfarrel.dev/latihan_web/models"
)

func CreateUser(user struct {
	Username   string
	Permission string
}, password string) bool {
	_, err := Dbpool.Query(context.Background(), "INSERT INTO users (username,password,permission) VALUES ($1,$2,$3)", user.Username, fmt.Sprintf("%x", sha256.Sum256([]byte(password))), user.Permission)

	return err == nil
}

func GetUserWithPassword(username string, password string) (models.User, error) {
	sqlScript := fmt.Sprintf("SELECT id,username FROM users WHERE username = '%s' AND password = '%s'", username, fmt.Sprintf("%x", sha256.Sum256([]byte(password))))
	rows, err := Dbpool.Query(context.Background(), sqlScript)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{}

	if rows.Next() {
		rows.Scan(&user.Id, &user.Username)
		return user, nil
	} else {
		return models.User{}, errors.New("user not found")
	}
}

func InsertTokenByUser(user_id int, token string) (string, error) {
	if sequence, err := Dbpool.Query(context.Background(), "SELECT NEXTVAL('user_tokens_id_seq')"); err != nil {
		return "", err
	} else {
		var value int
		sequence.Next()
		sequence.Scan(&value)
		_, err := Dbpool.Query(context.Background(), "INSERT INTO user_tokens (id,token,user_id) VALUES ($1,$2,$3)", value, fmt.Sprintf("%d|%s", value, token), user_id)
		if err != nil {
			return "", err
		} else {
			result, err := Dbpool.Query(context.Background(), "SELECT token FROM user_tokens WHERE id = $1", value)
			if err != nil {
				return "", err
			} else {
				var t string
				result.Next()
				result.Scan(&t)
				return t, nil
			}
		}
	}
}

func ValidateToken(token string) (string, bool) {
	if result, err := Dbpool.Query(context.Background(), "SELECT user_id from user_tokens WHERE token = $1", token); err != nil {
		fmt.Println(err)
		return "", false
	} else {
		if result.Next() {
			var i int
			result.Scan(&i)
			if result, err := Dbpool.Query(context.Background(), "SELECT permission from users WHERE id = $1", i); err != nil {
				return "", false
			} else {
				if result.Next() {
					var t string
					result.Scan(&t)
					return t, true
				} else {
					return "", false
				}
			}
		} else {
			return "", false
		}
	}
}
