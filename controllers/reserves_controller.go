package controllers

import (
	"balancy/services"
	"balancy/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

type ReservesController struct {
	Db *sql.DB
}

func (c ReservesController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	payload := utils.Validator{}.IsNumbers(r.URL.Query())

	reserve := services.ReserveService{}.Create(payload["user_id"], payload["order_id"], payload["service_id"], payload["amount"])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(reserve)

}
