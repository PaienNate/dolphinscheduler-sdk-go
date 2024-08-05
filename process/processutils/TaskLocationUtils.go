package processutils

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
)

// HorizontalLocation 生成任务位置列表，节点在水平方向上排列
// 如 (100, 200), (300, 200), (500, 200)
func HorizontalLocation(taskCodes ...int64) []process.TaskLocation {
	beginX := 100
	y := 100
	xStep := 300
	return HorizontalLocationCustom(beginX, y, xStep, taskCodes...)
}

// VerticalLocation 生成任务位置列表，节点在垂直方向上排列
// 如 (100, 100), (100, 300), (100, 500)
func VerticalLocation(taskCodes ...int64) []process.TaskLocation {
	x := 100
	beginY := 100
	yStep := 300
	return VerticalLocationCustom(x, beginY, yStep, taskCodes...)
}

// HorizontalLocationCustom 生成任务位置列表，节点在水平方向上按照自定义参数排列
// 如 (beginX, y), (beginX + step, y) ...
func HorizontalLocationCustom(beginX, y, xStep int, taskCodes ...int64) []process.TaskLocation {
	list := []process.TaskLocation{}
	for _, taskCode := range taskCodes {
		list = append(list, process.TaskLocation{TaskCode: taskCode, X: beginX, Y: y})
		beginX += xStep
	}
	return list
}

// VerticalLocationCustom 生成任务位置列表，节点在垂直方向上按照自定义参数排列
func VerticalLocationCustom(x, beginY, yStep int, taskCodes ...int64) []process.TaskLocation {
	list := []process.TaskLocation{}
	for _, taskCode := range taskCodes {
		list = append(list, process.TaskLocation{TaskCode: taskCode, X: x, Y: beginY})
		beginY += yStep
	}
	return list
}
