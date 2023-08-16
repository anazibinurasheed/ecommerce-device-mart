package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
)

type referralUseCase struct {
	referralRepo interfaces.ReferralRepository
	orderUseCase interfaces.OrderRepository
}

func NewRefferalUseCase(referraluseCase interfaces.ReferralRepository, orderUseCase interfaces.OrderRepository) services.ReferralUseCase {
	return &referralUseCase{
		referralRepo: referraluseCase,
		orderUseCase: orderUseCase,
	}
}

func (ru *referralUseCase) GetUserReferralCode(userID int) (response.Referral, error) {
	refferalCode, err := ru.referralRepo.FindRefferalCodeByUserId(userID)
	if err != nil {
		return response.Referral{}, fmt.Errorf("Failed to find refferal code by user id : %s", err)
	}
	if refferalCode.ID != 0 {
		return refferalCode, nil
	}

	code := helper.GenerateReferralCode()

	newRefferalCode, err := ru.referralRepo.InsertNewRefferalCode(userID, code)
	if err != nil {
		return response.Referral{}, fmt.Errorf("Failed to create new refferal code :%s", err)
	}
	if newRefferalCode.ID == 0 || newRefferalCode.Code == "" {
		return response.Referral{}, fmt.Errorf("Failed to verify the refferal code ")

	}
	return newRefferalCode, nil

}

func (ru *referralUseCase) VerifyReferralCode(refferalCode string, claimingUserID int) (int, error) {
	if refferalCode == "" {
		return -1, fmt.Errorf("No refferal code provided")
	}
	refferalCodeDetails, err := ru.referralRepo.FindRefferalCodeByCode(refferalCode)
	if err != nil {
		return -1, fmt.Errorf("Failed to find refferal code : %s", err)
	}

	if refferalCodeDetails.ID == 0 {
		return -1, fmt.Errorf("Invalid, refferal code doesn't exist")
	}

	if refferalCodeDetails.ID == uint(claimingUserID) {
		return -1, fmt.Errorf("Not allowed to use this coupon")
	}

	codeOwnerID := refferalCodeDetails.UserID
	return int(codeOwnerID), nil
}

func (ru *referralUseCase) ClaimReferralBonus(claimingUserID, codeOwnerID int) error {
	referredUserWallet, err := ru.orderUseCase.FindUserWallet(int(codeOwnerID))
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	}
	if referredUserWallet.ID == 0 {

		newReferredUserWallet, err := ru.orderUseCase.InitializeNewWallet(int(codeOwnerID))
		if err != nil {
			return fmt.Errorf("Failed to initialize wallet for code owner: %s", err)
		}
		if newReferredUserWallet.ID == 0 {
			return fmt.Errorf("Failed to verify code owner new wallet")
		}
	}

	bonusClaimingUser, err := ru.orderUseCase.FindUserWallet(int(claimingUserID))
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	}
	if bonusClaimingUser.ID == 0 {
		newClamingUserWallet, err := ru.orderUseCase.InitializeNewWallet(claimingUserID)
		if err != nil {
			return fmt.Errorf("Failed to initialize wallet for code owner: %s", err)
		}
		if newClamingUserWallet.ID == 0 {
			return fmt.Errorf("Failed to verify code owner new wallet")
		}

	}

	referredUserWallet, err = ru.orderUseCase.UpdateUserWallet(int(codeOwnerID), (referredUserWallet.Amount + 50))
	if err != nil {
		return fmt.Errorf("Failed update bonus on  : %s", err)
	}
	if referredUserWallet.ID == 0 {
		return fmt.Errorf("Failed to verify bonus updated wallet")
	}

	bonusClaimingUser, err = ru.orderUseCase.UpdateUserWallet(int(claimingUserID), (bonusClaimingUser.Amount + 50))
	if err != nil {
		return fmt.Errorf("Failed update bonus  : %s", err)
	}
	if referredUserWallet.ID == 0 {
		return fmt.Errorf("Failed to verify bonus updated wallet")
	}

	return nil
}
