package util

import (
	"os"
	"reflect"
)


func IsNilPointer(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() != reflect.Ptr || rv.IsNil()
}

func GetEnvStr(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}