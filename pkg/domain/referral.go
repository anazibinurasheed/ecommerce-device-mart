package domain

type Referral struct {
	ID     uint   `gorm:"not null,unique,primaryKey"`
	UserID uint   `gorm:"not null,unique"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Code   string `gorm:"not null,unique"`
}
