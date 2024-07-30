package process

import (
	"encoding/json"
)

// TaskRelation 任务关系结构体
type TaskRelation struct {
	Name            string `json:"name"`            // 名称
	PreTaskCode     int64  `json:"preTaskCode"`     // 前置任务代码
	PreTaskVersion  int    `json:"preTaskVersion"`  // 前置任务版本
	PostTaskCode    int64  `json:"postTaskCode"`    // 后置任务代码
	PostTaskVersion int    `json:"postTaskVersion"` // 后置任务版本
	ConditionType   int    `json:"conditionType"`   // 条件类型
	ConditionParams *int   `json:"conditionParams"` // 条件参数
}

// String 实现 Stringer 接口，返回对象的 JSON 字符串
func (t TaskRelation) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(data)
}
