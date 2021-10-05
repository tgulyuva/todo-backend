package utility

import (
	"github.com/tgulyuva/todo-backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskDataAccess struct {
}

func NewTaskDataAccess() *TaskDataAccess {
	return &TaskDataAccess{}
}

func (t TaskDataAccess) Insert(task models.Task) models.Task {
	task.Id = bson.NewObjectId()
	getSession().DB("todo-mongo").C("tasks").Insert(task)
	return task
}
func getSession() *mgo.Session {
	session, error := mgo.Dial("mongodb://localhost:2717")

	if error != nil {
		panic(error)
	}
	return session
}
