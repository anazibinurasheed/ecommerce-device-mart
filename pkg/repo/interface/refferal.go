package interfaces

import "github.com/anazibinurasheed/project-device-mart/pkg/util/response"

type ReferralRepository interface {
	InsertNewReferralCode(userID int, referralCode string) (response.Referral, error)
	FindReferralCodeByCode(referralCode string) (response.Referral, error)
	FindReferralCodeByUserID(userID int) (response.Referral, error)
}
