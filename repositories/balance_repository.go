package repositories

import (
	"balancy/database"
	"balancy/dto"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type BalanceRepository struct {
	conn *pgx.Conn
}

func (b *BalanceRepository) Show(userId int) *dto.Balance {
	var balance dto.Balance

	err := b.conn.QueryRow(context.Background(), "SELECT * FROM balance WHERE user_id=$1", userId).Scan(&balance.Id, &balance.UserId, &balance.Money)

	if err != nil {
		log.Fatal(err)
	}

	return &balance
}

func NewBalanceRepository() *BalanceRepository {
	return &BalanceRepository{
		conn: database.DBInstance(),
	}
}
