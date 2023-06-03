package controllers

import (
	"balancy/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type BalanceController struct {
}

func (c BalanceController) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		return
	}

	balance := services.BalanceService{}.Show(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(balance)
}
