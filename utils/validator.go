package utils

import (
	"fmt"
	"reflect"

	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
)

func IsTypeCorrect[T interface{}](value interface{}, debug bool) bool {
	if debug {
		fmt.Println(reflect.TypeOf(value))
	}
	switch value.(type) {
	case T:
		return true
	default:
		return false
	}
}

var tokenRepo pgsql.ITokenRepo = pgsql.NewTokenRepo()

func ValidateToken(token string) (result string, success bool) {
	result, err := tokenRepo.ValidateToken(token)
	if err != nil {
		success = false
	}
	return
}
