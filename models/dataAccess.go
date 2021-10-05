package models

type ITaskDal interface {
	Insert(task Task) Task
}
