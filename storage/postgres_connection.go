package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	Connection, err := sql.Open("postgres", "postgres://postgres:@localhost:5432/go_basics_dev")
	if err != nil {
		return nil, err
	}
	if err = Connection.Ping(); err != nil {
		return nil, err
	}
	return Connection, nil
}