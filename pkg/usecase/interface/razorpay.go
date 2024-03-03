package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type RazorpayUseCase interface {
	GetRazorPayDetails(userID int) (response.PaymentDetails, error)
	VerifyRazorPayPayment(signature string, razorpayOrderID string, paymentID string) error
}
