package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AdminRepository interface {
	FindAdminLoginCredentials() (config.AdminCredentials, error)
	GetAllUserDataFromDatabase() ([]response.UserData, error)
	BlockUserOnDatabase(id int) error
	UnBlockUserOnDatabase(id int) error
	FindUserByNameFromDatabase(name string) ([]response.UserData, error)
	SaveAdminOnDatabase(admin request.SignUpData) (response.UserData, error)
}
