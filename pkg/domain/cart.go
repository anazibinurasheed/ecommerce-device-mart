package domain

type Cart struct {
	ID        uint    `gorm:"not null;primaryKey"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Qty       int     `gorm:"not null"`
	// CouponID  int
}
