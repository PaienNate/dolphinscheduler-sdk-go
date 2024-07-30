package operator

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/remote"
)

// AbstractOperator 提供了一个抽象类，包含 Dolphin 调度器的基本操作。
type AbstractOperator struct {
	DolphinAddress     string
	token              string
	DolphinsRestClient *remote.DolphinsRestTemplate
}

// NewAbstractOperator 创建一个新的 AbstractOperator 实例。
func NewAbstractOperator(dolphinAddress, token string, dolphinsRestClient *remote.DolphinsRestTemplate) *AbstractOperator {
	return &AbstractOperator{
		DolphinAddress:     dolphinAddress,     // Dolphin 调度器的地址
		token:              token,              // 身份验证令牌
		DolphinsRestClient: dolphinsRestClient, // DolphinsRestTemplate 客户端
	}
}
