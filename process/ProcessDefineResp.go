package process

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/util"
)

// Parameter 任务参数或全局参数
type Parameter struct {
	Prop   string `json:"prop"`   // 属性
	Value  string `json:"value"`  // 值
	Direct string `json:"direct"` // 进出参数标识: in 或 out
	Type   string `json:"type"`   // 数据类型: VARCHAR, INTEGER, LONG 等
}

// ProcessDefineResp 处理定义响应，参考 org.apache.dolphinscheduler.dao.entity.ProcessDefinition
type ProcessDefineResp struct {
	ID                   int               `json:"id"`                   // id
	Code                 int64             `json:"code"`                 // code
	Name                 string            `json:"name"`                 // name
	Version              int               `json:"version"`              // version
	ReleaseState         string            `json:"releaseState"`         // release state: online/offline
	ProjectCode          int64             `json:"projectCode"`          // project code
	Description          string            `json:"description"`          // description
	GlobalParams         string            `json:"globalParams"`         // user defined parameters
	GlobalParamList      []Parameter       `json:"globalParamList"`      // user defined parameter list
	GlobalParamMap       map[string]string `json:"globalParamMap"`       // user define parameter map
	CreateTime           util.Time         `json:"createTime"`           // create time
	UpdateTime           util.Time         `json:"updateTime"`           // update time
	Flag                 string            `json:"flag"`                 // process is valid: yes/no
	UserID               int               `json:"userId"`               // process user id
	UserName             string            `json:"userName"`             // user name
	ProjectName          string            `json:"projectName"`          // project name
	Locations            string            `json:"locations"`            // locations array for web
	ScheduleReleaseState string            `json:"scheduleReleaseState"` // schedule release state: online/offline
	Timeout              int               `json:"timeout"`              // process warning time out. unit: minute
	TenantID             int               `json:"tenantId"`             // tenant id
	TenantCode           string            `json:"tenantCode"`           // tenant code
	ModifyBy             string            `json:"modifyBy"`             // modify user name
	WarningGroupID       int               `json:"warningGroupId"`       // warningGroupId
	ExecutionType        string            `json:"executionType"`        // execution type
}
