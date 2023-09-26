package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
)

type referralUseCase struct {
	referralRepo interfaces.ReferralRepository
	orderUseCase interfaces.OrderRepository
}

func NewReferralUseCase(referraluseCase interfaces.ReferralRepository, orderUseCase interfaces.OrderRepository) services.ReferralUseCase {
	return &referralUseCase{
		referralRepo: referraluseCase,
		orderUseCase: orderUseCase,
	}
}

func (ru *referralUseCase) GetUserReferralCode(userID int) (response.Referral, error) {
	referralCode, err := ru.referralRepo.FindReferralCodeByUserID(userID)
	if err != nil {
		return response.Referral{}, fmt.Errorf("Failed to find referral code by user id : %s", err)
	}
	if referralCode.ID != 0 {
		return referralCode, nil
	}

	code := helper.GenerateReferralCode()

	newReferralCode, err := ru.referralRepo.InsertNewReferralCode(userID, code)
	if err != nil {
		return response.Referral{}, fmt.Errorf("Failed to create new refferal code :%s", err)
	}
	if newReferralCode.ID == 0 || newReferralCode.Code == "" {
		return response.Referral{}, fmt.Errorf("Failed to verify the refferal code ")

	}
	return newReferralCode, nil

}

func (ru *referralUseCase) VerifyReferralCode(referralCode string, claimingUserID int) (int, error) {
	if referralCode == "" {
		return -1, fmt.Errorf("No referral code provided")
	}

	referralCodeDetails, err := ru.referralRepo.FindReferralCodeByCode(referralCode)
	if err != nil {
		return -1, fmt.Errorf("Failed to find referral code : %s", err)
	}

	if referralCodeDetails.ID == 0 {
		return -1, fmt.Errorf("Invalid, referral code doesn't exist")
	}

	if referralCodeDetails.ID == uint(claimingUserID) {
		return -1, fmt.Errorf("Not allowed to use this coupon")
	}

	referralCodeOwnerID := referralCodeDetails.UserID
	return int(referralCodeOwnerID), nil
}

func (ru *referralUseCase) ClaimReferralBonus(claimingUserID, codeOwnerID int) error {
	referredUserWallet, err := ru.orderUseCase.FindUserWalletByID(int(codeOwnerID))
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	}
	if referredUserWallet.ID == 0 {

		newReferredUserWallet, err := ru.orderUseCase.InitializeNewUserWallet(int(codeOwnerID))
		if err != nil {
			return fmt.Errorf("Failed to initialize wallet for code owner: %s", err)
		}
		if newReferredUserWallet.ID == 0 {
			return fmt.Errorf("Failed to verify code owner new wallet")
		}
	}

	bonusClaimingUserWallet, err := ru.orderUseCase.FindUserWalletByID(int(claimingUserID))
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	}
	if bonusClaimingUserWallet.ID == 0 {
		bonusClaimingUserWallet, err := ru.orderUseCase.InitializeNewUserWallet(claimingUserID)
		if err != nil {
			return fmt.Errorf("Failed to initialize wallet for code owner: %s", err)
		}
		if bonusClaimingUserWallet.ID == 0 {
			return fmt.Errorf("Failed to verify code owner new wallet")
		}

	}

	referredUserWallet, err = ru.orderUseCase.UpdateUserWalletBalance(int(codeOwnerID), (referredUserWallet.Amount + 50))
	if err != nil {
		return fmt.Errorf("Failed update bonus on  : %s", err)
	}
	if referredUserWallet.ID == 0 {
		return fmt.Errorf("Failed to verify bonus updated wallet")
	}

	bonusClaimingUserWallet, err = ru.orderUseCase.UpdateUserWalletBalance(int(claimingUserID), (bonusClaimingUserWallet.Amount + 50))
	if err != nil {
		return fmt.Errorf("Failed update bonus  : %s", err)
	}
	if referredUserWallet.ID == 0 {
		return fmt.Errorf("Failed to verify bonus updated wallet")
	}

	return nil
}
