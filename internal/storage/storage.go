package storage

import "errors"

// general for all storage implementetion

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

func GetInitDBQueries() []string {
	return []string{
		`CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name varchar(255) NOT NULL,
			username varchar(255) NOT NULL UNIQUE,
			password_hash varchar(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS todo_list(
			id SERIAL PRIMARY KEY,
			title varchar(255) NOT NULL,
			description varchar(255)
		);`,
		`CREATE TABLE IF NOT EXISTS users_lists(
			id SERIAL PRIMARY KEY,
			user_id int NOT NULL,
			list_id int NOT NULL,

			FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
			FOREIGN KEY (list_id) REFERENCES todo_lists ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS todo_item(
			id SERIAL PRIMARY KEY,
			title varchar(255) NOT NULL,
			description varchar(255),
			done boolean NOT NULL DEFAULT false
		);`,
		`CREATE TABLE IF NOT EXISTS list_items(
			id SERIAL PRIMARY KEY,
			item_id int NOT NULL,
			list_id int NOT NULL,

			FOREIGN KEY (item_id) REFERENCES todo_items ON DELETE CASCADE,
			FOREIGN KEY (list_id) REFERENCES todo_lists ON DELETE CASCADE
		);`,
	}
}
