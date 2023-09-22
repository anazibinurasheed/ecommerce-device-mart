package usecase

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	// razorpay "github.com/razorpay/razorpay-go"
)

type paymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
}

func NewPaymentUseCase(PaymentUseCase interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentRepo: PaymentUseCase,
	}
}
