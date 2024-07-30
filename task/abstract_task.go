package task

import (
	"encoding/json"
	"github.com/PaienNate/dolphinscheduler-sdk-go/util"
)

// 用于定义Task的名称
type Task interface {
	GetTaskType() string
}

// 复现DolphinScheduler-JAVA-API
type AbstractTask struct {
	Dependence       json.RawMessage // 依赖字段 dependence field
	ConditionResult  json.RawMessage // 条件结果 condition result
	WaitStartTimeout json.RawMessage // 等待开始超时 wait start timeout
	SwitchResult     json.RawMessage // 开关结果 switch result
	TaskType         string          // 任务类型
}

// 初始化新的 AbstractTask
func NewAbstractTask(tasktype string) *AbstractTask {
	return &AbstractTask{
		Dependence:       util.CreateObjectNode(),
		ConditionResult:  util.CreateObjectNode(),
		WaitStartTimeout: util.CreateObjectNode(),
		SwitchResult:     util.CreateObjectNode(),
		TaskType:         tasktype,
	}
}

// GetTaskType 返回任务类型
func (at *AbstractTask) GetTaskType() string {
	return at.TaskType
}
