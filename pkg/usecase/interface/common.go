package interfaces

import (
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
)

type CommonUseCase interface {
	ValidateSignupRequest(phone request.Phone) (int, error)
}
