package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type OrderUseCase interface {
	CheckOutDetails(userID int) (response.Checkout, error)
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
	CreateInvoice(orderID int) (response.Invoice, error)
	MonthlySalesReport() (response.MonthlySalesReport, error)
	GetWalletHistory(userID int) ([]response.WalletTransactionHistory, error)
}

// type OrderUseCase interface {
// 	// GetCheckoutDetails retrieves checkout details for a user's order.
// 	GetCheckoutDetails(userID int) (response.CheckOut, error)

// 	// ConfirmOrder confirms an order for a user using a specified payment method.
// 	ConfirmOrder(userID int, paymentMethodID int) error

// 	// GetUserOrderHistory retrieves the order history for a user with pagination.
// 	GetUserOrderHistory(userID, page, count int) ([]response.Orders, error)

// 	// GetOrderManagement retrieves order management details with pagination.
// 	GetOrderManagement(page, count int) (response.OrderManagement, error)

// 	// UpdateOrderStatus updates the status of an order.
// 	UpdateOrderStatus(statusID, orderID int) error

// 	// GetAllOrderOverview retrieves an overview of all orders with pagination.
// 	GetAllOrderOverview(page, count int) ([]response.Orders, error)

// 	// GetRazorPayDetails retrieves RazorPay payment details for a user.
// 	GetRazorPayDetails(userID int) (response.PaymentDetails, error)

// 	// VerifyRazorPayPayment verifies a RazorPay payment using the provided signature and IDs.
// 	VerifyRazorPayPayment(signature string, razorpayOrderID string, paymentID string) error

// 	// CancelOrder cancels an order by its ID.
// 	CancelOrder(orderID int) error

// 	// ProcessReturnRequest processes a return request for an order.
// 	ProcessReturnRequest(orderID int) error

// 	// GetUserWallet retrieves the wallet details for a user.
// 	GetUserWallet(userID int) (response.Wallet, error)

// 	// CreateUserWallet creates a wallet for a user.
// 	CreateUserWallet(userID int) error

// 	// ValidateWalletPayment validates a payment using a user's wallet.
// 	ValidateWalletPayment(userID int) error

// 	// CreateInvoice generates an invoice for an order and returns it as a byte slice.
// 	CreateInvoice(orderID int) ([]byte, error)
// }
