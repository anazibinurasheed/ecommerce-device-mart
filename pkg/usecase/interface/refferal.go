package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type RefferalUseCase interface {
	GetUserRefferalCode(userID int) (response.Referral, error)
	ClaimRefferalBonus(claimingUserID, codeOwnerID int) error
	VerifyRefferalCode(refferalCode string, claimingUserID int) (int, error)
}
