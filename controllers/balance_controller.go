package controllers

import (
	"balancy/repositories"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type BalanceController struct {
	Db *sql.DB
}

func (c BalanceController) Show(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil {
		// Обработка ошибки, если преобразование не удалось
		fmt.Println("Ошибка преобразования строки в число:", err)
		return
	}

	repositories.BalanceRepository{}.Show(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
