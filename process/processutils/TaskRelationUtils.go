package processutils

import (
	"errors"
	"github.com/PaienNate/dolphinscheduler-sdk-go/process"
)

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

// GetJustARelation 只获取两个taskCodes的联系。返回一个TaskRelation对象。
// 其主要用于DAG中按顺序获取的情况
func GetJustARelation(taskCodes ...int64) (error, *process.TaskRelation) {
	if len(taskCodes) != 2 {
		return errors.New("获取单个TaskRelation的参数有误"), nil
	}
	var node *process.TaskRelation
	for i, taskCode := range taskCodes {
		// 跳过第一个node参数，直接处理第二个
		if i == 0 {
			continue
		}
		node = &process.TaskRelation{
			PreTaskCode:  taskCodes[i-1],
			PostTaskCode: taskCode,
		}
	}
	return nil, node
}

// 该方法是GPT说用于构建多行任务关系，例如(1,2,3)->(4,5)->(6,7)
func MultiLineRelation(taskCodeLists ...[]int64) []process.TaskRelation {
	var list []process.TaskRelation
	var previousTaskCodes []int64

	for _, taskCodes := range taskCodeLists {
		var currentList []process.TaskRelation
		for _, taskCode := range taskCodes {
			if len(previousTaskCodes) == 0 {
				// 第一个列表的任务
				currentList = append(currentList, process.TaskRelation{
					PreTaskCode:  0,
					PostTaskCode: taskCode,
				})
			} else {
				// 之后列表的任务，每个任务的 PreTaskCode 都是前一个列表中的所有任务
				for _, prevTaskCode := range previousTaskCodes {
					currentList = append(currentList, process.TaskRelation{
						PreTaskCode:  prevTaskCode,
						PostTaskCode: taskCode,
					})
				}
			}
		}
		// 更新前一个任务列表为当前任务列表
		previousTaskCodes = taskCodes
		list = append(list, currentList...)
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
