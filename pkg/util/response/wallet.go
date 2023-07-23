package response

type Wallet struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Amount float32 `json:"amount"`
}
