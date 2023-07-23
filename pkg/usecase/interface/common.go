package interfaces

import (
	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
)

type CommonUseCase interface {
	ValidateSignupRequest(phone request.Phone) (int, error)
	// PhoneValidater(code string, c *gin.Context) (string, error)
}
