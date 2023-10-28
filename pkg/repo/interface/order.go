package interfaces

import (
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type OrderRepository interface {
	GetUserOrderHistory(userID, startIndex, endIndex int) ([]response.Orders, error)
	InsertOrder(request.NewOrder) (response.OrderLine, error)
	ChangeOrderStatusByID(statusID int, orderID int) (response.OrderLine, error)
	FindOrderByUserIDAndProductID(userID, productID int) (response.OrderLine, error)
	FindOrderStatusByID(statusID int) (string, error)
	GetAllOrderData(startIndex, endIndex int) ([]response.Orders, error)
	FindOrderByID(orderID int) (response.OrderLine, error)
	FindOrdersBoughtUsingCoupon(couponID int) ([]response.OrderLine, error)
	InitializeNewUserWallet(userID int) (response.Wallet, error)
	FindUserWalletByID(userID int) (response.Wallet, error)
	UpdateUserWalletBalance(userID int, amount float32) (response.Wallet, error)
	GetStatusReturned() (response.OrderStatus, error)
	GetStatusCancelled() (response.OrderStatus, error)
	GetStatusPending() (response.OrderStatus, error)
	GetOrderStatuses() ([]response.OrderStatus, error)
	GetInvoiceDataByID(orderID int) (response.Orders, error)
	//
	UpdateWalletTransactionHistory(update request.WalletTransactionHistory) (response.WalletTransactionHistory, error)
	GetWalletHistoryByUserID(userID int) ([]response.WalletTransactionHistory, error)

	TopSellingProduct(startDate, endDate time.Time) (response.TopSelling, error)
	GetTotalSaleCount(startDate, endDate time.Time) (int, error)
	GetAverageOrderValue(startDate, endDate time.Time) (float32, error)
	GetTotalRevenue(returnID int, startDate, endDate time.Time) ([]response.OrderLine, error)
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
