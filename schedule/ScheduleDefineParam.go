package schedule

type ScheduleDefineParam struct {
	Schedule                *Schedule `json:"schedule,omitempty" form:"schedule"`
	FailureStrategy         string    `json:"failureStrategy,omitempty" form:"failureStrategy"`
	WarningType             string    `json:"warningType,omitempty" form:"warningType"`
	ProcessInstancePriority string    `json:"processInstancePriority,omitempty" form:"processInstancePriority"`
	WarningGroupId          string    `json:"warningGroupId,omitempty" form:"warningGroupId"`
	WorkerGroup             string    `json:"workerGroup,omitempty" form:"workerGroup"`
	EnvironmentCode         string    `json:"environmentCode,omitempty" form:"environmentCode"`
	ProcessDefinitionCode   *int64    `json:"processDefinitionCode,omitempty" form:"processDefinitionCode"`
}

type Schedule struct {
	StartTime  string `json:"startTime,omitempty" form:"startTime"`
	EndTime    string `json:"endTime,omitempty" form:"endTime"`
	Crontab    string `json:"crontab,omitempty" form:"crontab"`
	TimezoneId string `json:"timezoneId'" form:"timezoneId"`
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
