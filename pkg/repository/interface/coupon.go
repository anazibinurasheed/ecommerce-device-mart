package interfaces

import (
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type CouponRepository interface {
	InsertNewCoupon(couponData request.Coupon) (response.Coupon, error)
	VeiwAllCoupons() ([]response.Coupon, error)
	BlockCoupon(couponID int) (response.Coupon, error)
	UnblockCoupon(couponID int) (response.Coupon, error)
	UpdateCoupon(couponData request.Coupon) (response.Coupon, error)
	FindCouponByCode(couponCode string) (response.Coupon, error)
	FindCouponById(couponID int) (response.Coupon, error)
	InsertIntoCouponTracking(userID, couponID int) (response.CouponTracking, error)
	UpdateCouponUsage(userID int) (response.CouponTracking, error)
	FindCouponTracking(userID, couponID int) (response.CouponTracking, error)
	CheckForAppliedCoupon(userID int) (response.CouponTracking, error)
	ChangeCoupon(couponID, userID int) (response.CouponTracking, error)
	FindAppliedCouponByUserId(userID int) (response.CouponTracking, error)
	FetchAvailabeCouponsForUser(userID int, currentTime time.Time) ([]response.Coupon, error)
	RemoveCouponFromCouponTracking(couponID, userID int) (response.CouponTracking, error)
}
