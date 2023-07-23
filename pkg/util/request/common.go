package request

type Otp struct {
	Otp string `json:"otp" binding:"required"`
	ID  string `json:"unique_id" binding:"required"`
}
type ChangePassword struct {
	NewPassword   string `json:"new_password"`
	ReNewPassword string `json:"re_new_password"`
}
type OldPassword struct {
	Password string `json:"old_password"`
}
type Phone struct {
	Phone int `json:"phone" binding:"required"`
}
