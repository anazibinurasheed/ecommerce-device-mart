package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"

	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
)

type refferalUseCase struct {
	refferalRepo interfaces.RefferalRepository
	orderUseCase interfaces.OrderRepository
}

func NewRefferalUseCase(refferaluseCase interfaces.RefferalRepository, orderUseCase interfaces.OrderRepository) services.RefferalUseCase {
	return &refferalUseCase{
		refferalRepo: refferaluseCase,
		orderUseCase: orderUseCase,
	}
}

func (ru *refferalUseCase) GetUserRefferalCode(userID int) (response.Referral, error) {

	RefferalCode, err := ru.refferalRepo.FindRefferalCodeByUserId(userID)
	fmt.Println("::::::::::1", RefferalCode)
	if err != nil {
		return response.Referral{}, fmt.Errorf("Failed to find refferal code by user id : %s", err)
	} else if RefferalCode.ID != 0 {
		fmt.Println(RefferalCode)
		return RefferalCode, nil
	}

	Code := helper.GenerateReferralCode()
	NewRefferalCode, err := ru.refferalRepo.InsertNewRefferalCode(userID, Code)
	fmt.Println(":::::::::::2", NewRefferalCode)
	if err != nil {
		return response.Referral{}, fmt.Errorf("Failed to create new refferal code :%s", err)
	} else if NewRefferalCode.ID == 0 || NewRefferalCode.Code == "" {
		return response.Referral{}, fmt.Errorf("Failed to verify the refferal code ")

	}
	return NewRefferalCode, nil

}

func (ru *refferalUseCase) VerifyRefferalCode(refferalCode string, claimingUserID int) (int, error) {

	if refferalCode == "" {
		return -1, fmt.Errorf("No refferal code provided")
	}
	RefferalDetails, err := ru.refferalRepo.FindRefferalCodeByCode(refferalCode)
	if err != nil {
		return -1, fmt.Errorf("Failed to find refferal code : %s", err)
	}

	if RefferalDetails.ID == 0 {
		return -1, fmt.Errorf("Invalid,Refferal code doesnt exist")
	}

	if RefferalDetails.ID == uint(claimingUserID) {
		return -1, fmt.Errorf("Not allowed to use this coupon")
	}
	CodeOwnerID := RefferalDetails.UserID
	return int(CodeOwnerID), nil
}

func (ru *refferalUseCase) ClaimRefferalBonus(claimingUserID, codeOwnerID int) error {

	CodeOwnerWallet, err := ru.orderUseCase.FindUserWallet(int(codeOwnerID))
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	} else if CodeOwnerWallet.ID == 0 {
		NewCodeOwnerWallet, err := ru.orderUseCase.InitializeNewWallet(int(codeOwnerID))
		if err != nil {
			return fmt.Errorf("Failed to initialize wallet for code owner: %s", err)
		} else if NewCodeOwnerWallet.ID == 0 {
			return fmt.Errorf("Failed to verify code owner new wallet")
		}

	}
	ClamingUserWallet, err := ru.orderUseCase.FindUserWallet(int(claimingUserID))
	if err != nil {
		return fmt.Errorf("Failed to find user wallet : %s", err)
	} else if ClamingUserWallet.ID == 0 {
		NewClamingUserWallet, err := ru.orderUseCase.InitializeNewWallet(claimingUserID)
		if err != nil {
			return fmt.Errorf("Failed to initialize wallet for code owner: %s", err)
		} else if NewClamingUserWallet.ID == 0 {
			return fmt.Errorf("Failed to verify code owner new wallet")
		}

	}

	CodeOwnerWallet, err = ru.orderUseCase.UpdateUserWallet(int(codeOwnerID), (CodeOwnerWallet.Amount + 50))
	if err != nil {
		return fmt.Errorf("Failed update bonus on  : %s", err)
	} else if CodeOwnerWallet.ID == 0 {
		return fmt.Errorf("Failed to verify bonus updated wallet")
	}

	ClamingUserWallet, err = ru.orderUseCase.UpdateUserWallet(int(claimingUserID), (ClamingUserWallet.Amount + 50))
	if err != nil {
		return fmt.Errorf("Failed update bonus  : %s", err)
	} else if CodeOwnerWallet.ID == 0 {
		return fmt.Errorf("Failed to verify bonus updated wallet")
	}
	return nil
}
