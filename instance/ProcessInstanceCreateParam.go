package instance

// ProcessInstanceCreateParam 用于创建流程实例的参数
type ProcessInstanceCreateParam struct {
	FailureStrategy           string `json:"failureStrategy" form:"failureStrategy"`
	ProcessDefinitionCode     int64  `json:"processDefinitionCode" form:"processDefinitionCode"`
	ProcessInstancePriority   string `json:"processInstancePriority" form:"processInstancePriority"`
	ScheduleTime              string `json:"scheduleTime" form:"scheduleTime"`
	WarningGroupId            int64  `json:"warningGroupId,omitempty" form:"warningGroupId"`
	WarningType               string `json:"warningType" form:"warningType"`
	DryRun                    int    `json:"dryRun" form:"dryRun"`
	EnvironmentCode           string `json:"environmentCode" form:"environmentCode"`
	ExecType                  string `json:"execType" form:"execType"`
	ExpectedParallelismNumber string `json:"expectedParallelismNumber" form:"expectedParallelismNumber"`
	RunMode                   string `json:"runMode" form:"runMode"`
	StartNodeList             string `json:"startNodeList" form:"startNodeList"`
	StartParams               string `json:"startParams" form:"startParams"`
	TaskDependType            string `json:"taskDependType" form:"taskDependType"`
	WorkerGroup               string `json:"workerGroup" form:"workerGroup"`
	// 必须添加参数：租户ID，因为不指定租户必定执行报错
	TenantCode string `json:"tenantCode" form:"tenantCode"`
}
