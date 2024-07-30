package process

import (
	"encoding/json"
)

// 定义执行类型常量
const (
	// EXECUTION_TYPE_PARALLEL 并行执行类型
	EXECUTION_TYPE_PARALLEL = "PARALLEL"
	// EXECUTION_TYPE_SERIAL_WAIT 串行等待执行类型
	EXECUTION_TYPE_SERIAL_WAIT = "SERIAL_WAIT"
	// EXECUTION_TYPE_SERIAL_DISCARD 串行丢弃执行类型
	EXECUTION_TYPE_SERIAL_DISCARD = "SERIAL_DISCARD"
	// EXECUTION_TYPE_SERIAL_PRIORITY 串行优先执行类型
	EXECUTION_TYPE_SERIAL_PRIORITY = "SERIAL_PRIORITY"
)

// ProcessDefineParam 定义了Dolphin Scheduler的进程参数
//
// Dolphin Scheduler 使用 POST 表单类型来创建进程，因此实际上每个属性都是字符串类型。
//
// 为了更方便的开发，使用了 TaskLocation、TaskDefinition 和 TaskRelation 对象，并重写了
// toString 方法，以便在发送请求时可以将这些对象转换为 JSON 字符串。
// Pinenutn: 需要将其转换为FORM表单
type ProcessDefineParam struct {
	// Name 工作流名称
	Name string `form:"name" json:"name"`

	// Locations 任务位置列表
	Locations []TaskLocation `form:"locations" json:"locations"`

	// TaskDefinitionJson 任务定义列表
	TaskDefinitionJson []TaskDefinition `form:"taskDefinitionJson" json:"taskDefinitionJson"`

	// TaskRelationJson 任务关系列表
	TaskRelationJson []TaskRelation `form:"taskRelationJson" json:"taskRelationJson"`

	// TenantCode 租户代码
	TenantCode string `form:"tenantCode" json:"tenantCode"`

	// Description 工作流描述
	Description string `form:"description" json:"description"`

	// ExecutionType 执行类型
	// 可能的值有 PARALLEL, SERIAL_WAIT, SERIAL_DISCARD, SERIAL_PRIORITY
	ExecutionType string `form:"executionType" json:"executionType"`

	// GlobalParams 全局参数列表
	GlobalParams []Parameter `form:"globalParams" json:"globalParams"`

	// Timeout 超时设置
	Timeout string `form:"timeout" json:"timeout"`
}

// NewParallelInstance 创建一个新的并行执行实例
func NewParallelInstance() *ProcessDefineParam {
	return &ProcessDefineParam{
		ExecutionType: EXECUTION_TYPE_PARALLEL,
	}
}

// NewSerialWaitInstance 创建一个新的串行等待执行实例
func NewSerialWaitInstance() *ProcessDefineParam {
	return &ProcessDefineParam{
		ExecutionType: EXECUTION_TYPE_SERIAL_WAIT,
	}
}

// NewSerialDiscardInstance 创建一个新的串行丢弃执行实例
func NewSerialDiscardInstance() *ProcessDefineParam {
	return &ProcessDefineParam{
		ExecutionType: EXECUTION_TYPE_SERIAL_DISCARD,
	}
}

// NewSerialPriorityInstance 创建一个新的串行优先执行实例
func NewSerialPriorityInstance() *ProcessDefineParam {
	return &ProcessDefineParam{
		ExecutionType: EXECUTION_TYPE_SERIAL_PRIORITY,
	}
}

// ToJSON 将 ProcessDefineParam 转换为 JSON 字符串
func (p *ProcessDefineParam) ToJSON() (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
