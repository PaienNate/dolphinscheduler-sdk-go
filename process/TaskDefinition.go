package process

import (
	"encoding/json"
	"github.com/PaienNate/dolphinscheduler-sdk-go/task"
)

// TaskDefinition 任务定义结构体
type TaskDefinition struct {
	Code                  int64     `json:"code"`                  // 任务代码
	Version               int       `json:"version"`               // 版本
	Name                  string    `json:"name"`                  // 任务节点名称
	Description           string    `json:"description"`           // 任务节点描述
	TaskType              string    `json:"taskType"`              // 任务类型，从 AbstractTask#getTaskType() 获取
	TaskParams            task.Task `json:"taskParams"`            // 任务参数
	Flag                  string    `json:"flag"`                  // 执行标志，NO: 不执行; YES: 执行; 默认值为 YES
	TaskPriority          string    `json:"taskPriority"`          // 任务优先级
	WorkerGroup           string    `json:"workerGroup"`           // 工作组
	FailRetryTimes        string    `json:"failRetryTimes"`        // 失败重试次数
	FailRetryInterval     string    `json:"failRetryInterval"`     // 失败重试间隔
	TimeoutFlag           string    `json:"timeoutFlag"`           // 超时标志
	TimeoutNotifyStrategy string    `json:"timeoutNotifyStrategy"` // 超时通知策略
	Timeout               int       `json:"timeout"`               // 超时时间，默认值为 0
	DelayTime             string    `json:"delayTime"`             // 延迟时间，默认值为 "0"
	EnvironmentCode       int       `json:"environmentCode"`       // 环境代码，默认值为 -1
	TaskExecuteType       string    `json:"taskExecuteType"`       // 任务执行类型
	CpuQuota              int       `json:"cpuQuota"`              // CPU 配额，默认值为 -1
	MemoryMax             int64     `json:"memoryMax"`             // 最大内存，默认值为 -1L
	IsCache               string    `json:"isCache"`               // 缓存标志，YES 或 NO，默认值为 NO
}

// String 实现 Stringer 接口，返回对象的 JSON 字符串
func (t TaskDefinition) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(data)
}
