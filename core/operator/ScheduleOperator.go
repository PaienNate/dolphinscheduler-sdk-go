package operator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PaienNate/dolphinscheduler-sdk-go/remote"
	"github.com/PaienNate/dolphinscheduler-sdk-go/schedule"
	"github.com/PaienNate/dolphinscheduler-sdk-go/util"
	"strconv"
)

type ScheduleOperator struct {
	*AbstractOperator
}

func NewScheduleOperator(dolphinAddress, token string, dolphinsRestTemplate *remote.DolphinsRestTemplate) *ScheduleOperator {
	return &ScheduleOperator{
		AbstractOperator: NewAbstractOperator(dolphinAddress, token, dolphinsRestTemplate),
	}
}

// 创建定时
func (p *ScheduleOperator) Create(projectCode int64, scheduleDefineParam *schedule.ScheduleDefineParam) (*schedule.ScheduleInfoResp, error) {
	url := fmt.Sprintf("%s/projects/%d/schedules", p.DolphinAddress, projectCode)
	result, err := p.DolphinsRestClient.PostForm(url, util.ToFormValues(*scheduleDefineParam))
	if err != nil {
		return nil, err
	}
	var scheduleInfoResp schedule.ScheduleInfoResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &scheduleInfoResp)
		if err != nil {
			return nil, err
		}
		return &scheduleInfoResp, nil
	}
	return nil, errors.New("请求不成功")
}

// 根据工作流代码获取定时
func (p *ScheduleOperator) GetByWorkflowCode(projectCode, processDefinitionCode int64) ([]*schedule.ScheduleInfoResp, error) {
	url := fmt.Sprintf("%s/projects/%d/schedules", p.DolphinAddress, projectCode)
	query := make(map[string]string)
	query["pageNo"] = "1"
	query["pageSize"] = "10"
	query["processDefinitionCode"] = strconv.FormatInt(processDefinitionCode, 10)
	result, err := p.DolphinsRestClient.Get(url, query)
	var scheduleInfoResp []*schedule.ScheduleInfoResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &scheduleInfoResp)
		if err != nil {
			return nil, err
		}
		return scheduleInfoResp, nil
	}
	return nil, nil
}

// 更新定时
func (p *ScheduleOperator) Update(projectCode, scheduleId int64, scheduleDefineParam *schedule.ScheduleDefineParam) (*schedule.ScheduleInfoResp, error) {
	url := fmt.Sprintf("%s/projects/%d/schedules/%d", p.DolphinAddress, projectCode, scheduleId)
	result, err := p.DolphinsRestClient.PutForm(url, util.ToFormValues(*scheduleDefineParam))
	if err != nil {
		return nil, err
	}
	var scheduleInfoResp schedule.ScheduleInfoResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &scheduleInfoResp)
		if err != nil {
			return nil, err
		}
		return &scheduleInfoResp, nil
	}
	return nil, nil
}

// 上线定时
// TODO：优化不必要的bool
func (p *ScheduleOperator) Online(projectCode, scheduleId int64) (bool, error) {
	url := fmt.Sprintf("%s/projects/%d/schedules/%d/online", p.DolphinAddress, projectCode, scheduleId)
	result, err := p.DolphinsRestClient.PostForm(url, nil)
	if err != nil {
		return false, err
	}
	var scheduleInfoResp schedule.ScheduleInfoResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &scheduleInfoResp)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 下线定时
func (p *ScheduleOperator) Offline(projectCode, scheduleId int64) (bool, error) {
	url := fmt.Sprintf("%s/projects/%d/schedules/%d/offline", p.DolphinAddress, projectCode, scheduleId)
	result, err := p.DolphinsRestClient.PostForm(url, nil)
	if err != nil {
		return false, err
	}
	var scheduleInfoResp schedule.ScheduleInfoResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &scheduleInfoResp)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 删除定时
func (p *ScheduleOperator) Delete(projectCode, scheduleId int64) (bool, error) {
	url := fmt.Sprintf("%s/projects/%d/schedules/%d", p.DolphinAddress, projectCode, scheduleId)
	result, err := p.DolphinsRestClient.Delete(url, nil)
	var scheduleInfoResp schedule.ScheduleInfoResp
	if result.Success {
		// 获取并转换为对应的ProcessDefineResp
		data := result.Data
		err = json.Unmarshal(data, &scheduleInfoResp)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
