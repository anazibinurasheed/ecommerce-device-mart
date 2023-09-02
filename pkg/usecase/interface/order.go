package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type OrderUseCase interface {
	CheckOutDetails(userID int) (response.CheckOut, error)
	ConfirmedOrder(userID int, paymentMethodID int) error
	GetUserOrderHistory(userID, page, count int) ([]response.Orders, error)
	GetOrderManagement(page, count int) (response.OrderManagement, error)
	UpdateOrderStatus(statusID, orderID int) error
	AllOrderOverView(page, count int) ([]response.Orders, error)
	GetRazorPayDetails(userID int) (response.PaymentDetails, error)
	VerifyRazorPayPayment(signature string, razorpayOrderId string, paymentId string) error
	OrderCancellation(orderID int) error
	ProcessReturnRequest(orderID int) error
	GetUserWallet(userID int) (response.Wallet, error)
	CreateUserWallet(userID int) error
	ValidateWalletPayment(userID int) error
	CreateInvoice(orderID int) ([]byte, error)
}
