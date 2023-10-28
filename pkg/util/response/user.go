package response

import "time"

type UserData struct {
	ID        int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	Password  string    `json:"-"`
	IsAdmin   bool      `json:"-"`
	IsBlocked bool      `json:"is_blocked"`
	CreatedAt time.Time `json:"created_at"`
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
	StateID          int    `json:"state_id"`
	State            string `json:"state"`
	Landmark         string `json:"landmark"`
	AlternativePhone string `json:"alternative_phone"`
	IsDefault        bool   `json:"is_default"`
}

type States struct {
	ID   uint   `json:"id"`
	Name string `json:"state_name"`
}

type Profile struct {
	ID        int       `json:"user_id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	Addresses []Address `json:"addresses"`
}
