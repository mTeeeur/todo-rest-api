package repository

import (
	"github.com/jmoiron/sqlx"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (s *AuthPostgres) CreateUser(user todo_rest_api.User) (int, error) {
	return 0, int
}
