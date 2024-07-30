package task

import (
	"encoding/json"
	"fmt"
)

// TaskResource 代表任务资源
type TaskResource struct {
	ID int64 `json:"id"`
}

// NewTaskResource 创建一个新的 TaskResource 实例
func NewTaskResource(id int64) *TaskResource {
	return &TaskResource{
		ID: id,
	}
}

// ToString 将 TaskResource 对象转换为 JSON 字符串
func (tr *TaskResource) ToString() (string, error) {
	jsonData, err := json.Marshal(tr)
	if err != nil {
		return "", fmt.Errorf("error converting TaskResource to JSON string: %v", err)
	}
	return string(jsonData), nil
}
