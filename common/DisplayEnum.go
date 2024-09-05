package common

var ProcessInstanceStatusMap = map[string]string{
	"SUBMITTED_SUCCESS":    "提交成功",
	"RUNNING_EXECUTION":    "正在运行",
	"READY_PAUSE":          "准备暂停",
	"PAUSE":                "暂停",
	"READY_STOP":           "准备停止",
	"STOP":                 "停止",
	"FAILURE":              "失败",
	"SUCCESS":              "成功",
	"NEED_FAULT_TOLERANCE": "需要容错",
	"KILL":                 "KILL",
	"WAITING_THREAD":       "等待线程",
	"WAITING_DEPEND":       "等待依赖完成",
	"DELAY_EXECUTION":      "延时执行",
	"FORCED_SUCCESS":       "强制成功",
	"SERIAL_WAIT":          "串行等待",
	"DISPATCH":             "派发",
	"READY_BLOCK":          "准备阻断",
	"BLOCK":                "阻断",
	"WAIT_TO_RUN":          "等待执行",
}
