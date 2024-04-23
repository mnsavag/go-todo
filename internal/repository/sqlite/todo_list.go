package sqlite

import (
	"database/sql"
	"fmt"
	"goTodo/internal/model"
	cmnStorage "goTodo/internal/repository/cmn-storage"
)

type TodoListSqlite struct {
	db *sql.DB
}

func NewTodoListSqlite(db *sql.DB) *TodoListSqlite {
	return &TodoListSqlite{db: db}
}

func (r *TodoListSqlite) Create(userId int64, list model.TodoList) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ('%s', '%s')",
		cmnStorage.TodoListsTable, list.Title, list.Description,
	)
	resListQuery, err := r.db.Exec(createListQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	listId, err := resListQuery.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, list_id) VALUES ('%x', '%x')",
		cmnStorage.UsersListsTable, userId, listId,
	)
	_, err = r.db.Exec(createUsersListQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}
