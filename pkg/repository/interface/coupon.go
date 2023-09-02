package interfaces

import (
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type CouponRepository interface {
	CreateCoupon(couponData request.Coupon) (response.Coupon, error)
	GetAllCoupons() ([]response.Coupon, error)
	BlockCouponByID(couponID int) (response.Coupon, error)
	UnblockCouponByID(couponID int) (response.Coupon, error)
	UpdateCouponDetails(couponData request.Coupon) (response.Coupon, error)
	FindCouponByCode(couponCode string) (response.Coupon, error)
	FindCouponByID(couponID int) (response.Coupon, error)
	AddCouponTracking(userID, couponID int) (response.CouponTracking, error)
	UpdateCouponUsage(userID int) (response.CouponTracking, error)

	FindTrackingCoupon(userID, couponID int) (response.CouponTracking, error)
	//actual check applied coupon
	//CheckAppliedCoupon(userID int) (response.CouponTracking, error)

	//the findappliedcoupon is this name changed
	CheckAppliedCoupon(userID int) (response.CouponTracking, error)

	ChangeUserCoupon(couponID, userID int) (response.CouponTracking, error)
	GetAvailableCouponsForUser(userID int, currentTime time.Time) ([]response.Coupon, error)
	RemoveCouponFromTracking(couponID, userID int) (response.CouponTracking, error)
}

// type CouponRepository interface {
//     // Coupon management
//     CreateCoupon(couponData request.Coupon) (response.Coupon, error)
//     GetAllCoupons() ([]response.Coupon, error)
//     BlockCouponByID(couponID int) (response.Coupon, error)
//     UnblockCouponByID(couponID int) (response.Coupon, error)
//     UpdateCouponData(couponData request.Coupon) (response.Coupon, error)
//     FindCouponByCode(couponCode string) (response.Coupon, error)
//     FindCouponByID(couponID int) (response.Coupon, error)

//     // Coupon tracking
//     AddCouponTracking(userID, couponID int) (response.CouponTracking, error)
//     UpdateCouponUsage(userID int) (response.CouponTracking, error)
//     FindCouponTracking(userID, couponID int) (response.CouponTracking, error)
//     CheckAppliedCoupon(userID int) (response.CouponTracking, error)
//     ChangeUserCoupon(couponID, userID int) (response.CouponTracking, error)
//     FindAppliedCouponByUserID(userID int) (response.CouponTracking, error)
//     GetAvailableCouponsForUser(userID int, currentTime time.Time) ([]response.Coupon, error)
//     RemoveCouponFromTracking(couponID, userID int) (response.CouponTracking, error)
// }
