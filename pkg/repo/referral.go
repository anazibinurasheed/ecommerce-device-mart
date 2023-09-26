package repo

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type referralDatabase struct {
	DB *gorm.DB
}

func NewReferralRepository(DB *gorm.DB) interfaces.ReferralRepository {
	return &referralDatabase{
		DB: DB,
	}
}

func (rd *referralDatabase) InsertNewReferralCode(userID int, referralCode string) (response.Referral, error) {
	var insertedDetails response.Referral
	query := `INSERT INTO referrals (user_id,code)VALUES($1,$2) RETURNING * ;`
	err := rd.DB.Raw(query, userID, referralCode).Scan(&insertedDetails).Error
	return insertedDetails, err
}

func (rd *referralDatabase) FindReferralCodeByCode(referralCode string) (response.Referral, error) {
	var referralDetails response.Referral
	query := `SELECT * FROM referrals WHERE code = $1 ;`
	err := rd.DB.Raw(query, referralCode).Scan(&referralDetails).Error
	return referralDetails, err
}

func (rd *referralDatabase) FindReferralCodeByUserID(userID int) (response.Referral, error) {
	var referralDetails response.Referral
	query := `SELECT * FROM referrals WHERE user_id = $1 ;`
	err := rd.DB.Raw(query, userID).Scan(&referralDetails).Error
	return referralDetails, err
}
