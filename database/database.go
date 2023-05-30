package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DatabaseInit() *sql.DB {
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s, password=%s, dbname=%s, sslmode=diable", username, password, dbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
