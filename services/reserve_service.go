package services

import (
	"balancy/dto"
	"balancy/repositories"
)

type ReserveService struct {
	ReserveRepository repositories.ReserveRepository
}

func (p ReserveService) Create(userId int, orderId int, serviceId int, amount int) *dto.Reserve {
	return repositories.NewReserveRepository().Create(userId, orderId, serviceId, amount)
}
