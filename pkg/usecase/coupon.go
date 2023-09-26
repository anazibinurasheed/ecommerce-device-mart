package usecase

import (
	"fmt"
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type couponUseCase struct {
	couponRepo interfaces.CouponRepository
}

func NewCouponUseCase(couponRepo interfaces.CouponRepository) services.CouponUseCase {
	return &couponUseCase{
		couponRepo: couponRepo,
	}

}

func (cu *couponUseCase) CreateCoupons(couponData request.Coupon) error {
	validDays := couponData.ValidityDays
	couponData.ValidFrom = time.Now()
	couponData.ValidTill = time.Now().AddDate(0, 0, validDays)

	insertedCoupon, err := cu.couponRepo.CreateCoupon(couponData)
	if err != nil {
		return fmt.Errorf("Failed to insert coupon :%s", err)
	}
	if insertedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify inserted coupon ")

	}
	return nil
}

func (cu *couponUseCase) ViewAllCoupons() ([]response.Coupon, error) {
	coupons, err := cu.couponRepo.GetAllCoupons()
	if err != nil {
		return nil, fmt.Errorf("Failed to insert coupon :%s", err)
	}

	return coupons, nil
}
func (cu *couponUseCase) UpdateCoupon(couponData request.Coupon, couponID int) error {
	validDays := couponData.ValidityDays
	couponData.ValidFrom = time.Now()
	couponData.ValidTill = time.Now().AddDate(0, 0, validDays)
	couponData.ID = uint(couponID)

	updatedCoupon, err := cu.couponRepo.UpdateCouponDetails(couponData)
	if err != nil {
		return fmt.Errorf("Failed to update coupon : %s", err)
	}

	if updatedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify updated coupon by id ")
	}

	return nil
}

func (cu *couponUseCase) BlockCoupon(couponID int) error {
	blockedCoupon, err := cu.couponRepo.BlockCouponByID(couponID)
	if err != nil {
		return fmt.Errorf("Failed to block coupon : %s ", err)
	}
	if blockedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify blocked coupon by id ")
	}
	return nil

}
func (cu *couponUseCase) UnBlockCoupon(couponID int) error {
	unBlockedCoupon, err := cu.couponRepo.UnblockCouponByID(couponID)
	if err != nil {
		return fmt.Errorf("Failed to unblock coupon : %s ", err)
	}
	if unBlockedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify unblocked coupon by id ")
	}
	return nil

}

func (cu *couponUseCase) ProcessApplyCoupon(couponCode string, userID int) error {

	coupon, err := cu.couponRepo.FindCouponByCode(couponCode)

	if err != nil {
		return fmt.Errorf("Failed to find coupon  :%s", err)
	}

	if coupon.IsBlocked || !helper.IsCouponValid(coupon.ValidTill) || coupon.ID == 0 {
		return fmt.Errorf("coupon cant use ,invalid coupon")
	}

	couponTrackingDetails, err := cu.couponRepo.FindTrackingCoupon(userID, coupon.ID)

	if err != nil {
		return fmt.Errorf("Failed to find coupon tracking details : %s", err)
	}

	if couponTrackingDetails.CouponID == coupon.ID && couponTrackingDetails.IsUsed {
		return fmt.Errorf("Failed coupon already used")
	} else if couponTrackingDetails.CouponID == coupon.ID && !couponTrackingDetails.IsUsed {
		return nil
	}

	previousCoupon, err := cu.couponRepo.CheckAppliedCoupon(userID)

	if err != nil {
		return fmt.Errorf("Failed to find  previous coupon details from coupon tracking ")
	}
	if previousCoupon.ID != 0 && !previousCoupon.IsUsed {

		changedCoupon, err := cu.couponRepo.ChangeUserCoupon(coupon.ID, userID)
		if err != nil {
			return fmt.Errorf("Failed to change coupon : %s", err)

		}
		if changedCoupon.ID == 0 {
			return fmt.Errorf("Failed to verify changed coupon from coupon tracking")
		}

	} else {
		InsertedRecord, err := cu.couponRepo.AddCouponTracking(userID, coupon.ID)
		if err != nil {
			return fmt.Errorf("Failed to insert tracking record : %s", err)
		}
		if InsertedRecord.ID == 0 {

			return fmt.Errorf("Failed to verify inserted coupon tracking by id ")
		}
	}
	return nil
}

func (cu *couponUseCase) ListOutAvailableCouponsToUser(userID int) ([]response.Coupon, error) {
	currentTime := time.Now()
	allCoupons, err := cu.couponRepo.GetAvailableCouponsForUser(userID, currentTime)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch available  coupons")
	}
	return allCoupons, nil

}

func (cu *couponUseCase) RemoveFromCouponTracking(couponID, userID int) error {
	removedCoupon, err := cu.couponRepo.RemoveCouponFromTracking(couponID, userID)
	if err != nil {
		return fmt.Errorf("Failed to remove coupon from coupon tracking :%s", err)
	}
	if removedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify the removed coupon")
	}
	return nil
}
