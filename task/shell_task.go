package task

// ShellTask 是一个具体的任务，嵌入了 AbstractTask
type ShellTask struct {
	*AbstractTask
	ResourceList []TaskResource `json:"resourceList"`
	//LocalParams  []Parameter `json:"localParams"`
	RawScript string `json:"rawScript"`
}

// NewShellTask 创建一个新的 ShellTask 实例
func NewShellTask() *ShellTask {
	return &ShellTask{
		AbstractTask: NewAbstractTask("SHELL"),
		ResourceList: []TaskResource{},
		//LocalParams:  []process.Parameter{},
	}
}

// GetTaskType 返回任务类型
func (st *ShellTask) GetTaskType() string {
	return "SHELL"
}
