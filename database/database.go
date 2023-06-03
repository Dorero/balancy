package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	dbOnce sync.Once
	conn   *pgx.Conn
)

func DBInstance() *pgx.Conn {

	err := loadEnvFromFile(".env")

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

func loadEnvFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("Неверный формат строки:", line)
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return nil
}
