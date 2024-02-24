package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	// razorpay "github.com/razorpay/razorpay-go"
)

type paymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
	cartUseCase services.CartUseCase
	userRepo    interfaces.UserRepository
}

func NewPaymentUseCase(PaymentUseCase interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentRepo: PaymentUseCase,
	}
}

func (ou *paymentUseCase) GetRazorPayDetails(userID int) (response.PaymentDetails, error) {
	userCart, err := ou.cartUseCase.ViewCart(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to retrieve userCart :%s", err)
	}

	userData, err := ou.userRepo.FindUserByID(userID)
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to find user  %s", err)
	}

	razorPayOrderID, err := helper.MakeRazorPayPaymentId(int(userCart.Total * 100))
	if err != nil {
		return response.PaymentDetails{}, fmt.Errorf("Failed to get razorpay id %s", err)
	}

	return response.PaymentDetails{
		Username:        userData.UserName,
		RazorPayOrderID: razorPayOrderID,
		Amount:          int(userCart.Total),
	}, nil
}

func (ou *paymentUseCase) VerifyRazorPayPayment(signature string, razorpayOrderID string, paymentID string) error {
	err := helper.VerifyRazorPayPayment(signature, razorpayOrderID, paymentID)
	if err != nil {
		return fmt.Errorf("Failed payment not success : %s", err)
	}
	return nil
}
