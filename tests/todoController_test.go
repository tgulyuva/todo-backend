package tests

import (
	"testing"

	"github.com/tgulyuva/todo-backend/controllers"
	"github.com/tgulyuva/todo-backend/mocks"
	"github.com/tgulyuva/todo-backend/models"
)

func SuccesCreateTask(t *testing.T) {

	mockTask := mocks.MockTask{}

	request := models.Task{
		Id:   "1111",
		Name: "task testing",
	}

	expectedResponse := models.ResponseModel{
		Success: true,
	}

	response := controllers.NewTodoController(mockTask).CreateResponseByTask(request)

	if response != expectedResponse {
		t.Fail()
	}
}

func FailCreateTask(t *testing.T) {

	mockTask := mocks.MockTask{}
	request := models.Task{
		Id:   "0",
		Name: "task testing",
	}

	expectedResponse := models.ResponseModel{
		Success: false,
	}
	response := controllers.NewTodoController(mockTask).CreateResponseByTask(request)

	if response != expectedResponse {
		t.Fail()
	}
}
func AddTaskWithDoubleSpace(t *testing.T) {
	mockTask := mocks.MockTask{}
	request := models.Task{
		Id:   "0",
		Name: "task double  space",
	}

	expectedResponse := models.ResponseModel{
		Success: true,
	}
	response := controllers.NewTodoController(mockTask).CreateResponseByTask(request)

	if response != expectedResponse {
		t.Fail()
	}
}
