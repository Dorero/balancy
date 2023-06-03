package dto

type Unreserve struct {
	Id        int   `json:"id"`
	Amount    int64 `json:"amount"`
	OrderId   int   `json:"order_id"`
	ServiceId int   `json:"service_id"`
}
