package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type PaymentRepository interface {
	GetPaymentMethods() ([]response.PaymentMethod, error)
	GetPaymentMethodCodId() (int, error)
	GetPaymentMethodRazorpayId() (int, error)
	FindPaymentMethodById(methodID int) (response.PaymentMethod, error)
}
