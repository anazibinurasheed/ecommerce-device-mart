package usecase

import (
	"fmt"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
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
