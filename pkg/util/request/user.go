package request

type SignUpData struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    int    `json:"-"`
	Password string `json:"password" binding:"required"`
	Uuid     string `json:"uuid" validate:"required"` //for retrieve user phone from the map
}

type LoginData struct {
	Phone    int    `json:"phone" validate:"required" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Address struct {
	ID               uint   `json:"-"`
	UserID           uint   `json:"-"`
	Name             string `json:"name" binding:"required"`
	PhoneNumber      string `json:"phone_number" binding:"required"`
	Pincode          string `json:"pincode" binding:"required"`
	Locality         string `json:"locality" `
	AddressLine      string `json:"address_line" binding:"required"`
	District         string `json:"district"`
	StateID          int    `json:"state_id"`
	Landmark         string `json:"landmark"`
	AlternativePhone string `json:"alternative_phone"`
	IsDefault        bool   `json:"-"`
}

type Name struct {
	Name string `json:"name" validate:"required"`
}
