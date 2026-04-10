package db

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func Connect() *sql.DB {
	once.Do(func() {
		connStr := "postgres://postgres:password@db:5432/benchlab?sslmode=disable"
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		}
	})
	return db
}
