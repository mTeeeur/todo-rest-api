package repository

import (
	"github.com/jmoiron/sqlx"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
)

type Authorization interface {
	CreateUser(user todo_rest_api.User) (int, error)
	GetUser(username, password string) (todo_rest_api.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
