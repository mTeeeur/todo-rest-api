package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo_rest_api.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	listQuery := fmt.Sprintf("INSERT INTO %s (titile, description) VALUES  ($1, $2) RETURNING id", todoLists)
	row := tx.QueryRow(listQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	userListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES  ($1, $2) RETURNING id", usersLists)
	_, err = tx.Exec(userListQuery, userId, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
