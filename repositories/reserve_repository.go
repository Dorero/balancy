package repositories

import (
	"balancy/database"
	"balancy/dto"
	"github.com/jackc/pgx/v5"
)

type ReserveRepository struct {
	conn *pgx.Conn
}

func (b *ReserveRepository) Create(userId int, orderId int, serviceId int, amount int) *dto.Reserve {
	var reserve dto.Reserve

	return &reserve
}

func NewReserveRepository() *ReserveRepository {
	return &ReserveRepository{
		conn: database.DBInstance(),
	}
}
