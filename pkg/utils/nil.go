package utils

import "reflect"

func IsInterfaceNil(value interface{}) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}
