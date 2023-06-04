package repositories

import (
	"balancy/database"
	"balancy/dto"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type UnreserveRepository struct {
	conn *pgx.Conn
}

func (b *UnreserveRepository) Create(userId int, orderId int, serviceId int, amount int) *dto.Unreserve {
	var unreserve dto.Unreserve

	tx, err := b.conn.Begin(context.Background())

	if err != nil {
		log.Println(err)
		return &unreserve
	}

	sql := "insert into unreserves (user_id, order_id, service_id, amount) VALUES ($1, $2, $3, $4) returning id, user_id, order_id, service_id, amount"
	row := tx.QueryRow(context.Background(), sql, userId, orderId, serviceId, amount)
	err = row.Scan(&unreserve.Id, &unreserve.UserId, &unreserve.OrderId, &unreserve.ServiceId, &unreserve.Amount)

	if err != nil {
		tx.Rollback(context.Background())
		log.Println(err)
		return &unreserve
	}

	balance := NewBalanceRepository().Show(userId)

	_, err = tx.Exec(context.Background(), "insert into payments (balance_id, amount) VALUES ($1, $2)", balance.Id, amount)

	if err != nil {
		tx.Rollback(context.Background())
		log.Println(err)
		return &unreserve
	}

	_, err = tx.Exec(context.Background(), "update balance set money=$1 where user_id=$2", balance.Money+int64(amount), userId)

	if err != nil {
		tx.Rollback(context.Background())
		log.Println(err)
		return &unreserve
	}

	err = tx.Commit(context.Background())

	if err != nil {
		log.Println(err)
		return &unreserve
	}

	return &unreserve
}

func NewUnreserveRepository() *UnreserveRepository {
	return &UnreserveRepository{
		conn: database.DBInstance(),
	}
}
