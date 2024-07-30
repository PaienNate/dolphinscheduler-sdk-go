package util

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Java代码等效转换的，所以很多名称其实没个B用
// TODO: 去除无效代码

// 创建一个新的 ObjectNode
func CreateObjectNode() json.RawMessage {
	node := make(map[string]interface{})
	raw, _ := json.Marshal(node)
	return raw
}

// 创建一个新的 ArrayNode
func CreateArrayNode() json.RawMessage {
	node := make([]interface{}, 0)
	raw, _ := json.Marshal(node)
	return raw
}

// 解析 JSON 字符串为 JsonNode
func ParseNode(jsonStr string) (json.RawMessage, error) {
	var node json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &node)
	if err != nil {
		return nil, fmt.Errorf("解析 JSON 节点失败: %w", err)
	}
	return node, nil
}

// 按路径表达式搜索 JsonNode
func SearchJsonNode(jsonNode json.RawMessage, path string) (json.RawMessage, error) {
	var node interface{}
	err := json.Unmarshal(jsonNode, &node)
	if err != nil {
		return nil, fmt.Errorf("解析 JSON 节点失败: %w", err)
	}

	keys := splitPath(path)
	for _, key := range keys {
		if m, ok := node.(map[string]interface{}); ok {
			node = m[key]
		} else {
			return nil, fmt.Errorf("路径错误，无法找到节点: %s", key)
		}
	}

	raw, _ := json.Marshal(node)
	return raw, nil
}

// 按路径表达式从字符串中搜索 JsonNode
func SearchJsonNodeFromString(jsonStr string, path string) (json.RawMessage, error) {
	node, err := ParseNode(jsonStr)
	if err != nil {
		return nil, err
	}
	return SearchJsonNode(node, path)
}

// 获取 ArrayNode
func GetAsArrayNode(jsonNode json.RawMessage, path string) (json.RawMessage, error) {
	node, err := SearchJsonNode(jsonNode, path)
	if err != nil {
		return nil, err
	}

	var arrayNode []interface{}
	err = json.Unmarshal(node, &arrayNode)
	if err != nil {
		return nil, fmt.Errorf("路径: %s 不是一个 JSON 数组", path)
	}

	raw, _ := json.Marshal(arrayNode)
	return raw, nil
}

// 获取 ArrayNode 从字符串
func GetAsArrayNodeFromString(jsonStr string, path string) (json.RawMessage, error) {
	node, err := SearchJsonNodeFromString(jsonStr, path)
	if err != nil {
		return nil, err
	}

	var arrayNode []interface{}
	err = json.Unmarshal(node, &arrayNode)
	if err != nil {
		return nil, fmt.Errorf("路径: %s 不是一个 JSON 数组", path)
	}

	raw, _ := json.Marshal(arrayNode)
	return raw, nil
}

// 获取字符串值
func GetAsText(jsonNode json.RawMessage, path string) (string, error) {
	node, err := SearchJsonNode(jsonNode, path)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(node, &result)
	if err != nil {
		return "", fmt.Errorf("路径: %s 不是一个字符串", path)
	}
	return result, nil
}

// 获取字符串值从字符串
func GetAsTextFromString(jsonStr string, path string) (string, error) {
	node, err := SearchJsonNodeFromString(jsonStr, path)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(node, &result)
	if err != nil {
		return "", fmt.Errorf("路径: %s 不是一个字符串", path)
	}
	return result, nil
}

// 获取长整型值
func GetAsLong(jsonNode json.RawMessage, path string) (int64, error) {
	node, err := SearchJsonNode(jsonNode, path)
	if err != nil {
		return 0, err
	}

	var result int64
	err = json.Unmarshal(node, &result)
	if err != nil {
		return 0, fmt.Errorf("路径: %s 不是一个长整型", path)
	}
	return result, nil
}

// 获取长整型值从字符串
func GetAsLongFromString(jsonStr string, path string) (int64, error) {
	node, err := SearchJsonNodeFromString(jsonStr, path)
	if err != nil {
		return 0, err
	}

	var result int64
	err = json.Unmarshal(node, &result)
	if err != nil {
		return 0, fmt.Errorf("路径: %s 不是一个长整型", path)
	}
	return result, nil
}

// 获取整型值
func GetAsInteger(jsonNode json.RawMessage, path string) (int, error) {
	node, err := SearchJsonNode(jsonNode, path)
	if err != nil {
		return 0, err
	}

	var result int
	err = json.Unmarshal(node, &result)
	if err != nil {
		return 0, fmt.Errorf("路径: %s 不是一个整型", path)
	}
	return result, nil
}

// 获取整型值从字符串
func GetAsIntegerFromString(jsonStr string, path string) (int, error) {
	node, err := SearchJsonNodeFromString(jsonStr, path)
	if err != nil {
		return 0, err
	}

	var result int
	err = json.Unmarshal(node, &result)
	if err != nil {
		return 0, fmt.Errorf("路径: %s 不是一个整型", path)
	}
	return result, nil
}

// 获取布尔值
func GetAsBoolean(jsonNode json.RawMessage, path string) (bool, error) {
	node, err := SearchJsonNode(jsonNode, path)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(node, &result)
	if err != nil {
		return false, fmt.Errorf("路径: %s 不是一个布尔值", path)
	}
	return result, nil
}

// 获取布尔值从字符串
func GetAsBooleanFromString(jsonStr string, path string) (bool, error) {
	node, err := SearchJsonNodeFromString(jsonStr, path)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(node, &result)
	if err != nil {
		return false, fmt.Errorf("路径: %s 不是一个布尔值", path)
	}
	return result, nil
}

// 将对象转换为 JSON 字符串
func ToJSONString(object interface{}) (string, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return "", fmt.Errorf("转换为 JSON 字符串失败: %w", err)
	}
	return string(data), nil
}

// 将 JSON 字符串解析为对象
func ParseObject[T any](jsonStr string) (T, error) {
	var obj T
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return obj, fmt.Errorf("解析 JSON 对象失败: %w", err)
	}
	return obj, nil
}

// 将 JSON 字符串流解析为对象
func ParseObjectFromStream[T any](stream io.Reader) (T, error) {
	var obj T
	err := json.NewDecoder(stream).Decode(&obj)
	if err != nil {
		return obj, fmt.Errorf("解析 JSON 对象失败: %w", err)
	}
	return obj, nil
}

// splitPath 分割路径字符串
func splitPath(path string) []string {
	// 使用 strings.Split 按 "." 分割路径字符串
	return strings.Split(path, ".")
}
