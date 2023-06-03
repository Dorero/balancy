package controllers

import (
	"balancy/services"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type PaymentsController struct {
	Db *sql.DB
}

func (c PaymentsController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, errId := strconv.Atoi(r.URL.Query().Get("user_id"))
	amount, errAmount := strconv.Atoi(r.URL.Query().Get("amount"))

	if errId != nil {
		log.Println("Ошибка преобразования строки в число:", errId)
		return
	}

	if errId != nil {
		log.Println("Ошибка преобразования строки в число:", errAmount)
		return
	}

	payment := services.PaymentService{}.Create(userId, amount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(payment)
}
