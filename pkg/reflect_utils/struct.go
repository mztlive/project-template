package reflect_utils

import (
	"reflect"

	"github.com/mztlive/project-template/pkg/string_utils"
)

// GetNameFromStruct 获取结构体的名字
func GetNameFromStruct(s any) string {
	// 通过反射获取结构体的名字
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val.Type().Name()
}

// GetSnakeNameFromStruct 获取结构体的名字(蛇形命名)
func GetSnakeNameFromStruct(s any) string {
	// 通过反射获取结构体的名字
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return string_utils.ToSnakeCase(val.Type().Name())
}
