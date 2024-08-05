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
	codeList, err := client.OpsForProcess().GenerateTaskCode(13917543937344, 2)
	codeList2, err := client.OpsForProcess().GenerateTaskCode(13917543937344, 1)
	if err != nil {
		return
	}
	// 获取对应的节点code
	//code := codeList[0]
	// 创建一个新任务
	shelltask := task.NewShellTask()
	shelltask.RawScript = "echo 'hello dolphin scheduler java sdk'"
	shelltask2 := task.NewShellTask()
	shelltask2.RawScript = "echo 'hello dolphin scheduler java sdk'"
	shelltask3 := task.NewShellTask()
	shelltask3.RawScript = "echo 'hello dolphin scheduler java sdk'"
	// 创建第一组任务并装两个
	var test = make([]process.TaskDefinition, 2)
	definition := processutils.CreateDefaultTaskDefinition(codeList[0], shelltask)
	test[0] = *definition
	// 获取任务定义
	definition2 := processutils.CreateDefaultTaskDefinition(codeList[1], shelltask)
	test[1] = *definition2

	var test2 = make([]process.TaskDefinition, 1)
	// 获取任务定义
	definition3 := processutils.CreateDefaultTaskDefinition(codeList2[0], shelltask)
	test2[0] = *definition3

	merged := append(test, test2...)
	merged2 := append(codeList, codeList2...)

	// 创建对应的任务信息

	var param = process.ProcessDefineParam{
		Name: "TESTTTQWQ",
		// 使用展开运算符传递参数
		Locations: processutils.VerticalLocation(merged2...),
		// 使用合并的任务信息
		TaskDefinitionJson: merged,
		TaskRelationJson:   processutils.MultiLineRelation(codeList, codeList2),
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
