package interfaces

import (
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AuthUseCase interface {
	AdminLogin(sudoData request.AdminLogin) error
	SignUp(user request.SignUpData) error
	ValidateSignUpRequest(phone request.Phone) (int, error)
	ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error)
}
