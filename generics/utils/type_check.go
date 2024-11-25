package utils

import (
	"reflect"
)

func IsInstanceOf[T any](r any) bool {
	ofType := reflect.TypeOf((*T)(nil)).Elem()
	instType := reflect.TypeOf(r)

	if ofType == instType {
		return true
	}
	return false
}
