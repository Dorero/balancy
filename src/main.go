package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=postgres password=postgres dbname=balancy sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connect correct done")

	result, err := db.Exec("SELECT * FROM accounts")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
