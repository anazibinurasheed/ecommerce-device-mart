package interfaces

import (
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
)

type CommonUseCase interface {
	ValidateSignUpRequest(phone request.Phone) (int, error)
}
