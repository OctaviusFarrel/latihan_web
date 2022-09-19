package utils

import (
	"context"
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

func ValidateToken(token string, ctx context.Context) (result string, success bool) {
	result, err := tokenRepo.ValidateToken(token, ctx)
	if err != nil {
		success = false
	}
	return
}
