package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type ReferralUseCase interface {
	GetUserReferralCode(userID int) (response.Referral, error)
	ClaimReferralBonus(claimingUserID, codeOwnerID int) error
	VerifyReferralCode(referralCode string, claimingUserID int) (int, error)
}
