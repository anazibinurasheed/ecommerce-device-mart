package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// gorm is creating tables with underscore when the capital encounter .
type Category struct {
	ID            uint   ` gorm:"primaryKey;AutoIncrement;unique"`
	Category_Name string ` gorm:"unique;not null"`
	Images        JSONB
	IsBlocked     bool `gorm:"default:false"`
}

// type Product struct {
// 	ID                 uint     `gorm:"primaryKey;unique;autoIncrement;not null"`
// 	CategoryID         uint     `gorm:"not null"`
// 	Category           Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	Brand              string
// 	Price              int    `gorm:"not null"`
// 	SKU                string `gorm:"not null"`
// 	ProductName        string `gorm:"not null"`
// 	ProductDescription string `gorm:"not null"`
// 	ProductImage       string
// 	IsBlocked          bool `gorm:"default:false"`
// }

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

// type CategoryImages struct {
// 	ID         uint     `gorm:"primaryKey;unique;autoIncrement;not null"`
// 	CategoryID uint     `gorm:"not null"`
// 	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	ImageURL   string
// }

// type ProductImages struct {
// 	ID        uint    `gorm:"primaryKey;unique;autoIncrement;not null"`
// 	ProductID int     `gorm:"not null"`
// 	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	ImageURL  string
// }

type JSONB map[string]interface{}

func NewJsonB() JSONB {
	return make(JSONB)
}

// Value transforms the type to database driver compatible type.
func (j JSONB) Value() (driver.Value, error) {
	v, err := json.Marshal(j)
	return v, err
}

// Scan take the raw data that comes from database and convert it as a Go type.
// The reverse process of Value method.
func (j *JSONB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return nil
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interfaceP{}) failed")
	}

	return nil
}
