package response

type Referral struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Code   string `json:"refferal_code"`
}
