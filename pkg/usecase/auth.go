package usecase

import (
	"errors"
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"golang.org/x/crypto/bcrypt"
)

var (
	InvalidCredentials = errors.New("Invalid credentials")
)

type authUseCase struct {
	userRepo  interfaces.UserRepository
	adminRepo interfaces.AdminRepository
}

func NewCommonUseCase(userRepo interfaces.UserRepository, adminRepo interfaces.AdminRepository) services.AuthUseCase {
	return &authUseCase{
		userRepo:  userRepo,
		adminRepo: adminRepo,
	}

}

func (ac *authUseCase) AdminLogin(sudoData request.AdminLogin) error {
	adminCredentials, err := ac.adminRepo.FindAdminCredentials()
	if err != nil {
		return err
	}
	if sudoData.Username == "" || sudoData.Password == "" {
		return fmt.Errorf("Credentials is empty")

	} else if adminCredentials.AdminUsername == sudoData.Username && adminCredentials.AdminPassword == sudoData.Password {
		return nil

	} else {
		return InvalidCredentials
	}
}

func (cu *authUseCase) ValidateSignUpRequest(phone request.Phone) (int, error) {
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

func (u *authUseCase) SignUp(user request.SignUpData) error {
	userData, err := u.userRepo.FindUserByPhone(user.Phone)
	if err != nil {
		return fmt.Errorf("Failed to find user by phone :%s", err)
	}
	if userData.ID != 0 {
		return fmt.Errorf("User already exist with this phone number")
	}

	userData, err = u.userRepo.FindUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("Failed to find user by email :%s", err)
	}
	if userData.ID != 0 {
		return fmt.Errorf("User already exist with this email address")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return fmt.Errorf("failed to generate hash from password :%s", err)
	}

	user.Password = string(hashedPassword)

	if _, err := u.userRepo.CreateUser(user); err != nil {
		return fmt.Errorf("Failed to save user on db, user sign up failed :%s", err)
	}

	return nil
}

func (u *authUseCase) ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error) {
	userData, err := u.userRepo.FindUserByPhone(user.Phone)
	if err != nil {
		return response.UserData{}, err
	}

	if userData.ID == 0 {
		return response.UserData{}, fmt.Errorf("User don't have an account")
	}

	if userData.IsBlocked {
		return response.UserData{}, fmt.Errorf("User has been blocked")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return response.UserData{}, fmt.Errorf("incorrect password")
	}

	userData.Password = ""
	return userData, nil

}
