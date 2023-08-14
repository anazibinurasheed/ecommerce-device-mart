package usecase

import (
	"errors"
	"fmt"
	"log"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
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

func (u *userUseCase) SignUp(user request.SignUpData) error {
	userData, err := u.userRepo.FindUserByPhone(user.Phone)
	if err != nil {
		return fmt.Errorf("failed to find user by phone :%s", err)
	}
	if userData.Id != 0 {
		return fmt.Errorf("user already exist with this phone number")
	}

	userData, err = u.userRepo.FindUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("failed to find user by email :%s", err)
	}
	if userData.Id != 0 {
		return fmt.Errorf("user already exist with this email address")
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// if err != nil {
	// 	return fmt.Errorf("failed to generate hash from password :%s", err)
	// }

	// user.Password = string(hashedPassword)

	if _, err := u.userRepo.SaveUserOnDatabase(user); err != nil {
		return fmt.Errorf("failed to save user on db, user sign up failed :%s", err)
	}

	return nil
}

func (u *userUseCase) ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error) {
	userData, err := u.userRepo.FindUserByPhone(user.Phone)
	if err != nil {
		return response.UserData{}, err
	} else if userData.Id == 0 {
		return response.UserData{}, errors.New("user dont have an account")
	} else if userData.IsBlocked {
		return response.UserData{}, errors.New("user have been blocked")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return response.UserData{}, errors.New("incorrect password")
	}

	userData.Password = ""
	return userData, nil

}

func (u *userUseCase) FindUserById(id int) (response.UserData, error) {
	UserData, err := u.userRepo.FindUserById(id)
	if err != nil {
		return response.UserData{}, err
	}
	return UserData, nil
}

func (u *userUseCase) ReadAllCategories() ([]response.Category, error) {
	ListOfAllCategories, err := u.userRepo.ReadCategory()
	if err != nil {
		return nil, err
	}
	return ListOfAllCategories, nil

}

func (u *userUseCase) DisplayListOfStates() ([]response.States, error) {
	ListOfStates, err := u.userRepo.GetListOfStates()

	if err != nil {
		return nil, err
	}
	if len(ListOfStates) == 0 {
		return nil, fmt.Errorf("no states found")
	}

	return ListOfStates, nil
}

func (u *userUseCase) AddNewAddress(userID int, address request.Address) error {
	createdAddress, err := u.userRepo.AddAdressToDatabase(userID, address)
	if err != nil {
		return fmt.Errorf("failed to add address to db :%s", err)
	}

	if createdAddress.ID == 0 {
		return fmt.Errorf("failed to verify created address")

	}

	defaultAddress, err := u.userRepo.FindDefaultAddressById(userID)
	if err != nil {
		return err
	}

	if defaultAddress.ID == 0 {

		setDefaultAddress, err := u.userRepo.SetIsDefaultStatusOnAddress(true, int(createdAddress.ID), userID)
		if err != nil {
			return fmt.Errorf("failed to set default address :%s", err)
		}

		if setDefaultAddress.ID == 0 {
			return fmt.Errorf("failed to set default address as new address")
		}
	}
	return nil
}

func (u *userUseCase) FindDefaultAddress(userId int) (response.Address, error) {
	DefaultAddress, err := u.userRepo.FindDefaultAddressById(userId)

	if err != nil {
		return response.Address{}, fmt.Errorf("failed to find default address :%s ", err)
	}

	if DefaultAddress.ID == 0 {
		return response.Address{}, fmt.Errorf("failed to verify retrieved default address")
	}
	return DefaultAddress, nil
}

func (u *userUseCase) UpdateUserAddress(address request.Address, addressID int, userID int) error {
	UpdatedAddress, err := u.userRepo.UpdateAddress(address, addressID, userID)

	if err != nil {
		return err
	}
	if UpdatedAddress.ID == 0 {
		return errors.New("failed to update address")
	}
	return nil

}

func (u *userUseCase) GetUserAddresses(userId int) ([]response.Address, error) {
	ListOfAddresses, err := u.userRepo.GetAllUserAddresses(userId)

	if err != nil {
		return nil, err
	}
	return ListOfAddresses, nil
}

func (u *userUseCase) DeleteUserAddress(addressId int) error {
	DeletedAddress, err := u.userRepo.DeleteAddressFromDatabase(addressId)
	if err != nil {
		return err
	}
	if DeletedAddress.ID == 0 {
		return errors.New("Address not found.")
	}
	userID := DeletedAddress.UserID
	UserAddress, err := u.userRepo.FindUserAddress(int(userID))
	if err != nil {
		return fmt.Errorf("Failed to set default address %s", err)
	}
	if UserAddress.ID != 0 {
		_, err := u.userRepo.SetIsDefaultStatusOnAddress(true, int(UserAddress.ID), int(UserAddress.UserID))
		if err != nil {
			return fmt.Errorf("Failed to update default address :%s", err)
		}
	}
	return nil
}
func (u *userUseCase) SetDefaultAddress(userID, addressID int) error {
	DefaultAddress, err := u.userRepo.FindDefaultAddressById(userID)
	if err != nil {
		return fmt.Errorf("Failed to find default address :%s", err)
	}
	Address, err := u.userRepo.SetIsDefaultStatusOnAddress(false, int(DefaultAddress.ID), userID)
	if Address.ID == 0 || Address.IsDefault || err != nil {
		return fmt.Errorf("Failed to change default address : %s", err)
	}
	NewDefaultAddress, err := u.userRepo.SetIsDefaultStatusOnAddress(true, addressID, userID)
	if err != nil || !NewDefaultAddress.IsDefault {
		return fmt.Errorf("Failed to set  address to default : %s", err)
	}
	return nil
}
func (u *userUseCase) GetProfile(userId int) (response.Profile, error) {
	var profile response.Profile

	UserData, err := u.userRepo.FindUserById(userId)
	if err != nil {
		return response.Profile{}, fmt.Errorf("Failed to find user : %s", err)
	}

	if UserData.Id == 0 {
		return response.Profile{}, errors.New("User not found .")
	}

	profile.Id = UserData.Id
	profile.UserName = UserData.UserName
	profile.Email = UserData.Email
	profile.Phone = UserData.Phone
	profile.Addresses, err = u.userRepo.GetAllUserAddresses(userId)
	if err != nil {
		return response.Profile{}, fmt.Errorf("Failed to find user addresses:  %s", err)
	}

	return profile, nil
}

// profile
func (u *userUseCase) ForgotPassword(userid int, c *gin.Context) error {
	UserData, err := u.userRepo.FindUserById(userid)
	if err != nil {
		return err
	}

	if UserData.Id == 0 {
		return fmt.Errorf("User not found.")
	}

	err = helper.SendOtp(fmt.Sprint(UserData.Phone))
	if err != nil {
		return fmt.Errorf("Failed to send otp")
	}
	helper.SetToCookie(userid, "PasswordChange", c)

	return nil
}

func (u *userUseCase) ChangeUserPassword(password request.ChangePassword, userId int, c *gin.Context) error {

	if password.NewPassword != password.ReNewPassword {
		return errors.New("Password is not matching")
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), 10)
	if err != nil {
		log.Println("FAILED TO HASH PASSWORD", err)
		return errors.New("Unable to process the request ")
	}

	err = u.userRepo.ChangePassword(userId, string(HashedPassword))
	if err != nil {
		return fmt.Errorf("Failed to change password %s", err)
	}
	return nil
}

func (u *userUseCase) CheckUserOldPassword(password request.OldPassword, userId int) error {

	UserData, err := u.userRepo.FindUserById(userId)

	err = bcrypt.CompareHashAndPassword([]byte(UserData.Password), []byte(password.Password))
	if err != nil {
		return fmt.Errorf("Entered wrong password : %s", err)
	}

	return nil
}

func (u *userUseCase) UpdateUserName(username string, userID int) error {

	UserData, err := u.userRepo.UpdateUserName(username, userID)
	if err != nil {

		return fmt.Errorf("Failed to update username : %s", err)
	}

	if UserData.UserName != username {
		return fmt.Errorf("Failed to update username")
	}
	return nil
}
