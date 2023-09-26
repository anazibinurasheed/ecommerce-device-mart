package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

// type AdminRepository interface {
// 	FindAdminLoginCredentials() (config.AdminCredentials, error)
// 	GetAllUserData() ([]response.UserData, error)
// 	BlockUser(ID int) error
// 	UnBlockUser(ID int) error
// 	FindUserByNameFromDatabase(name string) ([]response.UserData, error)
// 	SaveAdminOnDatabase(admin request.SignUpData) (response.UserData, error)
// }

type AdminRepository interface {
	FindAdminCredentials() (config.AdminCredentials, error)
	FetchAllUserData() ([]response.UserData, error)
	BlockUserByID(userID int) error
	UnblockUserByID(userID int) error
	FindUsersByName(name string) ([]response.UserData, error)
}
