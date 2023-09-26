package domain

import "time"

type Wallet struct {
	ID     uint    `gorm:"primaryKey,unique,not null"`
	UserID int     `gorm:"not null"`
	User   User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount float32 `gorm:"default:0"`
}

type WalletTransactionHistory struct {
	ID              uint      `gorm:"primaryKey,unique,not null"`
	TransactionTime time.Time `gorm:"not null"`
	UserID          int       `gorm:"not null"`
	User            User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount          float32   `gorm:"not null"`
	TransactionType string    `gorm:"not null"` // "credit" or "debit"
}
