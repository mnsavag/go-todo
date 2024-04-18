package sqlite

import (
	"database/sql"
	"fmt"

	"goTodo/internal/storage"

	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

type Storage struct {
	db *sql.DB
}

func NewSqlite(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.NewSqlite"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	initStmts := storage.GetInitDBQueries()
	for _, query := range initStmts {
		stmt, err := db.Prepare(query)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		_, err = stmt.Exec()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return &Storage{db: db}, nil
}
