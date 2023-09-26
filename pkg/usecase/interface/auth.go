package interfaces

import (
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
)

type AuthUseCase interface {
	ValidateSignUpRequest(phone request.Phone) (int, error)
	SignUp(user request.SignUpData) error
	ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error)
}
