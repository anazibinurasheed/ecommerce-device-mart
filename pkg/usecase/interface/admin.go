package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

// type AdminUseCase interface {
// 	SudoAdminLogin(sudoData request.SudoLoginData) error
// 	GetAllUserData() ([]response.UserData, error)
// 	BlockUserByID(ID int) error
// 	UnBlockUserByID(ID int) error
// 	AdminSignup(admin request.SignUpData) error
// }

type AdminUseCase interface {
	// SudoAdminLogin logs in as a super admin with sudo privileges.
	SudoAdminLogin(sudoData request.SudoLoginData) error

	// GetAllUserData retrieves a list of user data.
	GetAllUserData() ([]response.UserData, error)

	// BlockUserByID blocks a user by their ID.
	BlockUserByID(userID int) error

	// UnBlockUserByID unblocks a user by their ID.
	UnBlockUserByID(userID int) error

	// AdminSignUp signs up a new admin user.
	AdminSignUp(adminData request.SignUpData) error
}
