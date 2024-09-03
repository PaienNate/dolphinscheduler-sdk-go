package schedule

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/util"
)

type ScheduleInfoResp struct {
	Id                      int       `json:"id"`
	ProcessDefinitionCode   int64     `json:"processDefinitionCode"`
	ProcessDefinitionName   string    `json:"processDefinitionName"`
	ProjectName             string    `json:"projectName"`
	DefinitionDescription   string    `json:"definitionDescription,omitempty"`
	StartTime               util.Time `json:"startTime"`
	EndTime                 util.Time `json:"endTime"`
	TimezoneId              string    `json:"timezoneId"`
	Crontab                 string    `json:"crontab"`
	FailureStrategy         string    `json:"failureStrategy"`
	WarningType             string    `json:"warningType"`
	CreateTime              util.Time `json:"createTime"`
	UpdateTime              util.Time `json:"updateTime"`
	UserId                  int       `json:"userId"`
	UserName                string    `json:"userName"`
	ReleaseState            string    `json:"releaseState"`
	WarningGroupId          int       `json:"warningGroupId"`
	ProcessInstancePriority string    `json:"processInstancePriority"`
	WorkerGroup             string    `json:"workerGroup"`
	EnvironmentCode         *int64    `json:"environmentCode,omitempty"`
}
