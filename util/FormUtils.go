package util

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// ToFormValues 将结构体转换为 map[string]string
func ToFormValues(req interface{}) map[string]string {
	values := make(map[string]string)
	v := reflect.ValueOf(req)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		formKey := fieldType.Tag.Get("form")
		if formKey == "" {
			continue
		}
		if field.Len() == 0 {
			continue // Skip empty slices
		}
		switch field.Kind() {
		case reflect.String:
			values[formKey] = field.String()
		case reflect.Int:
			values[formKey] = strconv.Itoa(int(field.Int()))
		case reflect.Bool:
			values[formKey] = strconv.FormatBool(field.Bool())
		case reflect.Slice:
			// 将结构体字段序列化为 JSON 字符串
			jsonData, err := json.Marshal(field.Interface())
			if err != nil {
				values[formKey] = "" // 处理错误的情况下可以设为空字符串
			} else {
				values[formKey] = string(jsonData)
			}
		default:
			values[formKey] = field.String()
		}
	}
	return values
}
