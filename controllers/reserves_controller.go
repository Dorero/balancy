package controllers

import (
	"database/sql"
	"net/http"
)

type ReservesController struct {
	Db *sql.DB
}

func (c ReservesController) Index(http.ResponseWriter, *http.Request) {

}
