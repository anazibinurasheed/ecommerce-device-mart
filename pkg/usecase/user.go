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
		return err
	}

	if userData.Id != 0 {
		return errors.New("User already exist with this phone number")
	}
	userData, err = u.userRepo.FindUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if userData.Id != 0 {
		return errors.New("User already exist with this email address")
	}
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("FAILED TO HASH PASSWORD", err)
		return errors.New("Unable to process the request ")
	}
	user.Password = string(HashedPassword)
	if _, err := u.userRepo.SaveUserOnDatabase(user); err != nil {
		log.Println("FAILED TO SAVE USER ON DATABASE")
		return err
	}

	return nil
}

func (u *userUseCase) ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error) {
	UserData, err := u.userRepo.FindUserByPhone(user.Phone)
	if err != nil {
		return response.UserData{}, err
	} else if UserData.Id == 0 {
		return response.UserData{}, errors.New("user dont have an account.")
	} else if UserData.IsBlocked {
		return response.UserData{}, errors.New("user have been blocked.")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(UserData.Password), []byte(user.Password)); err != nil {
		return response.UserData{}, errors.New("Incorrect password.")
	}

	return UserData, nil

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
		log.Println("  FAILED WHILE READING ALL CATEGORIES ")
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
		return nil, errors.New("No states found")
	}

	return ListOfStates, nil
}

func (u *userUseCase) AddNewAddress(userId int, address request.Address) error {
	CreatedAddress, err := u.userRepo.AddAdressToDatabase(userId, address)

	if err != nil {
		return err
	}

	if CreatedAddress.ID == 0 {
		return errors.New("Failed to create new address")

	}

	DefaultAddress, err := u.userRepo.FindDefaultAddressById(userId)
	if err != nil {
		return err
	}

	if DefaultAddress.ID == 0 {
		SetDefaultAddress, err := u.userRepo.SetIsDefaultStatusOnAddress(true, int(CreatedAddress.ID), userId)
		if err != nil {
			return err
		}

		if SetDefaultAddress.ID == 0 {
			return errors.New("Failed to set default address as new address ")
		}
	}
	return nil
}

func (u *userUseCase) FindDefaultAddress(userId int) (response.Address, error) {
	DefaultAddress, err := u.userRepo.FindDefaultAddressById(userId)

	if err != nil {
		return response.Address{}, err
	}

	if DefaultAddress.ID == 0 {
		return response.Address{}, errors.New("Failed to find default address")
	}
	return DefaultAddress, nil
}

// func (u *userUseCase) GetAllUserAddresses(userId int) ([]response.Address, error) {
// 	DefaultAddress, err := u.userRepository.GetAllUserAddresses(userId)
// 	//double check
// 	if err != nil {
// 		return nil, err
// 	}

// 	return DefaultAddress, nil

// }

func (u *userUseCase) UpdateUserAddress(address request.Address, addressID int, userID int) error {

	UpdatedAddress, err := u.userRepo.UpdateAddress(address, addressID, userID)

	if err != nil {
		return err
	}
	if UpdatedAddress.ID == 0 {
		return errors.New("Failed to update address")
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
	for key, val := range profile.Addresses {
		fmt.Println("UserAddress ", key, "   :", val)
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

	// _, err1 := c.Cookie("PasswordChange")
	// if err1 != nil {

	// 	return errors.New("Unauthorized api call")

	// }

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
	// helper.DeleteCookie("PasswordChange", c)
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
		return fmt.Errorf("Failed to update username ")
	}
	return nil
}
