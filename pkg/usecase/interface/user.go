package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	FindUserById(userID int) (response.UserData, error)
	AddNewAddress(userID int, address request.Address) error
	DisplayListOfStates() ([]response.States, error)
	UpdateUserAddress(address request.Address, addressID int, userID int) error
	GetUserAddresses(userID int) ([]response.Address, error)
	DeleteUserAddress(addressID int) error
	GetProfile(userID int) (response.Profile, error)
	ForgotPassword(userID int, c *gin.Context) error
	ChangeUserPassword(password request.ChangePassword, userID int, c *gin.Context) error
	SetDefaultAddress(userID, addressID int) error
	CheckUserOldPassword(password request.OldPassword, userID int) error
	UpdateUserName(username string, userID int) error
}
