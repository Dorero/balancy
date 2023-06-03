package main

import (
	"balancy/controllers"
	"log"
	"net/http"
)

func main() {
	p := controllers.PaymentsController{}
	b := controllers.BalanceController{}
	res := controllers.ReservesController{}

	http.HandleFunc("/payments", p.Create)
	http.HandleFunc("/balance", b.Show)
	http.HandleFunc("/reserve/{id}/", res.Create)
	http.HandleFunc("/unreserve", res.Create)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
