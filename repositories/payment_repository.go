package repositories

import (
	"balancy/database"
	"balancy/dto"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type PaymentRepository struct {
	conn *pgx.Conn
}

func (b *PaymentRepository) Create(userId int, amount int) *dto.Payment {
	var payment dto.Payment
	balance := NewBalanceRepository().Show(userId)
	tx, err := b.conn.Begin(context.Background())

	if err != nil {
		log.Println(err)
		return &payment
	}
	row := tx.QueryRow(context.Background(), "insert into payments (balance_id, amount) VALUES ($1, $2) returning id, balance_id, amount", balance.Id, amount)
	err = row.Scan(&payment.Id, &payment.BalanceId, &payment.Amount)

	if err != nil {
		log.Println(err)
		return &payment
	}

	_, err = tx.Exec(context.Background(), "update balance set money=$1 where user_id=$2", balance.Money+int64(amount), userId)

	err = tx.Commit(context.Background())

	if err != nil {
		log.Println(err)
		return &payment
	}

	return &payment
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		conn: database.DBInstance(),
	}
}
