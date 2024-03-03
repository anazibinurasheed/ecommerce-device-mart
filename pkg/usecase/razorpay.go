package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	// razorpay "github.com/razorpay/razorpay-go"
)

type razorpayUseCase struct {
	paymentRepo interfaces.PaymentRepository
	cartUseCase services.CartUseCase
	userRepo    interfaces.UserRepository
}

func NewRazorpayUseCase(paymentRepo interfaces.PaymentRepository,
	cartUseCase services.CartUseCase,
	userRepo interfaces.UserRepository) services.RazorpayUseCase {
	return &razorpayUseCase{
		paymentRepo: paymentRepo,
		cartUseCase: cartUseCase,
		userRepo:    userRepo,
	}
}

func (ou *razorpayUseCase) GetRazorPayDetails(userID int) (response.PaymentDetails, error) {
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

func (ou *razorpayUseCase) VerifyRazorPayPayment(signature string, razorpayOrderID string, paymentID string) error {
	err := helper.VerifyRazorPayPayment(signature, razorpayOrderID, paymentID)
	if err != nil {
		return fmt.Errorf("Failed payment not success : %s", err)
	}
	return nil
}
