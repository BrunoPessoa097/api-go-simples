package database

import (
	"database/sql"
	"log"
)

func Conectar() *sql.DB {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
