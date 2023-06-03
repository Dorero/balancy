package dto

type Payment struct {
	Id        int   `json:"id"`
	BalanceId int   `json:"balance_id"`
	Amount    int64 `json:"amount"`
}
