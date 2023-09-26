package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AdminUseCase interface {

	// GetAllUserData retrieves a list of user data.
	GetAllUserData() ([]response.UserData, error)

	// BlockUserByID blocks a user by their ID.
	BlockUserByID(userID int) error

	// UnBlockUserByID unblocks a user by their ID.
	UnBlockUserByID(userID int) error
}
