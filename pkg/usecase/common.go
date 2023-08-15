package usecase

import (
	"errors"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
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

// var (
// 	users      = make(map[string]string)
// 	usersMutex sync.Mutex
// )

func (cu *commonUseCase) ValidateSignupRequest(phone request.Phone) (int, error) {
	UserData, err := cu.userRepo.FindUserByPhone(phone.Phone)
	if err != nil {
		return 0, err
	}
	if UserData.Id != 0 {
		return 0, errors.New("User already exist with this phone number")
	}

	// usersMutex.Lock()
	// users[helper.GenerateUniqueID()] = fmt.Sprintf("%d", phone)
	// usersMutex.Unlock()

	// number := strconv.Itoa(phone.Phone)
	// err = helper.SendOtp(number)
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to send otp%s", err)
	// }

	return phone.Phone, nil
}
