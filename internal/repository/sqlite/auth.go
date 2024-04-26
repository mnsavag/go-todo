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

func (r *AuthSqlite) CreateUser(user model.User) (int64, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) VALUES ('%s', '%s', '%s')",
		cmnStorage.UsersTable, user.Name, user.Username, user.Password,
	)

	res, err := r.db.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (r *AuthSqlite) GetRegisteredUser(username, password string) (model.User, error) {
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE username='%s' AND password_hash='%s'",
		cmnStorage.UsersTable, username, password,
	)
	var user model.User

	err := r.db.QueryRow(query).Scan(&user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}
