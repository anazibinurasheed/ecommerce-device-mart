package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type OrderRepository interface {
	GetUserOrderHistory(userID, startIndex, endIndex int) ([]response.Orders, error)
	GetStatusPending() (response.OrderStatus, error)
	GetOrderStatuses() ([]response.OrderStatus, error)
	InsertOrder(request.NewOrder) (response.OrderLine, error)
	ChangeOrderStatusByID(statusID int, orderID int) (response.OrderLine, error)
	FindOrderDataByUserIDAndProductID(userID, productID int) (response.OrderLine, error)
	FindOrderStatusByID(statusID int) (string, error)
	GetAllOrderData(startIndex, endIndex int) ([]response.Orders, error)
	FindOrderByID(orderID int) (response.OrderLine, error)
	FindOrdersBoughtUsingCoupon(couponID int) ([]response.OrderLine, error)
	GetStatusCancelled() (response.OrderStatus, error)
	InitializeNewUserWallet(userID int) (response.Wallet, error)
	FindUserWalletByID(userID int) (response.Wallet, error)
	UpdateUserWalletBalance(userID int, amount float32) (response.Wallet, error)
	GetStatusReturned() (response.OrderStatus, error)
	GetInvoiceDataByID(orderID int) (response.Orders, error)
	//
	UpdateWalletTransactionHistory(update response.WalletTransactionHistory) (response.WalletTransactionHistory, error)
	GetWalletHistoryByUserID(userID int) ([]response.WalletTransactionHistory, error)
}

// GetUserOrderHistory
// GetPendingOrderStatus
// GetAllOrderStatuses
// InsertOrder
// ChangeOrderStatusByID
// FindOrderDataByUserIDAndProductID
// FindOrderStatusByID
// GetAllOrderData
// FindOrderByID
// FindOrdersUsedByCoupon
// GetCancelledOrderStatus
// InitializeNewUserWallet
// FindUserWalletByID
// UpdateUserWalletBalance
// GetReturnedOrderStatus
// GetInvoiceDataByID
// UpdateWalletTransactionHistoryEntry
// GetWalletHistoryByUserID
