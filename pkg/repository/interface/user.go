package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

//	type UserRepository interface {
//		SaveUserOnDatabase(user request.SignUpData) (response.UserData, error)
//		FindUserByPhone(phone int) (response.UserData, error)
//		FindUserByEmail(email string) (response.UserData, error)
//		FindUserById(id int) (response.UserData, error)
//		ReadCategory() ([]response.Category, error)
//		AddAdressToDatabase(userId int, address request.Address) (response.Address, error)
//		GetListOfStates() ([]response.States, error)
//		FindDefaultAddressById(userId int) (response.Address, error)
//		SetIsDefaultStatusOnAddress(status bool, addressId int, userId int) (response.Address, error)
//		GetAllUserAddresses(userId int) ([]response.Address, error)
//		UpdateAddress(address request.Address, addressID int, userID int) (response.Address, error)
//		DeleteAddressFromDatabase(adressId int) (response.Address, error)
//		ChangePassword(userId int, newPassword string) error
//		FindUserAddress(userID int) (response.Address, error)
//		UpdateUserName(name string, userID int) (response.UserData, error)
//		FindAddressByAddressID(addressID int) (response.Address, error)
//	}
type UserRepository interface {
	CreateUser(user request.SignUpData) (response.UserData, error)
	FindUserByPhone(phone int) (response.UserData, error)
	FindUserByEmail(email string) (response.UserData, error)
	FindUserByID(id int) (response.UserData, error)
	ReadCategories() ([]response.Category, error)
	AddAddress(userID int, address request.Address) (response.Address, error)
	GetListOfStates() ([]response.States, error)
	FindDefaultAddress(userID int) (response.Address, error)
	SetDefaultAddressStatus(status bool, addressID, userID int) (response.Address, error)
	GetAllUserAddresses(userID int) ([]response.Address, error)
	UpdateAddress(address request.Address, addressID, userID int) (response.Address, error)
	DeleteAddress(addressID int) (response.Address, error)
	ChangePassword(userID int, newPassword string) error
	FindUserAddress(userID int) (response.Address, error)
	UpdateUserName(name string, userID int) (response.UserData, error)
	FindAddressByID(addressID int) (response.Address, error)
}
