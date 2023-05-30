package main

import (
	"balancy/controllers"
	"balancy/database"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db := database.DatabaseInit()
	ac := controllers.BalanceController{Db: db}
	res := controllers.ReservesController{Db: db}

	http.HandleFunc("/payments", ac.Create)
	http.HandleFunc("/balance", ac.Show)
	http.HandleFunc("/reserve", res.Create)
	http.HandleFunc("/unreserve", res.Create)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
