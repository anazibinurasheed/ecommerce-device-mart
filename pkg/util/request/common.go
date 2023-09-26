package request

type Otp struct {
	Otp  string `json:"otp" binding:"required"`
	UUID string `json:"uuid" binding:"required"`
}
type ChangePassword struct {
	NewPassword   string `json:"new_password" binding:"required"`
	ReNewPassword string `json:"re_new_password" binding:"required"`
}
type OldPassword struct {
	Password string `json:"old_password" binding:"required"`
}
type Phone struct {
	Phone int `json:"phone" validate:"required,min=10" binding:"required"`
}
