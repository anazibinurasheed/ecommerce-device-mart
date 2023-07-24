package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type UserRepository interface {
	SaveUserOnDatabase(user request.SignUpData) error
	FindUserByPhone(phone int) (response.UserData, error)
	FindUserByEmail(email string) (response.UserData, error)
	FindUserById(id int) (response.UserData, error)
	ReadCategory() ([]response.Category, error)
	AddAdressToDatabase(userId int, address request.Address) (response.Address, error)
	GetListOfStates() ([]response.States, error)
	FindDefaultAddressById(userId int) (response.Address, error)
	SetIsDefaultStatusOnAddress(status bool, addressId int, userId int) (response.Address, error)
	GetAllUserAddresses(userId int) ([]response.Address, error)
	UpdateAddress(address request.Address, addressID int, userID int) (response.Address, error)
	DeleteAddressFromDatabase(adressId int) (response.Address, error)
	ChangePassword(userId int, newPassword string) error
	FindUserAddress(userID int) (response.Address, error)
	UpdateUserName(name string, userID int) (response.UserData, error)
	FindAddressByAddressID(addressID int) (response.Address, error)
}
