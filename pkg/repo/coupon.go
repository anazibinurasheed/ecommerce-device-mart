package repo

import (
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type couponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponDatabase{
		DB: DB,
	}
}

func (cd *couponDatabase) CreateCoupon(couponData request.Coupon) (response.Coupon, error) {
	var InsertedCoupon response.Coupon

	query := `INSERT INTO coupons (coupon_name,code,min_order_value,discount_percent,discount_max_amount,valid_till,valid_from,valid_days)VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *;`
	err := cd.DB.Raw(query, couponData.CouponName, couponData.Code, couponData.MinOrderValue, couponData.DiscountPercent, couponData.DiscountMaxAmount, couponData.ValidTill, couponData.ValidFrom, couponData.ValidityDays).Scan(&InsertedCoupon).Error

	return InsertedCoupon, err
}

func (cd *couponDatabase) UpdateCouponDetails(couponData request.Coupon) (response.Coupon, error) {
	var UpdatedCoupon response.Coupon
	query := `UPDATE coupons SET coupon_name = $1 ,code = $2 ,min_order_value = $3 ,discount_percent = $4 ,discount_max_amount = $5,valid_till= $6 ,valid_from= $7,valid_days = $8 WHERE id = $9  RETURNING *;`
	err := cd.DB.Raw(query, couponData.CouponName, couponData.Code, couponData.MinOrderValue, couponData.DiscountPercent, couponData.DiscountMaxAmount, couponData.ValidTill, couponData.ValidFrom, couponData.ValidityDays, couponData.ID).Scan(&UpdatedCoupon).Error

	return UpdatedCoupon, err
}
func (cd *couponDatabase) BlockCouponByID(couponID int) (response.Coupon, error) {
	var BlockedCoupon response.Coupon
	IsBlocked := true
	query := `UPDATE coupons SET Is_blocked = $1 WHERE id = $2 RETURNING *;`
	err := cd.DB.Raw(query, IsBlocked, couponID).Scan(&BlockedCoupon).Error

	return BlockedCoupon, err
}

func (cd *couponDatabase) UnblockCouponByID(couponID int) (response.Coupon, error) {
	var BlockedCoupon response.Coupon
	IsBlocked := false
	query := `UPDATE coupons SET Is_blocked = $1 WHERE id = $2 RETURNING *;`
	err := cd.DB.Raw(query, IsBlocked, couponID).Scan(&BlockedCoupon).Error
	return BlockedCoupon, err
}

func (cd *couponDatabase) GetAllCoupons() ([]response.Coupon, error) {
	var Coupons = make([]response.Coupon, 0)

	query := `SELECT * FROM coupons ORDER BY id DESC`
	err := cd.DB.Raw(query).Scan(&Coupons).Error

	return Coupons, err
}

func (cd *couponDatabase) FindCouponByCode(couponCode string) (response.Coupon, error) {
	var Coupon response.Coupon
	query := `SELECT * FROM coupons WHERE code = $1 ;`
	err := cd.DB.Raw(query, couponCode).Scan(&Coupon).Error

	return Coupon, err

}

func (cd *couponDatabase) FindCouponByID(couponID int) (response.Coupon, error) {
	var Coupon response.Coupon
	query := `SELECT * FROM coupons WHERE id = $1 ;`
	err := cd.DB.Raw(query, couponID).Scan(&Coupon).Error

	return Coupon, err

}

func (cd *couponDatabase) AddCouponTracking(userID, couponID int) (response.CouponTracking, error) {
	var TrackedCoupon response.CouponTracking

	query := `INSERT INTO coupon_trackings (coupon_id,user_id)VALUES($1,$2) RETURNING * ;`
	err := cd.DB.Raw(query, couponID, userID).Scan(&TrackedCoupon).Error
	return TrackedCoupon, err
}

func (cd *couponDatabase) UpdateCouponUsage(userID int) (response.CouponTracking, error) {
	var UpdatedCouponTracking response.CouponTracking
	IsUsed := true
	query := `UPDATE coupon_trackings SET is_used = $1 WHERE user_id = $2 AND is_used = false RETURNING * ;`
	err := cd.DB.Raw(query, IsUsed, userID).Scan(&UpdatedCouponTracking).Error
	return UpdatedCouponTracking, err
}

func (cd *couponDatabase) FindTrackingCoupon(userID, couponID int) (response.CouponTracking, error) {
	var coupon response.CouponTracking
	query := `SELECT * FROM coupon_trackings WHERE coupon_id = $1 AND user_id = $2;`
	err := cd.DB.Raw(query, couponID, userID).Scan(&coupon).Error
	return coupon, err
}

func (cd *couponDatabase) CheckAppliedCoupon(userID int) (response.CouponTracking, error) {
	var AppliedCoupon response.CouponTracking
	IsUsed := false
	query := `SELECT * FROM coupon_trackings WHERE user_id = $1 AND is_used = $2;`
	err := cd.DB.Raw(query, userID, IsUsed).Scan(&AppliedCoupon).Error
	return AppliedCoupon, err
}

// func (cd *couponDatabase) FindAppliedCouponByUserID(userID int) (response.CouponTracking, error) {
// 	var Coupon response.CouponTracking
// 	query := `SELECT * FROM coupon_trackings WHERE user_id = $1 AND is_used !=true ;`
// 	err := cd.DB.Raw(query, userID).Scan(&Coupon).Error

// 	return Coupon, err

// }

func (cd *couponDatabase) ChangeUserCoupon(couponID, userID int) (response.CouponTracking, error) {
	var ChangedCoupon response.CouponTracking
	query := `UPDATE coupon_trackings SET coupon_id = $1 WHERE user_id = $2 RETURNING *;`
	err := cd.DB.Raw(query, couponID, userID).Scan(&ChangedCoupon).Error
	return ChangedCoupon, err
}

func (cd *couponDatabase) GetAvailableCouponsForUser(userID int, currentTime time.Time) ([]response.Coupon, error) {
	var Coupons = make([]response.Coupon, 0)
	query := `SELECT *
	FROM coupons c
	LEFT JOIN coupon_trackings ct ON c.id = ct.coupon_id AND ct.user_id = $1
	WHERE c.is_blocked = false
	  AND c.valid_till >$2
	  AND (ct.is_used = false OR ct.is_used IS NULL);`
	err := cd.DB.Raw(query, userID, currentTime).Scan(&Coupons).Error

	return Coupons, err

}

func (cd *couponDatabase) RemoveCouponFromTracking(couponID, userID int) (response.CouponTracking, error) {
	var RemovedCoupon response.CouponTracking
	query := `DELETE FROM coupon_trackings WHERE coupon_id = $1 AND user_id = $2 RETURNING *;`
	err := cd.DB.Raw(query, couponID, userID).Scan(&RemovedCoupon).Error

	return RemovedCoupon, err

}
