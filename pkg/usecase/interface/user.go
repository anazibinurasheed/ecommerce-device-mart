package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

// user service is collection of method signature in the usecase package .
// usecase have a struct that struct have repository/interfaces.UserRespository type variable ,with including that methods we
// created a new method , that method signature will hold in this  UserService interface .
type UserUseCase interface {
	SignUp(user request.SignUpData) error
	ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error)
	FindUserById(id int) (response.UserData, error)
	AddNewAddress(userId int, address request.Address) error
	DisplayListOfStates() ([]response.States, error)
	UpdateUserAddress(address request.Address, addressID int, userID int) error
	GetUserAddresses(userId int) ([]response.Address, error)
	DeleteUserAddress(addressId int) error
	GetProfile(userId int) (response.Profile, error)
	ForgotPassword(userId int, c *gin.Context) error
	ChangeUserPassword(password request.ChangePassword, userId int, c *gin.Context) error
	SetDefaultAddress(userID, addressID int) error
	CheckUserOldPassword(password request.OldPassword, userId int) error
	UpdateUserName(username string, userID int) error
}
