// internal/infrastructure/database/mysql.go
package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OrquideaDB() (*sql.DB, error) {

	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_SERVER := os.Getenv("DB_SERVER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_URI := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_SERVER + ":" + DB_PORT + ")/" + DB_NAME
	db, err := sql.Open("mysql", DB_URI)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
