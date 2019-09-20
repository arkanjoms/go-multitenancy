package repository

import (
	"context"
	"go-multitenancy/model"
	"go-multitenancy/multitenancy"
	"go-multitenancy/util"
)

type TodoRepository struct{}

func (t TodoRepository) GetAll(ctx context.Context, todos []model.Todo) (interface{}, error) {
	db := multitenancy.GetDatasource(ctx)
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return []model.Todo{}, err
	}

	todo := model.Todo{}
	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Description, &todo.Completed)
		todos = append(todos, todo)
	}

	return util.ResultData(todos, []model.Todo{}, err)
}

func (t TodoRepository) AddTodo(ctx context.Context, todo model.Todo) (interface{}, error) {
	db := multitenancy.GetDatasource(ctx)
	err := db.QueryRow("INSERT INTO todo (description,completed) VALUES ($1,false) RETURNING id;",
		todo.Description,
	).Scan(&todo.ID)

	return util.ResultData(todo.ID, 0, err)
}
