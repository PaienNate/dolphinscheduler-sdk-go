package processutils

import "github.com/PaienNate/dolphinscheduler-sdk-go/process"

// OneLineRelation 该方法用于构建单行任务关系，例如 t1->t2->t3->t4......
func OneLineRelation(taskCodes ...int64) []process.TaskRelation {
	var list []process.TaskRelation
	for i, taskCode := range taskCodes {
		if i == 0 {
			list = append(list, process.TaskRelation{PostTaskCode: taskCode})
		} else {
			list = append(list, process.TaskRelation{
				PreTaskCode:  taskCodes[i-1],
				PostTaskCode: taskCode,
			})
		}
	}
	return list
}

// NoRelation 该方法用于构建没有关系的任务集合，例如 t1 t2 t3
func NoRelation(taskCodes ...int64) []process.TaskRelation {
	var list []process.TaskRelation
	for _, taskCode := range taskCodes {
		list = append(list, process.TaskRelation{PostTaskCode: taskCode})
	}
	return list
}
