package main

import (
	"fmt"
	"github.com/PaienNate/dolphinscheduler-sdk-go/core"
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
	"github.com/PaienNate/dolphinscheduler-sdk-go/process/processutils"
	"github.com/PaienNate/dolphinscheduler-sdk-go/remote"
	"github.com/PaienNate/dolphinscheduler-sdk-go/task"
)

// TODO：提取所有的Model到同一层，以避免循环依赖（Go）问题

func main() {
	// 构建Header
	var BaseURL = "http://127.0.0.1:12345/dolphinscheduler"
	var Header = map[string]string{"token": "7f7b599651ed35c0872b29b6b3ef9649"}
	rest := remote.NewDolphinsRestTemplate(Header)
	client := core.NewDolphinClient("7f7b599651ed35c0872b29b6b3ef9649", BaseURL, rest)
	codeList, err := client.OpsForProcess().GenerateTaskCode(13917543937344, 1)
	if err != nil {
		return
	}
	// 获取对应的节点code
	code := codeList[0]
	// 创建一个新任务
	shelltask := task.NewShellTask()
	shelltask.RawScript = "echo 'hello dolphin scheduler java sdk'"
	var test = make([]process.TaskDefinition, 1)
	// 获取任务定义
	definition := processutils.CreateDefaultTaskDefinition(code, shelltask)
	test[0] = *definition
	// 创建对应的任务信息

	var param = process.ProcessDefineParam{
		Name:               "TESTTT",
		Locations:          processutils.VerticalLocation(code),
		TaskDefinitionJson: test,
		TaskRelationJson:   processutils.OneLineRelation(code),
		// 这个租户code回头再说
		//TenantCode:  "1",
		Description: "测试描述",
		// 并行执行
		ExecutionType: process.EXECUTION_TYPE_PARALLEL,
		GlobalParams:  nil,
		Timeout:       "0",
	}
	create, err := client.OpsForProcess().Create(13917543937344, &param)
	if err != nil {
		return
	}
	fmt.Println(create)

}
