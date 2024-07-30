package processutils

import (
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
	"github.com/PaienNate/dolphinscheduler-sdk-go/task"
	"strconv"
	"time"
)

// TaskDefinitionUtils 的常量定义
const (
	FLAG_YES = "YES" // 节点将被执行
	FLAG_NO  = "NO"  // 节点不会被执行

	PRIORITY_HIGHEST = "HIGHEST"
	PRIORITY_HIGH    = "HIGH"
	PRIORITY_MEDIUM  = "MEDIUM"
	PRIORITY_LOW     = "LOW"
	PRIORITY_LOWEST  = "LOWEST"

	IS_CACHE_YES = "YES" // 执行缓存
	IS_CACHE_NO  = "NO"  // 不执行缓存
)

// CreateDefaultTaskDefinition 创建具有默认配置的任务定义，可以满足基本需求
func CreateDefaultTaskDefinition(taskCode int64, task task.Task) *process.TaskDefinition {
	return CreateTaskDefinition("", taskCode, 0, task, FLAG_YES, PRIORITY_MEDIUM, nil, nil)
}

// CreateDefaultTaskDefinitionWithName 创建具有默认配置的任务定义，可以满足基本需求
func CreateDefaultTaskDefinitionWithName(taskName string, taskCode int64, task task.Task) *process.TaskDefinition {
	return CreateTaskDefinition(taskName, taskCode, 0, task, FLAG_YES, PRIORITY_MEDIUM, nil, nil)
}

// CreateDefaultTaskDefinitionWithVersion 创建具有默认配置的任务定义，可以满足基本需求
func CreateDefaultTaskDefinitionWithVersion(taskCode int64, version int, task task.Task) *process.TaskDefinition {
	return CreateTaskDefinition("", taskCode, version, task, FLAG_YES, PRIORITY_MEDIUM, nil, nil)
}

// CreateBannedTaskDefinition 创建不会执行的任务
func CreateBannedTaskDefinition(taskCode int64, version int, task task.Task) *process.TaskDefinition {
	return CreateTaskDefinition("", taskCode, version, task, FLAG_NO, PRIORITY_MEDIUM, nil, nil)
}

// CreateHighLevelTaskDefinition 创建高优先级任务
func CreateHighLevelTaskDefinition(taskCode int64, version int, task task.Task) *process.TaskDefinition {
	return CreateTaskDefinition("", taskCode, version, task, FLAG_YES, PRIORITY_HIGH, nil, nil)
}

// CreateTaskDefinition 创建任务定义
func CreateTaskDefinition(
	taskName string,
	taskCode int64,
	version int,
	task task.Task,
	flag string,
	taskPriority string,
	cpuQuota *int,
	memoryMax *int64,
) *process.TaskDefinition {
	if taskName == "" {
		taskName = task.GetTaskType() + strconv.FormatInt(time.Now().UnixMilli(), 10)
	}

	taskDefinition := &process.TaskDefinition{
		Code:                  taskCode,
		Version:               version,
		Name:                  taskName,
		Description:           "",
		TaskType:              task.GetTaskType(),
		TaskParams:            task,
		Flag:                  flag,
		TaskPriority:          taskPriority,
		WorkerGroup:           "default",
		FailRetryTimes:        "0",
		FailRetryInterval:     "1",
		TimeoutFlag:           "CLOSE",
		TimeoutNotifyStrategy: "WARN",
		IsCache:               IS_CACHE_NO,
	}

	if cpuQuota != nil {
		taskDefinition.CpuQuota = *cpuQuota
	}
	if memoryMax != nil {
		taskDefinition.MemoryMax = *memoryMax
	}

	return taskDefinition
}
