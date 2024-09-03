package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ToFormValues 将结构体转换为 map[string]string
func ToFormValues(req interface{}) map[string]string {
	values := make(map[string]string)
	v := reflect.ValueOf(req)

	// 检查并解引用指针
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		formKey := fieldType.Tag.Get("form")
		if formKey == "" {
			continue
		}
		if field.Kind() == reflect.Slice && field.Len() == 0 {
			continue // 跳过空的slice
		}
		// 检查字段是否为nil（仅适用于指针和接口）
		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
			if field.IsNil() {
				continue
			}
			// 对于指针，解引用
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}
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
			values[formKey] = fmt.Sprint(field.Interface()) // 处理其他类型
		}
	}
	return values
}

// ToMapString 结构体转为Map[string]string
func ToMapString(in interface{}, tagName string) map[string]string {
	out := make(map[string]string)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ { // 遍历结构体字段
		fi := t.Field(i)
		tagValue := fi.Tag.Get(tagName) //获取标签里面的值 例如`form:"page"` 则获取到form里面的值 page
		if tagValue != "" {
			switch fi.Type.String() {
			case "Int", "Int8", "Int16", "Int32", "Int64":
				out[tagValue] = fmt.Sprintf("%d", v.Field(i).Int())
			case "Float32", "Float64":
				out[tagValue] = fmt.Sprintf("%f", v.Field(i).Float())
			case "String":
				out[tagValue] = v.Field(i).String()
			case "Bool":
				out[tagValue] = fmt.Sprintf("%v", v.Field(i).Bool())
			default:
				out[tagValue] = fmt.Sprintf("%v", v.Field(i).Interface())
			}
		}
	}
	return out
}
