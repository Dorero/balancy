package repositories

import (
	"balancy/database"
	"balancy/dto"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type ReserveRepository struct {
	conn *pgx.Conn
}

func (b *ReserveRepository) Create(userId int, orderId int, serviceId int, amount int) *dto.Reserve {
	var reserve dto.Reserve
	sql := "insert into reserves (user_id, order_id, service_id, amount) VALUES ($1, $2, $3, $4) returning id, user_id, order_id, service_id, amount"
	row := b.conn.QueryRow(context.Background(), sql, userId, orderId, serviceId, amount)
	err := row.Scan(&reserve.Id, &reserve.UserId, &reserve.OrderId, &reserve.ServiceId, &reserve.Amount)

	if err != nil {
		log.Println(err)
	}

	return &reserve
}

func NewReserveRepository() *ReserveRepository {
	return &ReserveRepository{
		conn: database.DBInstance(),
	}
}
