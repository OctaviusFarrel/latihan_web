package utils

import (
	"fmt"
	"reflect"
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
