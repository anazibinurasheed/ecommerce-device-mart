package domain

import "time"

type Coupon struct {
	ID                uint      `gorm:"primaryKey,unique,not null"`
	Code              string    `gorm:"unique,not null"`
	CouponName        string    `gorm:"not null"`
	MinOrderValue     float64   `gorm:"not null"`
	DiscountPercent   float64   `gorm:"not null"`
	DiscountMaxAmount float64   `gorm:"not null"`
	ValidFrom         time.Time `gorm:"not null"`
	ValidTill         time.Time `gorm:"not null"`
	ValidDays         int       `gorm:"not null"`
	IsBlocked         bool      `gorm:"default:false"`
}

type CouponTracking struct {
	ID       uint   `gorm:"primaryKey,unique,not null"`
	CouponID int    `gorm:"not null"`
	Coupon   Coupon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsUsed   bool   `gorm:"default:false"`
}
