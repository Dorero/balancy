package services

import (
	"balancy/dto"
	"balancy/repositories"
)

type PaymentService struct {
	paymentRepository repositories.PaymentRepository
}

func (p PaymentService) Create(userId int, amount int) *dto.Payment {
	return repositories.NewPaymentRepository().Create(userId, amount)
}
