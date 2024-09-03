package schedule

type ScheduleDefineParam struct {
	Schedule                *Schedule `json:"schedule,omitempty"`
	FailureStrategy         string    `json:"failureStrategy,omitempty"`
	WarningType             string    `json:"warningType,omitempty"`
	ProcessInstancePriority string    `json:"processInstancePriority,omitempty"`
	WarningGroupId          string    `json:"warningGroupId,omitempty"`
	WorkerGroup             string    `json:"workerGroup,omitempty"`
	EnvironmentCode         string    `json:"environmentCode,omitempty"`
	ProcessDefinitionCode   *int64    `json:"processDefinitionCode,omitempty"`
}

type Schedule struct {
	StartTime  string `json:"startTime,omitempty"`
	EndTime    string `json:"endTime,omitempty"`
	Crontab    string `json:"crontab,omitempty"`
	TimezoneId string `json:"timezoneId,omitempty"`
}

// NewScheduleDefineParam creates a new instance of ScheduleDefineParam with default values
func NewScheduleDefineParam() *ScheduleDefineParam {
	return &ScheduleDefineParam{
		FailureStrategy:         "END",
		WarningType:             "NONE",
		ProcessInstancePriority: "MEDIUM",
		WarningGroupId:          "0",
		WorkerGroup:             "default",
		EnvironmentCode:         "",
		Schedule: &Schedule{
			TimezoneId: "Asia/Shanghai",
		},
	}
}
