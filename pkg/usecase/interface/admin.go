package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AdminUseCase interface {
	SudoAdminLogin(sudoData request.LoginSudoAdmin) error
	GetAllUserData() ([]response.UserData, error)
	BlockUserById(id int) error
	UnBlockUserById(id int) error
	AdminSignup(admin request.SignUpData) error
}
