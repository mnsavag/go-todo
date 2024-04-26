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

func (r *TodoListSqlite) GetAll(userId int64) ([]model.TodoList, error) {
	var lists []model.TodoList
	var list model.TodoList
	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id=ul.list_id WHERE ul.user_id = %x",
		cmnStorage.TodoListsTable, cmnStorage.UsersListsTable, userId,
	)

	rows, err := r.db.Query(query)
	if err != nil {
		return lists, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&list.Id, &list.Title, &list.Description)
		if err != nil {
			return lists, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *TodoListSqlite) GetById(userId, listId int64) (model.TodoList, error) {
	var list model.TodoList
	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl 
		INNER JOIN %s ul ON tl.id=ul.list_id WHERE ul.user_id = ? AND ul.list_id = ?`,
		cmnStorage.TodoListsTable, cmnStorage.UsersListsTable,
	)

	err := r.db.QueryRow(query, userId, listId).Scan(&list.Id, &list.Title, &list.Description)
	if err != nil && err != sql.ErrNoRows {
		return list, err
	}

	return list, nil
}
