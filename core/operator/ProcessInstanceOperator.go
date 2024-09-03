package operator

// 如果放在别处，似乎会因为循环依赖出现问题
// 所以全部放在operator内

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PaienNate/dolphinscheduler-sdk-go/instance"
	"github.com/PaienNate/dolphinscheduler-sdk-go/remote"
	"github.com/PaienNate/dolphinscheduler-sdk-go/util"
	"strconv"
)

// ProcessInstanceOperator 用于操作工作流实例的结构体
type ProcessInstanceOperator struct {
	*AbstractOperator
}

// NewProcessInstanceOperator 创建一个新的 ProcessInstanceOperator 实例
func NewProcessInstanceOperator(dolphinAddress, token string, dolphinsRestTemplate *remote.DolphinsRestTemplate) *ProcessInstanceOperator {
	return &ProcessInstanceOperator{
		AbstractOperator: NewAbstractOperator(dolphinAddress, token, dolphinsRestTemplate),
	}
}

// Start 启动/运行工作流实例
// API: /dolphinscheduler/projects/{projectCode}/executors/start-process-instance
func (p *ProcessInstanceOperator) Start(projectCode int64, processInstanceCreateParam *instance.ProcessInstanceCreateParam) (bool, error) {
	// 执行这个工作流
	url := fmt.Sprintf("%s/projects/%d/executors/start-process-instance", p.DolphinAddress, projectCode)
	result, err := p.DolphinsRestClient.PostForm(url, util.ToFormValues(processInstanceCreateParam))
	if err != nil {
		return false, err
	}
	if result.Success {
		return true, nil
	}
	return false, errors.New("执行请求失败")
}

// Page 分页查询工作流实例列表
func (p *ProcessInstanceOperator) Page(page, size int, projectCode, workflowCode int64) ([]instance.ProcessInstanceQueryResp, error) {
	url := fmt.Sprintf("%s/projects/%d/process-instances", p.DolphinAddress, projectCode)
	// TODO：抽离通用Page方法
	// 如果page和size为空，那么默认page为1，默认size为10
	defaultPage := 1
	defaultSize := 10
	if page == 0 {
		page = defaultPage
	}
	if size == 0 {
		size = defaultSize
	}
	// 发送Query map[string][string]添加数据
	query := make(map[string]string)
	query["pageNo"] = strconv.Itoa(page)
	query["pageSize"] = strconv.Itoa(size)
	// 转换为十进制字符串
	query["processDefineCode"] = strconv.FormatInt(workflowCode, 10)
	// 请求并获取返回值
	result, err := p.DolphinsRestClient.Get(url, query)
	if err != nil {
		return nil, err
	}
	var processInstanceQueryResp []instance.ProcessInstanceQueryResp
	if result.Success {
		// 获取并转换为对应的[]ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &processInstanceQueryResp)
		if err != nil {
			return nil, err
		}
		return processInstanceQueryResp, nil
	}

	return nil, errors.New("page获取失败")
}

// ReRun 重新运行工作流实例
func (p *ProcessInstanceOperator) ReRun(projectCode, processInstanceId int64) (bool, error) {
	return p.Execute(projectCode, processInstanceId, util.DefaultExecuteType.ReRun)
}

// Stop 停止工作流实例
func (p *ProcessInstanceOperator) Stop(projectCode, processInstanceId int64) (bool, error) {
	// 实现函数逻辑
	return p.Execute(projectCode, processInstanceId, util.DefaultExecuteType.Stop)
}

// Pause 暂停工作流实例
func (p *ProcessInstanceOperator) Pause(projectCode, processInstanceId int64) (bool, error) {
	return p.Execute(projectCode, processInstanceId, util.DefaultExecuteType.Pause)
}

// Execute 执行工作流实例，自定义执行类型
func (p *ProcessInstanceOperator) Execute(projectCode, processInstanceId int64, executeType string) (bool, error) {
	// 实现函数逻辑
	url := fmt.Sprintf("%s/projects/%d/executors/execute", p.DolphinAddress, projectCode)
	param := instance.ProcessInstanceRunParam{
		ProcessInstanceId: processInstanceId,
		ExecuteType:       executeType,
	}
	result, err := p.DolphinsRestClient.PostForm(url, util.ToFormValues(param))
	if err != nil {
		return false, err
	}
	if result.Success {
		return true, nil
	}
	return false, errors.New("execute执行错误")
}

// Delete 删除工作流实例
func (p *ProcessInstanceOperator) Delete(projectCode, processInstanceId int64) (bool, error) {
	url := fmt.Sprintf("%s/projects/%d/process-instances/%d", p.DolphinAddress, projectCode, processInstanceId)
	result, err := p.DolphinsRestClient.Delete(url, nil)
	if err != nil {
		return false, err
	}
	if result.Success {
		return true, nil
	}
	return false, errors.New("delete执行错误")
}
