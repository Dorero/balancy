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
	balanceId := NewBalanceRepository().Show(userId).Id

	row := b.conn.QueryRow(context.Background(), "insert into payments (balance_id, amount) VALUES ($1, $2) returning id, balance_id, amount", balanceId, amount)
	err := row.Scan(&payment.Id, &payment.BalanceId, &payment.Amount)

	if err != nil {
		log.Println(err)
	}

	return &payment
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		conn: database.DBInstance(),
	}
}
