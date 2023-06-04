package database

import (
	"balancy/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"sync"
)

var (
	dbOnce sync.Once
	conn   *pgx.Conn
)

func DBInstance() *pgx.Conn {

	err := utils.LoadEnvFromFile(".env")

	if err != nil {
		log.Fatal("Ошибка при загрузке файла .env:", err)
	}

	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	dbOnce.Do(func() {
		connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", username, password, dbName)
		instance, err := pgx.Connect(context.Background(), connStr)

		if err != nil {
			log.Fatal(err)
		}

		conn = instance
	})
	return conn
}
