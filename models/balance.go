package models

type Balance struct {
	Id     int   `json:"id"`
	UserId int   `json:"user_id"`
	Money  int64 `json:"money"`
}
