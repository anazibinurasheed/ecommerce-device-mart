package interfaces

import (
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AuthUseCase interface {
	// SudoAdminLogin logs in as a super admin with sudo privileges.
	SudoAdminLogin(sudoData request.SudoLoginData) error

	ValidateSignUpRequest(phone request.Phone) (int, error)
	SignUp(user request.SignUpData) error
	ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error)
}
