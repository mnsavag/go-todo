package sqlite

import (
	"database/sql"
	"fmt"

	"goTodo/internal/model"
	cmnStorage "goTodo/internal/repository/cmn-storage"
)

type TodoItemSqlite struct {
	db *sql.DB
}

func NewTodoItemSqlite(db *sql.DB) *TodoItemSqlite {
	return &TodoItemSqlite{db: db}
}

func (r *TodoItemSqlite) Create(listId int64, item model.TodoItem) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int64
	createItemQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ('%s', '%s')",
		cmnStorage.TodoItemsTable, item.Title, item.Description,
	)

	resItemQuery, err := r.db.Exec(createItemQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	itemId, err = resItemQuery.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf(
		"INSERT INTO %s (list_id, item_id) VALUES ('%x', '%x')",
		cmnStorage.ListsItemsTable, listId, itemId,
	)

	_, err = r.db.Exec(createListItemsQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}
