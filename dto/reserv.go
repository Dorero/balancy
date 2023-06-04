package dto

type Reserve struct {
	Id        int   `json:"id"`
	UserId    int   `json:"user_id"`
	OrderId   int   `json:"order_id"`
	ServiceId int   `json:"service_id"`
	Amount    int64 `json:"amount"`
}
