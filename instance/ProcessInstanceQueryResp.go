package instance

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
	"time"
)

// ProcessInstanceQueryResp 用于查询流程实例的响应结构体
type ProcessInstanceQueryResp struct {
	ID                       int                        `json:"id" form:"id"`
	ProcessDefinitionCode    int64                      `json:"processDefinitionCode" form:"processDefinitionCode"`
	ProcessDefinitionVersion int                        `json:"processDefinitionVersion" form:"processDefinitionVersion"`
	State                    string                     `json:"state" form:"state"`
	Recovery                 string                     `json:"recovery" form:"recovery"`
	StartTime                time.Time                  `json:"startTime" form:"startTime"`
	EndTime                  time.Time                  `json:"endTime" form:"endTime"`
	RunTimes                 int                        `json:"runTimes" form:"runTimes"`
	Name                     string                     `json:"name" form:"name"`
	Host                     string                     `json:"host" form:"host"`
	ProcessDefinition        *process.ProcessDefineResp `json:"processDefinition" form:"processDefinition"`
	CommandType              string                     `json:"commandType" form:"commandType"`
	CommandParam             string                     `json:"commandParam" form:"commandParam"`
	TaskDependType           string                     `json:"taskDependType" form:"taskDependType"`
	MaxTryTimes              int                        `json:"maxTryTimes" form:"maxTryTimes"`
	FailureStrategy          string                     `json:"failureStrategy" form:"failureStrategy"`
	WarningType              string                     `json:"warningType" form:"warningType"`
	WarningGroupId           *int                       `json:"warningGroupId" form:"warningGroupId"`
	ScheduleTime             time.Time                  `json:"scheduleTime" form:"scheduleTime"`
	CommandStartTime         time.Time                  `json:"commandStartTime" form:"commandStartTime"`
	GlobalParams             string                     `json:"globalParams" form:"globalParams"`
	ExecutorId               int                        `json:"executorId" form:"executorId"`
	ExecutorName             string                     `json:"executorName" form:"executorName"`
	TenantCode               string                     `json:"tenantCode" form:"tenantCode"`
	Queue                    string                     `json:"queue" form:"queue"`
	IsSubProcess             string                     `json:"isSubProcess" form:"isSubProcess"`
	Locations                string                     `json:"locations" form:"locations"`
	HistoryCmd               string                     `json:"historyCmd" form:"historyCmd"`
	DependenceScheduleTimes  string                     `json:"dependenceScheduleTimes" form:"dependenceScheduleTimes"`
	Duration                 string                     `json:"duration" form:"duration"`
	ProcessInstancePriority  string                     `json:"processInstancePriority" form:"processInstancePriority"`
	WorkerGroup              string                     `json:"workerGroup" form:"workerGroup"`
	EnvironmentCode          *int64                     `json:"environmentCode" form:"environmentCode"`
	Timeout                  int                        `json:"timeout" form:"timeout"`
	TenantId                 int                        `json:"tenantId" form:"tenantId"`
	VarPool                  string                     `json:"varPool" form:"varPool"`
	DryRun                   int                        `json:"dryRun" form:"dryRun"`
	RestartTime              time.Time                  `json:"restartTime" form:"restartTime"`
}
