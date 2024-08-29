package task

// TaskResource 代表任务资源
// 或许因为我们使用的ambari的DS，存储在HDFS上，导致这里通过ID根本无法获取任何信息
// 应该手动定义点Resource
type TaskResource []Resource

type Resource struct {
	ResourceName string `json:"resourceName"`
}

// NewTaskResource 是一个构造函数，用于通过传入路径字符串生成 TaskResource
// 路径字符串建议通过Resource相关函数去获得。
func NewTaskResource(paths ...string) TaskResource {
	var resources TaskResource
	for _, path := range paths {
		resources = append(resources, Resource{
			ResourceName: path,
		})
	}
	return resources
}
