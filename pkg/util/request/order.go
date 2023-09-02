package request

import "time"

type NewOrder struct {
	UserID          int
	ProductID       int
	AddressID       int
	Qty             int
	Price           int
	PaymentMethodID int
	OrderStatusID   int
	CouponID        int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
