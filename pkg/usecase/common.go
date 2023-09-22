package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
)

type commonUseCase struct {
	userRepo  interfaces.UserRepository
	adminRepo interfaces.AdminRepository
}

func NewCommonUseCase(userRepo interfaces.UserRepository, adminRepo interfaces.AdminRepository) services.CommonUseCase {
	return &commonUseCase{
		userRepo:  userRepo,
		adminRepo: adminRepo,
	}

}

func (cu *commonUseCase) ValidateSignUpRequest(phone request.Phone) (int, error) {
	userData, err := cu.userRepo.FindUserByPhone(phone.Phone)
	if err != nil {
		return 0, err
	}
	if userData.ID != 0 {
		return 0, fmt.Errorf("User already exist with this phone number")
	}

	//below code is necessary commented out because defined a predefined otp

	// number := strconv.Itoa(phone.Phone)
	// err = helper.SendOtp(number)
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to send otp%s", err)
	// }

	return phone.Phone, nil
}
