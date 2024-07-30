package operator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
	"github.com/PaienNate/dolphinscheduler-sdk-go/remote"
	"github.com/PaienNate/dolphinscheduler-sdk-go/util"
	"strconv"
)

// TODO：ERROR判断整体有误，需要修改

// 添加上默认的AbstractOperator，然后使用它初始化
// ProcessOperator 用于操作Dolphin Scheduler流程定义的操作类
type ProcessOperator struct {
	*AbstractOperator
}

// 暂时去掉引用

// NewProcessOperator 创建一个新的 ProcessOperator 实例
// 这个需要和之前的代码对接
func NewProcessOperator(dolphinAddress, token string, dolphinsRestTemplate *remote.DolphinsRestTemplate) *ProcessOperator {
	return &ProcessOperator{
		AbstractOperator: NewAbstractOperator(dolphinAddress, token, dolphinsRestTemplate),
	}
}

// Page 分页查询流程定义
//
// @param projectCode 项目代码
// @param page 页码
// @param size 每页大小
// @param searchVal 流程名称
// @return 流程定义响应列表
func (p *ProcessOperator) Page(projectCode int64, page, size *int, searchVal *string) ([]process.ProcessDefineResp, error) {
	url := fmt.Sprintf("%s/projects/%d/process-definition", p.DolphinAddress, projectCode)
	// TODO:输出日志

	// 如果page和size为空，那么默认page为1，默认size为10
	defaultPage := 1
	defaultSize := 10

	if page == nil {
		page = &defaultPage
	}
	if size == nil {
		size = &defaultSize
	}
	// 发送Query map[string][string]添加数据
	query := make(map[string]string)
	query["pageNo"] = strconv.Itoa(*page)
	query["pageSize"] = strconv.Itoa(*size)
	query["searchVal"] = *searchVal
	// 调用发送部分，发送对应信息
	var processDefineRespList []process.ProcessDefineResp
	result, err := p.DolphinsRestClient.Get(url, query)
	if err != nil {
		return nil, err
	}
	if result.Success {
		// 获取并转换为对应的[]ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &processDefineRespList)
		if err != nil {
			return nil, err
		}
		return processDefineRespList, nil
	}

	return nil, errors.New("请求不成功")
}

// Create 创建流程定义
// API: /dolphinscheduler/projects/{projectCode}/process-definition
//
// @param projectCode 项目代码
// @param processDefineParam 创建流程的参数
// @return 创建响应
func (p *ProcessOperator) Create(projectCode int64, processDefineParam *process.ProcessDefineParam) (*process.ProcessDefineResp, error) {
	// 先定义URL
	url := fmt.Sprintf("%s/projects/%d/process-definition", p.DolphinAddress, projectCode)
	// TODO:指定结构体FORM字段
	result, err := p.DolphinsRestClient.PostForm(url, util.ToFormValues(*processDefineParam))
	if err != nil {
		return nil, err
	}
	var SingleProcessDefineResp process.ProcessDefineResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &SingleProcessDefineResp)
		if err != nil {
			return nil, err
		}
		return &SingleProcessDefineResp, nil
	}
	return nil, errors.New("请求不成功")
}

// Update 更新流程定义
// API: /dolphinscheduler/projects/{projectCode}/process-definition/{process-definition-code}
//
// @param projectCode 项目代码
// @param processDefineParam 更新流程的参数
// @param processCode 流程代码
// @return 更新响应
func (p *ProcessOperator) Update(projectCode int64, processDefineParam *process.ProcessDefineParam, processCode int64) (*process.ProcessDefineResp, error) {
	// 定义URL
	url := fmt.Sprintf("%s/projects/%d/process-definition/%s", p.DolphinAddress, projectCode, processCode)
	// 执行代码
	// TODO:指定结构体FORM字段
	result, err := p.DolphinsRestClient.PutForm(url, util.ToFormValues(processDefineParam))
	// 返回是否成功
	var SingleProcessDefineResp process.ProcessDefineResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &SingleProcessDefineResp)
		if err != nil {
			return nil, err
		}
		return &SingleProcessDefineResp, nil
	}
	return nil, errors.New("请求不成功")
}

// Delete 删除流程定义
//
// @param projectCode 项目代码
// @param processCode 流程代码
// @return 成功则返回true，否则返回false
func (p *ProcessOperator) Delete(projectCode int64, processCode int64) (bool, error) {
	// 实现代码
	url := fmt.Sprintf("%s/projects/%d/process-definition/%s", p.DolphinAddress, projectCode, processCode)
	result, err := p.DolphinsRestClient.Delete(url, nil)
	// 返回是否成功
	if result.Success {
		return true, nil
	}
	return false, err
}

// Release 发布流程定义
// API: /dolphinscheduler/projects/{projectCode}/process-definition/{code}/release
//
// @param projectCode 项目代码
// @param code 流程ID
// @param processReleaseParam 发布参数
// @return 成功则返回true，否则返回false
func (p *ProcessOperator) Release(projectCode int64, code int64, processReleaseParam *process.ProcessReleaseParam) (bool, error) {
	// 实现代码
	url := fmt.Sprintf("%s/projects/%d/process-definition/%s/release", p.DolphinAddress, projectCode, code)
	result, err := p.DolphinsRestClient.PostForm(url, util.ToFormValues(processReleaseParam))
	if result.Success {
		return true, nil
	}
	return false, err
}

// Online 上线流程定义
//
// @param projectCode 项目代码
// @param code 流程ID
// @return 成功则返回true，否则返回false
func (p *ProcessOperator) Online(projectCode int64, code int64) (bool, error) {
	return p.Release(projectCode, code, nil)
}

// Offline 下线流程定义，替代 Release 方法
//
// @param projectCode 项目代码
// @param code 流程ID
// @return 成功则返回true，否则返回false
func (p *ProcessOperator) Offline(projectCode int64, code int64) (bool, error) {
	return p.Release(projectCode, code, nil)
}

// GenerateTaskCode 生成任务唯一标识代码
//
// @param projectCode 项目代码
// @param codeNumber 任务代码数量
// @return 任务代码列表
func (p *ProcessOperator) GenerateTaskCode(projectCode int64, codeNumber int) ([]int64, error) {
	// 实现代码
	url := fmt.Sprintf("%s/projects/%d/task-definition/gen-task-codes", p.DolphinAddress, projectCode)
	query := make(map[string]string)
	query["genNum"] = strconv.Itoa(codeNumber)
	result, err := p.DolphinsRestClient.Get(url, query)
	var trulyResponse []int64
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &trulyResponse)
		return trulyResponse, nil
	}
	return nil, err
}
