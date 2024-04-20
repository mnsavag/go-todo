package sqlite

import (
	"database/sql"
	"fmt"

	"goTodo/internal/model"
	cmnStorage "goTodo/internal/repository/cmn-storage"
)

type AuthSqlite struct {
	db *sql.DB
}

func NewAuthSqlite(db *sql.DB) *AuthSqlite {
	return &AuthSqlite{db: db}
}

func (s *AuthSqlite) CreateUser(user model.User) (int64, error) {
	const op = "repository.sqlite.CreateUser"

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES (?, ?, ?)", cmnStorage.UsersTable)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(user.Name, user.Username, user.Password)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}
