package domain

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey;unique;autoIncrement;not null"`
	UserName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Phone     int    `gorm:"not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool   `gorm:"default:false"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Addresses struct {
	ID               uint   `json:"id" gorm:"primaryKey;unique;autoIncrement;not null"`
	Name             string `json:"name"`
	PhoneNumber      string `json:"phone_number"`
	Pincode          string `gorm:"pincode"`
	Locality         string `json:"locality"`
	AddressLine      string `json:"address_line"`
	District         string `json:"district"`
	StateID          uint   `json:"state_id" gorm:"not null"`
	State            State  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Landmark         string `json:"landmark"`
	AlternativePhone string `json:"alternative_phone"`
	UserID           uint   `gorm:"not null"`
	User             User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsDefault        bool   `gorm:"default:false"`
}

type State struct {
	ID   uint   `gorm:"primaryKey;unique;autoIncrement;not null"`
	Name string `gorm:"not null;unique"`
}
