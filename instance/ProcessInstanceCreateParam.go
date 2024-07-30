package instance

// ProcessInstanceCreateParam 用于创建流程实例的参数
type ProcessInstanceCreateParam struct {
	FailureStrategy           string `json:"failureStrategy"`
	ProcessDefinitionCode     int64  `json:"processDefinitionCode"`
	ProcessInstancePriority   string `json:"processInstancePriority"`
	ScheduleTime              string `json:"scheduleTime"`
	WarningGroupId            int64  `json:"warningGroupId"`
	WarningType               string `json:"warningType"`
	DryRun                    int    `json:"dryRun"`
	EnvironmentCode           string `json:"environmentCode"`
	ExecType                  string `json:"execType"`
	ExpectedParallelismNumber string `json:"expectedParallelismNumber"`
	RunMode                   string `json:"runMode"`
	StartNodeList             string `json:"startNodeList"`
	StartParams               string `json:"startParams"`
	TaskDependType            string `json:"taskDependType"`
	WorkerGroup               string `json:"workerGroup"`
}

// NewProcessInstanceCreateParam 创建一个新的 ProcessInstanceCreateParam 实例
func NewProcessInstanceCreateParam() *ProcessInstanceCreateParam {
	return &ProcessInstanceCreateParam{}
}

// SetFailureStrategy 设置 FailureStrategy 并返回当前实例
func (p *ProcessInstanceCreateParam) SetFailureStrategy(failureStrategy string) *ProcessInstanceCreateParam {
	p.FailureStrategy = failureStrategy
	return p
}

// SetProcessDefinitionCode 设置 ProcessDefinitionCode 并返回当前实例
func (p *ProcessInstanceCreateParam) SetProcessDefinitionCode(processDefinitionCode int64) *ProcessInstanceCreateParam {
	p.ProcessDefinitionCode = processDefinitionCode
	return p
}

// SetProcessInstancePriority 设置 ProcessInstancePriority 并返回当前实例
func (p *ProcessInstanceCreateParam) SetProcessInstancePriority(processInstancePriority string) *ProcessInstanceCreateParam {
	p.ProcessInstancePriority = processInstancePriority
	return p
}

// SetScheduleTime 设置 ScheduleTime 并返回当前实例
func (p *ProcessInstanceCreateParam) SetScheduleTime(scheduleTime string) *ProcessInstanceCreateParam {
	p.ScheduleTime = scheduleTime
	return p
}

// SetWarningGroupId 设置 WarningGroupId 并返回当前实例
func (p *ProcessInstanceCreateParam) SetWarningGroupId(warningGroupId int64) *ProcessInstanceCreateParam {
	p.WarningGroupId = warningGroupId
	return p
}

// SetWarningType 设置 WarningType 并返回当前实例
func (p *ProcessInstanceCreateParam) SetWarningType(warningType string) *ProcessInstanceCreateParam {
	p.WarningType = warningType
	return p
}

// SetDryRun 设置 DryRun 并返回当前实例
func (p *ProcessInstanceCreateParam) SetDryRun(dryRun int) *ProcessInstanceCreateParam {
	p.DryRun = dryRun
	return p
}

// SetEnvironmentCode 设置 EnvironmentCode 并返回当前实例
func (p *ProcessInstanceCreateParam) SetEnvironmentCode(environmentCode string) *ProcessInstanceCreateParam {
	p.EnvironmentCode = environmentCode
	return p
}

// SetExecType 设置 ExecType 并返回当前实例
func (p *ProcessInstanceCreateParam) SetExecType(execType string) *ProcessInstanceCreateParam {
	p.ExecType = execType
	return p
}

// SetExpectedParallelismNumber 设置 ExpectedParallelismNumber 并返回当前实例
func (p *ProcessInstanceCreateParam) SetExpectedParallelismNumber(expectedParallelismNumber string) *ProcessInstanceCreateParam {
	p.ExpectedParallelismNumber = expectedParallelismNumber
	return p
}

// SetRunMode 设置 RunMode 并返回当前实例
func (p *ProcessInstanceCreateParam) SetRunMode(runMode string) *ProcessInstanceCreateParam {
	p.RunMode = runMode
	return p
}

// SetStartNodeList 设置 StartNodeList 并返回当前实例
func (p *ProcessInstanceCreateParam) SetStartNodeList(startNodeList string) *ProcessInstanceCreateParam {
	p.StartNodeList = startNodeList
	return p
}

// SetStartParams 设置 StartParams 并返回当前实例
func (p *ProcessInstanceCreateParam) SetStartParams(startParams string) *ProcessInstanceCreateParam {
	p.StartParams = startParams
	return p
}

// SetTaskDependType 设置 TaskDependType 并返回当前实例
func (p *ProcessInstanceCreateParam) SetTaskDependType(taskDependType string) *ProcessInstanceCreateParam {
	p.TaskDependType = taskDependType
	return p
}

// SetWorkerGroup 设置 WorkerGroup 并返回当前实例
func (p *ProcessInstanceCreateParam) SetWorkerGroup(workerGroup string) *ProcessInstanceCreateParam {
	p.WorkerGroup = workerGroup
	return p
}
