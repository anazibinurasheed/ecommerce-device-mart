package repository

import (
	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repository/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type refferalDatabase struct {
	DB *gorm.DB
}

func NewRefferalRepository(DB *gorm.DB) interfaces.ReferralRepository {
	return &refferalDatabase{
		DB: DB,
	}
}

func (rd *refferalDatabase) InsertNewReferralCode(userID int, refferalCode string) (response.Referral, error) {
	var InsertedDetails response.Referral
	query := `INSERT INTO referrals (user_id,code)VALUES($1,$2) RETURNING * ;`
	err := rd.DB.Raw(query, userID, refferalCode).Scan(&InsertedDetails).Error
	return InsertedDetails, err
}

func (rd *refferalDatabase) FindReferralCodeByCode(refferalCode string) (response.Referral, error) {
	var ReferrelDetails response.Referral
	query := `SELECT * FROM referrals WHERE code = $1 ;`
	err := rd.DB.Raw(query, refferalCode).Scan(&ReferrelDetails).Error
	return ReferrelDetails, err
}

func (rd *refferalDatabase) FindReferralCodeByUserID(userID int) (response.Referral, error) {
	var ReferrelDetails response.Referral
	query := `SELECT * FROM referrals WHERE user_id = $1 ;`
	err := rd.DB.Raw(query, userID).Scan(&ReferrelDetails).Error
	return ReferrelDetails, err
}
