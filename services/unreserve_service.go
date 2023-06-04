package services

import (
	"balancy/dto"
	"balancy/repositories"
)

type UnreserveService struct {
	ReserveRepository repositories.ReserveRepository
}

func (u UnreserveService) Create(userId int, orderId int, serviceId int, amount int) *dto.Unreserve {
	return repositories.NewUnreserveRepository().Create(userId, orderId, serviceId, amount)
}
