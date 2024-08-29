package task

// TaskResource 代表任务资源
// 或许因为我们使用的ambari的DS，存储在HDFS上，导致这里通过ID根本无法获取任何信息
// 应该手动定义点Resource
type TaskResource struct {
	ResourceName string `json:"resourceName"`
}
