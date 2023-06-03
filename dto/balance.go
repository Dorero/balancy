package dto

type Balance struct {
	Id     int   `json:"id" db:"id"`
	UserId int   `json:"user_id" db:"user_id"`
	Money  int64 `json:"money" db:"money"`
}
