package domain

// gorm is creating tables with underscore when the capital encounter .
type Category struct {
	ID            uint   ` gorm:"primaryKey;AutoIncrement;unique"`
	Category_Name string ` gorm:"unique;not null"`
	Images        JSONB
	IsBlocked     bool `gorm:"default:false"`
}

type Product struct {
	ID                 uint     `gorm:"primaryKey;unique;autoIncrement;not null"`
	CategoryID         uint     `gorm:"not null"`
	Category           Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Brand              string
	Price              int    `gorm:"not null"`
	SKU                string `gorm:"not null"`
	ProductName        string `gorm:"not null"`
	ProductDescription string `gorm:"not null"`
	Images             JSONB
	IsBlocked          bool `gorm:"default:false"`
}

type Rating struct {
	ID          uint   `gorm:"primaryKey;unique;autoIncrement;not null"`
	UserID      int    `gorm:"not null"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Rating      int    `gorm:"not null"`
	ProductID   int    `gorm:"not null"`
	Description string `gorm:"not null"`
}

type Wishlist struct {
	ID        uint    `json:"id"`
	UserID    int     `gorm:"not null"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID int     `gorm:"not null"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
