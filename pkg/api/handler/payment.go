package handler

import (
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	// "github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	// "github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	orderUseCase services.OrderUseCase
}

func NewPaymentHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}
