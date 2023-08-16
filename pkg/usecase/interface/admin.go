package interfaces

import (
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AdminUseCase interface {
	SudoAdminLogin(sudoData request.SudoLoginData) error
	GetAllUserData() ([]response.UserData, error)
	BlockUserByID(ID int) error
	UnBlockUserByID(ID int) error
	AdminSignup(admin request.SignUpData) error
}
