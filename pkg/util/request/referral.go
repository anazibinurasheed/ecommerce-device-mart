package request

type Referral struct {
	Code string `json:"code" binding:"required"`
}
