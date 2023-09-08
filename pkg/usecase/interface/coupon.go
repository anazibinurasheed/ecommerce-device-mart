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

// type CouponUseCase interface {
// 	// CreateCoupon creates a new coupon.
// 	CreateCoupon(couponData request.Coupon) error

// 	// ViewAllCoupons retrieves a list of all coupons.
// 	ViewAllCoupons() ([]response.Coupon, error)

// 	// UpdateCoupon updates an existing coupon by its ID.
// 	UpdateCoupon(couponData request.Coupon, couponID int) error

// 	// BlockCoupon blocks a coupon by its ID.
// 	BlockCoupon(couponID int) error

// 	// UnBlockCoupon unblocks a coupon by its ID.
// 	UnBlockCoupon(couponID int) error

// 	// ApplyCoupon processes the application of a coupon by its code to a user's order.
// 	ApplyCoupon(couponCode string, userID int) error

// 	// ListAvailableCouponsForUser lists out available coupons for a user.
// 	ListAvailableCouponsForUser(userID int) ([]response.Coupon, error)

// 	// RemoveCouponTracking removes a user's tracking of a specific coupon.
// 	RemoveCouponTracking(couponID, userID int) error
// }
