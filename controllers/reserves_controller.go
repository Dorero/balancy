package controllers

import (
	"balancy/services"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ReservesController struct {
	Db *sql.DB
}

func (c ReservesController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, errId := strconv.Atoi(r.URL.Query().Get("user_id"))
	orderId, errId := strconv.Atoi(r.URL.Query().Get("order_id"))
	serviceId, errId := strconv.Atoi(r.URL.Query().Get("service_id"))
	amount, errAmount := strconv.Atoi(r.URL.Query().Get("amount"))

	if errId != nil {
		log.Println("Ошибка преобразования строки в число:", errId)
		return
	}

	if errId != nil {
		log.Println("Ошибка преобразования строки в число:", errAmount)
		return
	}

	payment := services.ReserveService{}.Create(userId, orderId, serviceId, amount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(payment)

}
