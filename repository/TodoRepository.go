package repository

import (
	"database/sql"
	"go-multitenancy/model"
	"go-multitenancy/util"
)

type TodoRepository struct{}

func (t TodoRepository) GetAll(db *sql.DB, todos []model.Todo) (interface{}, error) {
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

func (t TodoRepository) AddTodo(db *sql.DB, todo model.Todo) (interface{}, error) {
	err := db.QueryRow("INSERT INTO todo (description,completed) VALUES ($1,false) RETURNING id;",
		todo.Description,
	).Scan(&todo.ID)

	return util.ResultData(todo.ID, 0, err)
}
