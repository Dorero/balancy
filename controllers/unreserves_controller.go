package controllers

import (
	"database/sql"
	"net/http"
)

type UnReservesController struct {
	Db *sql.DB
}

func (c UnReservesController) Create(w http.ResponseWriter, r *http.Request) {

}
