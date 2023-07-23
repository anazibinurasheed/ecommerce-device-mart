package usecase

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
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

// func (p *paymentUseCase) MakePaymentRazorPay(orderID string, userID int) (models.CombinedOrderDetails, string, error) {

// 	client := razorpay.NewClient("rzp_test_6m0J6O6Dngl96V", "F9zSviAWO3DIXnNAtKgrufzT")

// 	data := map[string]interface{}{
// 		"amount":   int(combinedOrderDetails.FinalPrice) * 100,
// 		"currency": "INR",
// 		"receipt":  "some_receipt_id",
// 	}
// 	body, err := client.Order.Create(data, nil)
// 	fmt.Println(body)
// 	razorPayOrderID := body["id"].(string)

// 	err = p.orderRepository.AddRazorPayDetails(orderID, razorPayOrderID)
// 	if err != nil {
// 		return models.CombinedOrderDetails{}, "", err
// 	}

// 	return combinedOrderDetails, razorPayOrderID, nil

// }
