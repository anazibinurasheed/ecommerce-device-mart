package usecase

import (
	"fmt"
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
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
	Days := couponData.ValidityDays
	couponData.ValidFrom = time.Now()
	couponData.ValidTill = time.Now().AddDate(0, 0, Days)

	InsertedCoupon, err := cu.couponRepo.InsertNewCoupon(couponData)
	if err != nil {
		return fmt.Errorf("Failed to insert coupon :%s", err)
	}
	if InsertedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify inserted coupon ")

	}
	return nil
}

func (cu *couponUseCase) ViewAllCoupons() ([]response.Coupon, error) {
	Coupons, err := cu.couponRepo.VeiwAllCoupons()
	if err != nil {
		return nil, fmt.Errorf("Failed to insert coupon :%s", err)
	}

	return Coupons, nil
}
func (cu *couponUseCase) UpdateCoupon(couponData request.Coupon, couponID int) error {
	Days := couponData.ValidityDays
	couponData.ValidFrom = time.Now()
	couponData.ValidTill = time.Now().AddDate(0, 0, Days)
	couponData.ID = uint(couponID)
	UpdatedCoupon, err := cu.couponRepo.UpdateCoupon(couponData)
	if err != nil {
		return fmt.Errorf("Failed to update coupon : %s", err)
	}

	if UpdatedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify updated coupon by id ")
	}

	return nil
}

func (cu *couponUseCase) BlockCoupon(couponID int) error {
	BlockedCoupon, err := cu.couponRepo.BlockCoupon(couponID)
	if err != nil {
		return fmt.Errorf("Failed to block coupon : %s ", err)
	}
	if BlockedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify blocked coupon by id ")
	}
	return nil

}
func (cu *couponUseCase) UnBlockCoupon(couponID int) error {
	UnBlockedCoupon, err := cu.couponRepo.UnblockCoupon(couponID)
	if err != nil {
		return fmt.Errorf("Failed to unblock coupon : %s ", err)
	}
	if UnBlockedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify unblocked coupon by id ")
	}
	return nil

}

func (cu *couponUseCase) ProcessApplyCoupon(couponCode string, userID int) error {

	Coupon, err := cu.couponRepo.FindCouponByCode(couponCode)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("1::::", Coupon)
	if err != nil {
		return fmt.Errorf("Failed to find coupon  :%s", err)
	}

	if Coupon.IsBlocked || !helper.IsCouponValid(Coupon.ValidTill) || Coupon.ID == 0 {
		return fmt.Errorf("coupon cant use ,invalid coupon")
	}

	TrackingDetails, err := cu.couponRepo.FindCouponTracking(userID, Coupon.ID)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("2:::COUPON TRACKING:", TrackingDetails)

	if err != nil {
		return fmt.Errorf("Failed to find coupon tracking details : %s", err)
	}

	if TrackingDetails.CouponID == Coupon.ID && TrackingDetails.IsUsed {
		return fmt.Errorf("Failed coupon already used")
	} else if TrackingDetails.CouponID == Coupon.ID && !TrackingDetails.IsUsed {
		return nil
	}
	PreviousCoupon, err := cu.couponRepo.CheckForAppliedCoupon(userID)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("3::::", PreviousCoupon)

	if err != nil {
		return fmt.Errorf("Failed to find  previous coupon details from coupon tracking ")
	}
	if PreviousCoupon.ID != 0 && !PreviousCoupon.IsUsed {
		ChangedCoupon, err := cu.couponRepo.ChangeCoupon(Coupon.ID, userID)
		if err != nil {
			return fmt.Errorf("Failed to change coupon : %s", err)

		} else if ChangedCoupon.ID == 0 {
			return fmt.Errorf("Failed to verify changed coupon from coupon tracking")
		}
	} else {
		InsertedRecord, err := cu.couponRepo.InsertIntoCouponTracking(userID, Coupon.ID)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("4::::", InsertedRecord)
		fmt.Println("")
		fmt.Println("")
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
	AllCoupons, err := cu.couponRepo.FetchAvailabeCouponsForUser(userID, currentTime)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch available  coupons")
	}
	return AllCoupons, nil

}
func (cu *couponUseCase) RemoveFromCouponTracking(couponID, userID int) error {
	fmt.Println("COUPON AND USERID", couponID, userID)
	RemovedCoupon, err := cu.couponRepo.RemoveCouponFromCouponTracking(couponID, userID)
	fmt.Println("REMOVE COUPON TRACKING ", RemovedCoupon)
	if err != nil {
		return fmt.Errorf("Failed to remove coupon from coupon tracking :%s", err)
	}
	if RemovedCoupon.ID == 0 {
		return fmt.Errorf("Failed to verify the removed coupon  ")
	}
	return nil
}