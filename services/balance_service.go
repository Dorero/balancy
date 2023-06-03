package services

import (
	"balancy/dto"
	"balancy/repositories"
)

type BalanceService struct {
	balanceRepository repositories.BalanceRepository
}

func (b BalanceService) Show(userId int) *dto.Balance {
	return repositories.NewBalanceRepository().Show(userId)
}
