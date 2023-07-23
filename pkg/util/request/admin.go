package request

// for bind the admin login request
type LoginSudoAdmin struct {
	Username string `json:"username" binding:"required"`
	// Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}
