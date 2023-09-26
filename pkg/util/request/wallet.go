package request

import "time"

type WalletTransactionHistory struct {
	ID              uint      `json:"id"`
	TransactionTime time.Time `json:"transaction_time"`
	UserID          int       `json:"user_id"`
	Amount          float32   `json:"amount"`
	TransactionType string    `json:"transaction_type"` // "credit" or "debit"
}
