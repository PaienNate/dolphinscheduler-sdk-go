package process

import (
	"encoding/json"
)

// TaskLocation 代表任务位置
type TaskLocation struct {
	TaskCode int64 `json:"taskCode"` // 任务代码
	X        int   `json:"x"`        // x 坐标
	Y        int   `json:"y"`        // y 坐标
}

// NewTaskLocation 创建一个新的 TaskLocation，带有默认值
func NewTaskLocation() *TaskLocation {
	return &TaskLocation{}
}

// WithTaskCode 设置任务代码并返回更新后的 TaskLocation
func (t *TaskLocation) WithTaskCode(taskCode int64) *TaskLocation {
	t.TaskCode = taskCode
	return t
}

// WithX 设置 x 坐标并返回更新后的 TaskLocation
func (t *TaskLocation) WithX(x int) *TaskLocation {
	t.X = x
	return t
}

// WithY 设置 y 坐标并返回更新后的 TaskLocation
func (t *TaskLocation) WithY(y int) *TaskLocation {
	t.Y = y
	return t
}

// String 返回 TaskLocation 的 JSON 表示
func (t TaskLocation) String() string {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
