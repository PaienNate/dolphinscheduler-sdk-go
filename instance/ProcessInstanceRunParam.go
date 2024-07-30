package instance

// ProcessInstanceRunParam 用于重新运行或恢复流程实例的参数结构体
type ProcessInstanceRunParam struct {
	ProcessInstanceId int64  `json:"processInstanceId"` // 流程实例ID
	ExecuteType       string `json:"executeType"`       // 执行类型
}
