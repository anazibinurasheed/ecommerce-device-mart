package domain

type Wallet struct {
	ID     uint    `gorm:"primaryKey,unique,not null"`
	UserID int     `gorm:"not null"`
	User   User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount float32 `gorm:"default:0"`
}
