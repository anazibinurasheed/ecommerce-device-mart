package interfaces

import (
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type OrderRepository interface {
	// InsertShopOrder(userID int, addressID int, paymentMethodID int, orderTotal float32, orderStatusId int) (response.ShopOrder, error)
	UserOrderHistory(userID, startIndex, endIndex int) ([]response.Orders, error)
	GetStatusPending() (response.OrderStatus, error)
	GetOrderStatuses() ([]response.OrderStatus, error)
	InsertOrderLine(userID int, productID int, addressID int, qty int, price int, paymentMethodID int, orderStatusID int, couponID int, createdAt time.Time, updatedAt time.Time) (response.OrderLine, error)
	ChangeOrderStatus(statusID int, orderID int) (response.OrderLine, error)
	FindOrderDataByUseridAndProductid(userID, productID int) (response.OrderLine, error)
	FindOrderStatusById(statusID int) (string, error)
	AllOrderData(startIndex, endIndex int) ([]response.Orders, error)
	FindOrderById(orderID int) (response.OrderLine, error)
	FindOrdersUsedByCoupon(couponID int) ([]response.OrderLine, error)
	GetStatusCancelled() (response.OrderStatus, error)
	InitializeNewWallet(userID int) (response.Wallet, error)
	FindUserWallet(userID int) (response.Wallet, error)
	UpdateUserWallet(userID int, amount float32) (response.Wallet, error)
	GetStatusReturned() (response.OrderStatus, error)
	GetInvoiceData(orderID int) (response.Orders, error)
}
