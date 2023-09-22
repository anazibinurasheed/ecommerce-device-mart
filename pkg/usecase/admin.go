package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepo interfaces.AdminRepository
	userRepo  interfaces.UserRepository
}

func NewAdminUseCase(adminUseCase interfaces.AdminRepository, userUseCase interfaces.UserRepository) services.AdminUseCase {
	return &adminUsecase{
		adminRepo: adminUseCase,
		userRepo:  userUseCase,
	}
}

func (ac *adminUsecase) AdminSignUp(admin request.SignUpData) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return fmt.Errorf("Failed to generate hash from password :%s", err)
	}

	admin.Password = string(hashedPassword)
	adminData, err := ac.adminRepo.CreateAdmin(admin)

	if err != nil {
		return fmt.Errorf("Failed to create user :%s", err)
	}
	if adminData.ID == 0 {
		return fmt.Errorf("Failed to verify created user")
	}

	return nil
}

func (ac *adminUsecase) SudoAdminLogin(sudoData request.SudoLoginData) error {
	adminCredentials, err := ac.adminRepo.FindAdminCredentials()
	if err != nil {
		return err
	}
	if sudoData.Username == "" || sudoData.Password == "" {
		return fmt.Errorf("Credentials is empty")

	} else if adminCredentials.AdminUsername == sudoData.Username && adminCredentials.AdminPassword == sudoData.Password {
		return nil

	} else {
		return fmt.Errorf("Invalid credentials")
	}
}

func (ac *adminUsecase) GetAllUserData() ([]response.UserData, error) {

	listOfAllUserData, err := ac.adminRepo.FetchAllUserData()
	if err != nil {
		return []response.UserData{}, fmt.Errorf("Failed to get user data's :%s", err)
	}

	return listOfAllUserData, nil

}

func (ac *adminUsecase) BlockUserByID(ID int) error {
	err := ac.adminRepo.BlockUserByID(ID)
	if err != nil {
		return err
	}
	return nil
}

func (ac *adminUsecase) UnBlockUserByID(ID int) error {
	err := ac.adminRepo.UnblockUserByID(ID)
	if err != nil {
		return err
	}
	return nil
}

func (ac *adminUsecase) FindUsersByName(name string) ([]response.UserData, error) {
	user, err := ac.adminRepo.FindUsersByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
