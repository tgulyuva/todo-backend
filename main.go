package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tgulyuva/todo-backend/controllers"
	da "github.com/tgulyuva/todo-backend/utility"
)

func main() {

	r := httprouter.New()
	tc := controllers.NewTodoController((da.NewTaskDataAccess()))
	r.POST("/task/add", tc.AddTask)
	http.ListenAndServe("localhost:9000", r)

}
