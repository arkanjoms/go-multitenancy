package controller

import (
	"encoding/json"
	"errors"
	"go-multitenancy/model"
	"go-multitenancy/multitenancy"
	"go-multitenancy/repository"
	"go-multitenancy/util"
	"net/http"
)

var (
	todos          []model.Todo
	todoRepository repository.TodoRepository
)

type TodoController struct{}

func (t TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	_ = json.NewDecoder(r.Body).Decode(&todo)

	if todo.Description == "" {
		util.SendBadRequest(w, errors.New("enter missing fields"))
		return
	}

	data, err := todoRepository.AddTodo(multitenancy.TenantResolver(r), todo)
	util.SendResult(w, data, err)
}

func (t TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := todoRepository.GetAll(multitenancy.TenantResolver(r), todos)
	util.SendResult(w, data, err)
}
