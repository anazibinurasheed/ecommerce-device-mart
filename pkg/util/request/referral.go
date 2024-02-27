package request

type Referral struct {
	Code string `json:"code" validate:"required"`
}
