package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tgulyuva/todo-backend/helpers"
	"github.com/tgulyuva/todo-backend/models"
)

type TodoController struct {
	taskDataAcces models.ITaskDal
}

func NewTodoController(taskDataAcces models.ITaskDal) *TodoController {
	return &TodoController{taskDataAcces: taskDataAcces}
}

func (tc TodoController) AddTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	task := models.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	savedTask := tc.SaveTask(task)
	response := tc.CreateResponseByTask(savedTask)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (controller TodoController) SaveTask(task models.Task) models.Task {

	if helpers.CheckDoubleSpace(task.Name) {
		task.Name = helpers.RemoveDoubleSpace(task.Name)
	}
	return controller.taskDataAcces.Insert(task)
}

func (controller TodoController) CreateResponseByTask(task models.Task) models.ResponseModel {
	var response models.ResponseModel

	if task.Id != "0" {
		response.Success = true
	} else {
		response.Success = false
	}
	return response
}
