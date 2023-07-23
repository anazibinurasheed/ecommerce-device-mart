package response

import "time"

type UserData struct {
	Id        int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Password  string `json:"password,omitempty"`
	IsAdmin   bool   `json:"is_admin"`
	IsBlocked bool   `json:"is_blocked"`
	CreatedAt time.Time
}

type Address struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"user_id"`
	Name             string `json:"name"`
	PhoneNumber      string `json:"phone_number"`
	Pincode          string `json:"pincode"`
	Locality         string `json:"locality"`
	AddressLine      string `json:"address_line"`
	District         string `json:"district"`
	StateId          int    `json:"state_id"`
	State            string `json:"state"`
	Landmark         string `json:"landmark"`
	AlternativePhone string `json:"alternative_phone"`
	IsDefault        bool
}
type States struct {
	ID   uint   `gorm:"primaryKey;unique;autoIncrement;not null"`
	Name string `gorm:"not null;unique"`
}

type Profile struct {
	Id        int
	UserName  string
	Email     string
	Phone     int
	Addresses []Address
}
