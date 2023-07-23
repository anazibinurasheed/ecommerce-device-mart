package usecase

import (
	"errors"
	"fmt"
	"sync"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
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

var (
	users      = make(map[string]string)
	usersMutex sync.Mutex
)

func (cu *commonUseCase) ValidateSignupRequest(phone request.Phone) (int, error) {
	UserData, err := cu.userRepo.FindUserByPhone(phone.Phone)
	if err != nil {
		return 0, err
	}
	if UserData.Id != 0 {
		return 0, errors.New("User already exist with this phone number")
	}

	usersMutex.Lock()
	users[helper.GenerateUniqueID()] = fmt.Sprintf("%d", phone)
	usersMutex.Unlock()
	// number := strconv.Itoa(phone.Phone)
	// err = helper.SendOtp(number)
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to send otp%s", err)
	// }

	return phone.Phone, nil
}

// func (cu *commonUseCase) PhoneValidater(code string, c *gin.Context) (string, error) {
// 	phone, _ := helper.GetFromCookie("UserPhoneForCheckOtp", c)
// 	number := strconv.Itoa(phone)
// 	fmt.Printf("NUMBER %#v ,CODE %#v", number, code)
// 	status, err := otp.CheckOtp(number, code)

// 	if err != nil {
// 		return "", err
// 	} else if status == "incorrect" {
// 		return status, nil
// 	}

// 	tokenString, err := token.GenerateJwtToken(0)
// 	maxAge := int(time.Now().Add(time.Minute * 30).Unix())
// 	c.SetCookie("Verified", tokenString, maxAge, "", "", false, true)
// 	c.SetSameSite(http.SameSiteLaxMode)
// 	return status, nil

// }
