package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
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

func (ac *adminUsecase) AdminSignup(admin request.SignUpData) error {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return fmt.Errorf("Failed to generate hash from password :%s", err)
	}

	admin.Password = string(HashedPassword)
	AdminData, err := ac.adminRepo.SaveAdminOnDatabase(admin)

	if err != nil {
		return fmt.Errorf("Failed to create user :%s", err)
	}
	if AdminData.ID == 0 {
		return fmt.Errorf("Failed to verify created user")
	}

	return nil
}

func (ac *adminUsecase) SudoAdminLogin(sudoData request.LoginSudoAdmin) error {
	adminCredentials, err := ac.adminRepo.FindAdminLoginCredentials()
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

	ListOfAllUserData, err := ac.adminRepo.GetAllUserDataFromDatabase()
	if err != nil {
		return []response.UserData{}, fmt.Errorf("Failed to get user data's :%s", err)
	}

	return ListOfAllUserData, nil

}
func (ac *adminUsecase) BlockUserById(id int) error {
	err := ac.adminRepo.BlockUserOnDatabase(id)
	if err != nil {
		return err
	}
	return nil
}

func (ac *adminUsecase) UnBlockUserById(id int) error {
	err := ac.adminRepo.UnBlockUserOnDatabase(id)
	if err != nil {
		return err
	}
	return nil
}

func (ac *adminUsecase) FindUsersByName(name string) ([]response.UserData, error) {
	user, err := ac.adminRepo.FindUserByNameFromDatabase(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
