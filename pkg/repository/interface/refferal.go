package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type RefferalRepository interface {
	InsertNewRefferalCode(userID int, refferalCode string) (response.Referral, error)
	FindRefferalCodeByCode(refferalCode string) (response.Referral, error)
	FindRefferalCodeByUserId(userID int) (response.Referral, error)
}
