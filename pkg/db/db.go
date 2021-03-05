package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(config interface{}) (db *sql.DB, err error) {
	//TODO use config
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "abc"
		dbname   = "nooble_task"
	)
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	return sql.Open("postgres", url)
}
