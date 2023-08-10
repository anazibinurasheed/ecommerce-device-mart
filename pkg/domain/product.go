package domain

// gorm is creating tables with underscore when the capital encounter .
type Category struct {
	ID            uint   ` gorm:"primaryKey;AutoIncrement;unique"`
	Category_Name string ` gorm:"unique;not null"`
	IsBlocked     bool
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
	ProductImage       string
	IsBlocked          bool `gorm:"default:false"`
}

type Rating struct {
	ID          uint    `gorm:"primaryKey;unique;autoIncrement;not null"`
	UserID      int     `gorm:"not null"`
	User        User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID   int     `gorm:"not null"`
	Product     Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Rating      int     `gorm:"not null"`
	Description string  `gorm:"not null"`
}

// type ProductItem struct {
// 	ID            uint    `gorm:"primaryKey;autoIncrement"`
// 	ProductID     uint    `gorm:"not null"`
// 	Product       Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// 	SKU           string  `gorm:"not null"`
// 	Product_Image string
// 	Price         string `gorm:"not null"`
// }

// type Variation struct {
// 	ID         uint     `gorm:"primaryKey;autoIncrement"`
// 	CategoryID uint     `gorm:"not null"`
// 	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// 	Name       string   ` gorm:"not null"`
// }

// type VariationOption struct {
// 	ID          uint      ` gorm:"primaryKey;autoIncrement"`
// 	VariationID uint      ` gorm:"not null"`
// 	Variation   Variation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// 	Value       string    ` gorm:"not null"`
// }

// type ProductConfiguration struct {
// 	ProductItemID     uint            `gorm:"not null"`
// 	ProductItem       ProductItem     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// 	VariationOptionID uint            `gorm:"not null"`
// 	VariationOption   VariationOption `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// }
