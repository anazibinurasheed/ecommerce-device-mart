package request

import "time"

type Coupon struct {
	ID                uint      `json:"-"`
	Code              string    `json:"code" binding:"required"`
	CouponName        string    `json:"coupon_name" binding:"required"`
	MinOrderValue     float64   `json:"min_order_value" binding:"required"`
	DiscountPercent   float64   `json:"discount_percentage" binding:"required"`
	DiscountMaxAmount float64   `json:"discount_max_amount" binding:"required"`
	ValidityDays      int       `json:"validity_days" binding:"required,gte=1"`
	ValidFrom         time.Time `json:"-"`
	ValidTill         time.Time `json:"-"`
}
