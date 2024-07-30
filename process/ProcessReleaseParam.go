package process

// ProcessReleaseParam 表示流程发布参数
type ProcessReleaseParam struct {
	Name         string `form:"name" json:"name"`                 // 流程名称，Dolphin Scheduler REST API 并不需要此字段
	ReleaseState string `form:"releaseState" json:"releaseState"` // 流程状态：ONLINE 或 OFFLINE
}

// 定义常量
const (
	OnlineState  = "ONLINE"
	OfflineState = "OFFLINE"
)

// NewOnlineInstance 创建一个状态为在线的 ProcessReleaseParam 实例
func NewOnlineInstance() *ProcessReleaseParam {
	return &ProcessReleaseParam{
		ReleaseState: OnlineState,
	}
}

// NewOfflineInstance 创建一个状态为离线的 ProcessReleaseParam 实例
func NewOfflineInstance() *ProcessReleaseParam {
	return &ProcessReleaseParam{
		ReleaseState: OfflineState,
	}
}
