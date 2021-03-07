package db

import (
	"database/sql"
	"fmt"
	"github.com/nooble/task/audio-short-api/pkg/config"

	_ "github.com/lib/pq"
)

func NewDB(config *config.Config) (db *sql.DB, err error) {
	var (
		host     = config.Postgres.Host
		port     = config.Postgres.Port
		user     = config.Postgres.Username
		password = config.Postgres.Password
		dbname   = config.Postgres.Database
	)
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	return sql.Open("postgres", url)
}
