package core

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/core/operator"
	"github.com/PaienNate/dolphinscheduler-sdk-go/remote"
)

// DolphinClient dolphin scheduler client to operate dolphin scheduler
type DolphinClient struct {
	dolphinsRestTemplate *remote.DolphinsRestTemplate
	dolphinAddress       string
	token                string

	//dataSourceOperator      *datasource.DataSourceOperator
	//resourceOperator        *resource.ResourceOperator
	processOperator *operator.ProcessOperator
	//processInstanceOperator *instance.ProcessInstanceOperator
	//scheduleOperator        *schedule.ScheduleOperator
	//projectOperator         *project.ProjectOperator
	//tenantOperator          *tenant.TenantOperator
}

func NewDolphinClient(token, dolphinAddress string, dolphinsRestTemplate *remote.DolphinsRestTemplate) *DolphinClient {
	client := &DolphinClient{
		token:                token,
		dolphinAddress:       dolphinAddress,
		dolphinsRestTemplate: dolphinsRestTemplate,
	}
	client.initOperators()
	return client
}

func (c *DolphinClient) initOperators() {
	// 数据源
	//c.dataSourceOperator = datasource.NewDataSourceOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
	//// 资源中心
	//c.resourceOperator = resource.NewResourceOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
	// 工作流
	c.processOperator = operator.NewProcessOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
	//c.processInstanceOperator = instance.NewProcessInstanceOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
	//c.scheduleOperator = schedule.NewScheduleOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
	//c.projectOperator = project.NewProjectOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
	//c.tenantOperator = tenant.NewTenantOperator(c.dolphinAddress, c.token, c.dolphinsRestTemplate)
}

//func (c *DolphinClient) OpsForDataSource() *datasource.DataSourceOperator {
//	return c.dataSourceOperator
//}
//
//func (c *DolphinClient) OpsForResource() *resource.ResourceOperator {
//	return c.resourceOperator
//}

func (c *DolphinClient) OpsForProcess() *operator.ProcessOperator {
	return c.processOperator
}

//func (c *DolphinClient) OpsForProcessInst() *instance.ProcessInstanceOperator {
//	return c.processInstanceOperator
//}
//
//func (c *DolphinClient) OpsForSchedule() *schedule.ScheduleOperator {
//	return c.scheduleOperator
//}
//
//func (c *DolphinClient) OpsForProject() *project.ProjectOperator {
//	return c.projectOperator
//}
//
//func (c *DolphinClient) OpsForTenant() *tenant.TenantOperator {
//	return c.tenantOperator
//}
