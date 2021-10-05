package mocks

import (
	"github.com/tgulyuva/todo-backend/models"
)

type MockTask struct {
}

func (t MockTask) Insert(task models.Task) models.Task {
	task.Id = "1111"
	return task
}
