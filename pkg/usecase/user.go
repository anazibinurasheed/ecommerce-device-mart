package usecase

import (
	"errors"
	"fmt"
	"log"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	services "github.com/anazibinurasheed/project-device-mart/pkg/usecase/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) FindUserById(userID int) (response.UserData, error) {
	userData, err := u.userRepo.FindUserByID(userID)
	if err != nil {
		return response.UserData{}, err
	}
	return userData, nil
}

func (u *userUseCase) ReadAllCategories() ([]response.Category, error) {
	listOfAllCategories, err := u.userRepo.ReadCategories()
	if err != nil {
		return nil, err
	}
	return listOfAllCategories, nil

}

func (u *userUseCase) DisplayListOfStates() ([]response.States, error) {
	listOfStates, err := u.userRepo.GetListOfStates()
	if err != nil {
		return nil, err
	}
	if len(listOfStates) == 0 {
		return nil, fmt.Errorf("No states found")
	}

	return listOfStates, nil
}

func (u *userUseCase) AddNewAddress(userID int, address request.Address) error {
	createdAddress, err := u.userRepo.AddAddress(userID, address)
	if err != nil {
		return fmt.Errorf("Failed to add address to db :%s", err)
	}

	if createdAddress.ID == 0 {
		return fmt.Errorf("Failed to verify created address")

	}

	defaultAddress, err := u.userRepo.FindDefaultAddress(userID)
	if err != nil {
		return err
	}

	if defaultAddress.ID == 0 {

		setDefaultAddress, err := u.userRepo.SetDefaultAddressStatus(true, int(createdAddress.ID), userID)
		if err != nil {
			return fmt.Errorf("Failed to set default address :%s", err)
		}

		if setDefaultAddress.ID == 0 {
			return fmt.Errorf("Failed to set default address as new address")
		}
	}
	return nil
}

func (u *userUseCase) FindDefaultAddress(userID int) (response.Address, error) {
	defaultAddress, err := u.userRepo.FindDefaultAddress(userID)

	if err != nil {
		return response.Address{}, fmt.Errorf("Failed to find default address :%s ", err)
	}

	if defaultAddress.ID == 0 {
		return response.Address{}, fmt.Errorf("Failed to verify retrieved default address")
	}
	return defaultAddress, nil
}

func (u *userUseCase) UpdateUserAddress(address request.Address, addressID int, userID int) error {
	updatedAddress, err := u.userRepo.UpdateAddress(address, addressID, userID)

	if err != nil {
		return err
	}
	if updatedAddress.ID == 0 {
		return errors.New("Failed to update address")
	}
	return nil

}

func (u *userUseCase) GetUserAddresses(userID int) ([]response.Address, error) {
	listOfAddresses, err := u.userRepo.GetAllUserAddresses(userID)

	if err != nil {
		return nil, err
	}
	return listOfAddresses, nil
}

func (u *userUseCase) DeleteUserAddress(addressID int) error {
	deletedAddress, err := u.userRepo.DeleteAddress(addressID)
	if err != nil {
		return err
	}
	if deletedAddress.ID == 0 {
		return errors.New("Address not found.")
	}

	userID := deletedAddress.UserID

	userAddress, err := u.userRepo.FindUserAddress(int(userID))
	if err != nil {
		return fmt.Errorf("Failed to set default address %s", err)
	}
	if userAddress.ID != 0 {
		_, err := u.userRepo.SetDefaultAddressStatus(true, int(userAddress.ID), int(userAddress.UserID))
		if err != nil {
			return fmt.Errorf("Failed to update default address :%s", err)
		}
	}
	return nil
}

func (u *userUseCase) SetDefaultAddress(userID, addressID int) error {
	defaultAddress, err := u.userRepo.FindDefaultAddress(userID)
	if err != nil {
		return fmt.Errorf("Failed to find default address :%s", err)
	}

	address, err := u.userRepo.SetDefaultAddressStatus(false, int(defaultAddress.ID), userID)
	if address.ID == 0 || address.IsDefault || err != nil {
		return fmt.Errorf("Failed to change default address : %s", err)
	}

	newDefaultAddress, err := u.userRepo.SetDefaultAddressStatus(true, addressID, userID)
	if err != nil || !newDefaultAddress.IsDefault {
		return fmt.Errorf("Failed to set  address to default : %s", err)
	}
	return nil
}

func (u *userUseCase) GetProfile(userID int) (response.Profile, error) {
	userData, err := u.userRepo.FindUserByID(userID)
	if err != nil {
		return response.Profile{}, fmt.Errorf("Failed to find user : %s", err)
	}

	if userData.ID == 0 {
		return response.Profile{}, fmt.Errorf("User not found")
	}

	userAddresses, err := u.userRepo.GetAllUserAddresses(userID)
	if err != nil {
		return response.Profile{}, fmt.Errorf("Failed to find user addresses :%s", err)
	}

	return response.Profile{
		ID:        userData.ID,
		UserName:  userData.UserName,
		Email:     userData.Email,
		Phone:     userData.Phone,
		Addresses: userAddresses,
	}, nil
}

func (u *userUseCase) ForgotPassword(userID int, c *gin.Context) error {
	userData, err := u.userRepo.FindUserByID(userID)
	if err != nil {
		return err
	}

	if userData.ID == 0 {
		return fmt.Errorf("User not found.")
	}

	err = helper.SendOtp(fmt.Sprint(userData.Phone))
	if err != nil {
		return fmt.Errorf("Failed to send otp")
	}

	helper.SetToCookie(userID, "PasswordChange", c)

	return nil
}

func (u *userUseCase) ChangeUserPassword(password request.ChangePassword, userID int, c *gin.Context) error {
	if password.NewPassword != password.ReNewPassword {
		return fmt.Errorf("Password is not matching")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), 10)
	if err != nil {
		log.Println("FAILED TO HASH PASSWORD", err)
		return fmt.Errorf("Unable to process the request ")
	}

	err = u.userRepo.ChangePassword(userID, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("Failed to change password :%s", err)
	}
	return nil
}

func (u *userUseCase) CheckUserOldPassword(password request.OldPassword, userID int) error {
	userData, err := u.userRepo.FindUserByID(userID)

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password.Password))
	if err != nil {
		return fmt.Errorf("Entered wrong password : %s", err)
	}

	return nil
}

func (u *userUseCase) UpdateUserName(username string, userID int) error {
	userData, err := u.userRepo.UpdateUserName(username, userID)
	if err != nil {
		return fmt.Errorf("Failed to update username : %s", err)
	}
	if userData.UserName != username {
		return fmt.Errorf("Failed to update username")
	}
	return nil
}
