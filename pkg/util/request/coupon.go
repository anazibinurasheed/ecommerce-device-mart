package request

import "time"

type Coupon struct {
	ID                uint      `json:"-"`
	Code              string    `json:"code" validate:"required"`
	CouponName        string    `json:"coupon_name" validate:"required"`
	MinOrderValue     float64   `json:"min_order_value" validate:"required"`
	DiscountPercent   float64   `json:"discount_percentage" validate:"required"`
	DiscountMaxAmount float64   `json:"discount_max_amount" validate:"required"`
	ValidityDays      int       `json:"validity_days"`
	ValidFrom         time.Time `json:"-"`
	ValidTill         time.Time `json:"-"`
}
