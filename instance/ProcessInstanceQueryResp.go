package instance

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
	"time"
)

// ProcessInstanceQueryResp 用于查询流程实例的响应结构体
type ProcessInstanceQueryResp struct {
	ID                       int                        `json:"id"`
	ProcessDefinitionCode    int64                      `json:"processDefinitionCode"`
	ProcessDefinitionVersion int                        `json:"processDefinitionVersion"`
	State                    string                     `json:"state"`
	Recovery                 string                     `json:"recovery"`
	StartTime                time.Time                  `json:"startTime"`
	EndTime                  time.Time                  `json:"endTime"`
	RunTimes                 int                        `json:"runTimes"`
	Name                     string                     `json:"name"`
	Host                     string                     `json:"host"`
	ProcessDefinition        *process.ProcessDefineResp `json:"processDefinition"`
	CommandType              string                     `json:"commandType"`
	CommandParam             string                     `json:"commandParam"`
	TaskDependType           string                     `json:"taskDependType"`
	MaxTryTimes              int                        `json:"maxTryTimes"`
	FailureStrategy          string                     `json:"failureStrategy"`
	WarningType              string                     `json:"warningType"`
	WarningGroupId           *int                       `json:"warningGroupId"`
	ScheduleTime             time.Time                  `json:"scheduleTime"`
	CommandStartTime         time.Time                  `json:"commandStartTime"`
	GlobalParams             string                     `json:"globalParams"`
	ExecutorId               int                        `json:"executorId"`
	ExecutorName             string                     `json:"executorName"`
	TenantCode               string                     `json:"tenantCode"`
	Queue                    string                     `json:"queue"`
	IsSubProcess             string                     `json:"isSubProcess"`
	Locations                string                     `json:"locations"`
	HistoryCmd               string                     `json:"historyCmd"`
	DependenceScheduleTimes  string                     `json:"dependenceScheduleTimes"`
	Duration                 string                     `json:"duration"`
	ProcessInstancePriority  string                     `json:"processInstancePriority"`
	WorkerGroup              string                     `json:"workerGroup"`
	EnvironmentCode          *int64                     `json:"environmentCode"`
	Timeout                  int                        `json:"timeout"`
	TenantId                 int                        `json:"tenantId"`
	VarPool                  string                     `json:"varPool"`
	DryRun                   int                        `json:"dryRun"`
	RestartTime              time.Time                  `json:"restartTime"`
}
