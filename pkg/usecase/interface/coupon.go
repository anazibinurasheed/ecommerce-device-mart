package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type CouponUseCase interface {
	CreateCoupons(couponData request.Coupon) error
	ViewAllCoupons() ([]response.Coupon, error)
	UpdateCoupon(couponData request.Coupon, couponID int) error
	BlockCoupon(couponID int) error
	UnBlockCoupon(couponID int) error
	ProcessApplyCoupon(couponCode string, userID int) error
	ListOutAvailableCouponsToUser(userID int) ([]response.Coupon, error)
	RemoveFromCouponTracking(couponID, userID int) error
}
