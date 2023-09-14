package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PgDB struct {
	db *sql.DB
}

func New(dsn string) (*PgDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &PgDB{db: db}, nil
}

func (d PgDB) CheckConnection() error {
	return d.db.Ping()
}
