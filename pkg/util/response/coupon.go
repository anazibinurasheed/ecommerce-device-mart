package response

import "time"

type Coupon struct {
	ID                int       `json:"id"`
	Code              string    `json:"code" `
	CouponName        string    `json:"coupon_name"`
	MinOrderValue     float64   `json:"min_order_value"`
	DiscountPercent   float64   `json:"discount_percentage"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidFrom         time.Time `json:"valid_from"`
	ValidTill         time.Time `json:"valid_till"` //date and hours left
	ValidDays         int       `json:"-"`
	IsBlocked         bool      `json:"is_blocked"`
}
type CouponTracking struct {
	ID       int
	CouponID int
	userID   int
	IsUsed   bool
}
